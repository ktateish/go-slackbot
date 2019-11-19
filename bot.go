package slackbot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"unsafe"

	"github.com/ktateish/go-slackbot/brain"
	"github.com/ktateish/go-slackbot/iface/islack"
	"github.com/nlopes/slack"
)

func assert(exp bool) {
	if exp {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	panic(fmt.Sprintf("ASSERT!!: %s:%d @0x%x\n", file, line, pc))
}

// Bot represents a slack bot framework
type Bot struct {
	islack.RTM
	logger *log.Logger
	info   *slack.Info
	brain  Brain

	loaded   bool
	onLoads  []func(bot *Bot) error
	hears    []func(*slack.Message) error
	responds []func(*slack.Message) error
}

// NewBot creates a Bot instance
func NewBot(rtm islack.RTM) (*Bot, error) {
	brain, err := NewBrain()
	if err != nil {
		return nil, fmt.Errorf("getting new brain: %w", err)
	}
	return &Bot{
		RTM:   rtm,
		brain: brain,
	}, nil
}

func (bot *Bot) OnLoaded(action func(bot *Bot) error) {
	bot.onLoads = append(bot.onLoads, action)
}

func (bot *Bot) Brain() Brain {
	return bot.brain
}

func (bot *Bot) SetBrain(br Brain) {
	bot.brain = br
}

// SetLogger sets logger for the bot
func (bot *Bot) SetLogger(logger *log.Logger) {
	bot.logger = logger
}

// Logf print a log message with format
func (bot *Bot) Logf(format string, v ...interface{}) {
	if bot.logger == nil {
		log.Printf(format, v...)
	} else {
		bot.logger.Printf(format, v...)
	}
}

// Logf print a log message
func (bot *Bot) Log(v ...interface{}) {
	if bot.logger == nil {
		log.Print(v...)
	} else {
		bot.logger.Print(v...)
	}
}

func (bot *Bot) setInfo(info *slack.Info) {
	bot.info = info
}

// IncomingEvents retract chan slack.RTMEvent in the undelying slack.RTM
// from bot.  It is needed because slack.RTM exposes the channel itself
// instead of API.  It panics if non-slack.RTM was given on calling
// NewBot()
func (bot *Bot) IncomingEvents() chan slack.RTMEvent {
	rtm, ok := bot.RTM.(*slack.RTM)
	if ok {
		return rtm.IncomingEvents
	}

	closed := make(chan slack.RTMEvent)
	close(closed)

	v := reflect.ValueOf(bot.RTM).Elem()
	f, ok := v.Type().FieldByName("IncomingEvents")
	if !ok {
		return closed
	}
	if f.Type != reflect.TypeOf(make(chan slack.RTMEvent)) {
		return closed
	}

	ptr := v.FieldByName("IncomingEvents").Addr().Pointer()
	ch := (*chan slack.RTMEvent)(unsafe.Pointer(ptr))

	return *ch
}

// Hear registers action for all messages that re matches.
// The bot proceeds action if such a message received.
func (bot *Bot) Hear(re *regexp.Regexp, action func(ctx context.Context, res *Response) error) {
	f := func(msg *slack.Message) error {
		match := re.FindStringSubmatch(msg.Text)
		if len(match) == 0 {
			return nil
		}
		res := bot.NewResponse(msg)
		res.SetMatch(match)
		return action(context.Background(), res)
	}
	bot.hears = append(bot.hears, f)
}

// Respond registers action only for the messages that mentiond to the bot and matched by the re.
// The bot proceeds action if such a message received.
func (bot *Bot) Respond(re *regexp.Regexp, action func(ctx context.Context, res *Response) error) {
	f := func(msg *slack.Message) error {
		match := re.FindStringSubmatch(msg.Text)
		if len(match) == 0 {
			return nil
		}
		res := bot.NewResponse(msg)
		res.SetMatch(match)
		return action(context.Background(), res)
	}
	bot.responds = append(bot.responds, f)
}

func (bot *Bot) handleMessageEvent(ev *slack.MessageEvent) {
	msg := (*slack.Message)(ev)
	msgjson, err := json.Marshal(msg)
	if err != nil {
		msgjson = []byte("{}")
	}
	bot.Log("MessageEvent:", string(msgjson))
	if len(msg.BotID) != 0 && msg.User == bot.info.User.ID {
		return
	}

	// menthion to bot?
	if m := strings.TrimSpace(msg.Text); strings.HasPrefix(m, fmt.Sprintf("<@%s>", bot.info.User.ID)) {
		for i, f := range bot.responds {
			bot.Logf("responds[%d]: start", i)
			err := f(msg)
			if err != nil {
				bot.Logf("responds[%d]: %s", i, err)
			}
			bot.Logf("responds[%d]: end", i)
		}
	}

	for i, f := range bot.hears {
		bot.Logf("hears[%d]: start", i)
		err := f(msg)
		if err != nil {
			bot.Logf("hears[%d]: %s", i, err)
		}
		bot.Logf("hears[%d]: end", i)
	}

	bot.Log("MessageEvent: done")
}

func (bot *Bot) Loaded() bool {
	return bot.loaded
}

func (bot *Bot) handleConnected(ev *slack.ConnectedEvent) {
	bot.Log("Connected")
	bot.Log("Infos:", ev.Info)
	bot.setInfo(ev.Info)
	bot.Log("Connection counter:", ev.ConnectionCount)

	bot.Log("Proceed on-load handlers")
	for i, f := range bot.onLoads {
		err := f(bot)
		if err != nil {
			bot.Logf("E: loading[%d]: %s", i, err.Error())
		}
	}
	bot.Log("done")
	bot.loaded = true
}

// Run starts listening slack events
func (bot *Bot) Run() error {
	go bot.ManageConnection()

	// for unloaded state
	for msg := range bot.IncomingEvents() {
		if bot.Loaded() {
			break
		}
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			bot.handleConnected(ev)

		case *slack.RTMError:
			bot.Logf("Event Received: Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			bot.Logf("Event Received: Invalid credentials")
			return fmt.Errorf("error\n")

		default:
			fmt.Printf("Event Replace: Unexpected: %v\n", msg.Data)
		}
	}

	// handling for loaded state
	for msg := range bot.IncomingEvents() {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			bot.handleMessageEvent(ev)

		case *slack.PresenceChangeEvent:
			fmt.Print("Event Received: ")
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Print("Event Received: ")
			fmt.Printf("Current latency: %v\n", ev.Value)

		//case *slack.DesktopNotificationEvent:
		//	fmt.Printf("Desktop Notification: %v\n", ev)

		case *slack.RTMError:
			fmt.Print("Event Received: ")
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Print("Event Received: ")
			fmt.Printf("Invalid credentials")
			return fmt.Errorf("error\n")

		default:
			fmt.Print("Event Received: ")
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
			fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
	return nil
}

// NewResponse creates a Response object and returns the pointer to it.
func (bot *Bot) NewResponse(msg *slack.Message) *Response {
	res := &Response{
		bot: bot,
		msg: msg,
		opts: []slack.MsgOption{
			slack.MsgOptionAsUser(true),
			slack.MsgOptionUser(bot.info.User.ID),
		},
	}
	return res
}

func (res *Response) SetMatch(match []string) {
	res.match = match
}

// Response represents response for the received message
type Response struct {
	bot   *Bot
	msg   *slack.Message
	opts  []slack.MsgOption
	match []string
}

// Match returns a slice of matched string from the message that bot received.
// 0th item is the whole text.
func (res *Response) Match() []string {
	return res.match
}

// Send sends msg to the channel
func (res *Response) Send(ctx context.Context, msg string) error {
	opts := append(res.opts, slack.MsgOptionText(msg, false))
	ch, txt, ts, err := res.bot.RTM.SendMessageContext(ctx, res.msg.Channel, opts...)
	if err != nil {
		return fmt.Errorf("sending message: %w", err)
	}
	res.bot.Logf("Send: ch=%s text=%s ts=%s", ch, txt, ts)
	return nil
}

type Brain interface {
	Load(key string) ([]byte, error)
	Save(key string, val []byte) error
}

func NewBrain() (Brain, error) {
	return brain.NewOnMemoryBrain(), nil
}
