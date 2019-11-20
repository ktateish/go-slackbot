package slackbot

import (
	"testing"

	"github.com/nlopes/slack"
)

type mockRTM struct {
	RTM
	channels       map[string]bool
	IncomingEvents chan slack.RTMEvent
}

func NewMockRTM() *mockRTM {
	res := &mockRTM{
		channels:       make(map[string]bool),
		IncomingEvents: make(chan slack.RTMEvent),
	}
	res.channels["foo"] = true
	return res
}

func TestBot_IncomingEvents(t *testing.T) {
	grtm := slack.New("none").NewRTM()
	bot, err := NewBot(grtm)
	if err != nil {
		t.Fatalf("NewBot() failed: %s\n", err)
	}

	ch := bot.IncomingEvents()
	if ch == nil {
		t.Errorf("IncomingEvents() failed for real slack RTM object\n")
	}

	bot, err = NewBot(NewMockRTM())
	if err != nil {
		t.Fatalf("NewBot() failed: %s\n", err)
	}
	ch = bot.IncomingEvents()
	if ch == nil {
		t.Errorf("IncomingEvents() failed for mockRTM object\n")
	}

	wrong := &struct {
		RTM
		IncomingEvents chan string
	}{}
	bot, err = NewBot(wrong)
	if err != nil {
		t.Fatalf("NewBot() failed: %s\n", err)
	}
	ch = bot.IncomingEvents()
	if ch == nil {
		t.Errorf("IncomingEvents() failed for wrong type object\n")
	}

	empty := &struct{ RTM }{}
	bot, err = NewBot(empty)
	if err != nil {
		t.Fatalf("NewBot() failed: %s\n", err)
	}
	ch = bot.IncomingEvents()
	if ch == nil {
		t.Errorf("IncomingEvents() failed for empty object\n")
	}
}
