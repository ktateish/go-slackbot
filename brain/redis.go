package brain

import (
	"context"
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

func (br *RedisBrain) Load(ctx context.Context, key string) ([]byte, error) {
	c := br.rc.WithContext(ctx)
	v, err := c.Get(key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("getting value for key='%s': %w", key, err)
	}
	return v, err
}

func (br *RedisBrain) Save(ctx context.Context, key string, val []byte) error {
	c := br.rc.WithContext(ctx)
	err := c.Set(key, val, 0).Err()
	if err != nil {
		return fmt.Errorf("setting value for key='%s': %w", key, err)
	}
	return nil
}
