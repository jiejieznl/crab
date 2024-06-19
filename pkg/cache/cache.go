package cache

import (
	"context"
	"errors"
	"log"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string, value any) error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

var (
	ErrNotRecord = errors.New("cache not record")
)

var (
	runtimeCache ICache
	conf         *Config
)

func Cache() ICache {
	if runtimeCache == nil {
		log.Fatal("cache not init")
	}
	return runtimeCache
}

func Register(cfg *Config) {
	conf = cfg
	var err error
	if conf.Source == "redis" {
		runtimeCache, err = newCacheRedis()
	} else {
		runtimeCache, err = newCacheMemory()
	}
	if err != nil {
		log.Fatal(err)
	}
}
