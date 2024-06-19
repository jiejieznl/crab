package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"

	"github.com/go-redis/redis/v8"
)

func newCacheRedis() (ICache, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Redis.Addr, conf.Redis.Port),
		Password:     conf.Redis.Password,
		DB:           conf.Redis.DB,
		DialTimeout:  time.Second * time.Duration(conf.Redis.DialTimeout),
		ReadTimeout:  time.Second * time.Duration(conf.Redis.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(conf.Redis.WriteTimeout),
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping error:%v", err)
	}
	fmt.Println("redis conn success")
	return &implCacheRedis{cli: cli}, nil
}

type implCacheRedis struct {
	cli *redis.Client
}

func (i *implCacheRedis) Get(ctx context.Context, key string, value any) error {
	data, err := i.cli.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrNotRecord
		}
		return err
	}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	return nil
}

func (i *implCacheRedis) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return i.cli.Set(ctx, key, data, expiration).Err()
}

func (i *implCacheRedis) Del(ctx context.Context, key string) error {
	return i.cli.Del(ctx, key).Err()
}
