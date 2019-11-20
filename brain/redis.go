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
	cmd := br.rc.Get(key)

	err := br.rc.ProcessContext(ctx, cmd)
	if err != nil {
		return nil, fmt.Errorf("getting from redis: %w", err)
	}

	v, err := cmd.Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("converting fetched result: %w", key, err)
	}
	return v, err
}

func (br *RedisBrain) Save(ctx context.Context, key string, val []byte) error {
	cmd := br.rc.Set(key, val, 0)

	err := br.rc.ProcessContext(ctx, cmd)
	if err != nil {
		return fmt.Errorf("setting value into redis: %w", err)
	}
	err = cmd.Err()
	if err != nil {
		return fmt.Errorf("setting value into redis: %w", err)
	}
	return nil
}
