package brain

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/ktateish/go-slackbot/iface/iredis/v7"
)

type RedisBrain struct {
	rc iredis.Client
}

func NewRedisBrain(rc iredis.Client) (*RedisBrain, error) {
	_, err := rc.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("initial checking: %w", err)
	}
	return &RedisBrain{
		rc: rc,
	}, nil
}

func (br *RedisBrain) Load(key string) ([]byte, error) {
	v, err := br.rc.Get(key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("getting value for key='%s': %w", key, err)
	}
	return v, err
}

func (br *RedisBrain) Save(key string, val []byte) error {
	err := br.rc.Set(key, val, 0).Err()
	if err != nil {
		return fmt.Errorf("setting value for key='%s': %w", key, err)
	}
	return nil
}
