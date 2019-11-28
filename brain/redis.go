package brain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisClient interface {
	Get(key string) *redis.StringCmd
	Ping() *redis.StatusCmd
	ProcessContext(ctx context.Context, cmd redis.Cmder) error
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type RedisBrain struct {
	rc RedisClient
}

func NewRedisBrain(rc RedisClient) (*RedisBrain, error) {
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
		if errors.Is(err, redis.Nil) {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("getting from redis: %w", err)
	}

	v, err := cmd.Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("converting fetched result: %w", err)
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
