package cache

import (
	"context"
	"errors"
	"time"

	"github.com/goccy/go-json"

	"github.com/patrickmn/go-cache"
)

func newCacheMemory() (ICache, error) {
	return &implCacheMemory{cli: cache.New(
		time.Minute*time.Duration(conf.Local.DefaultExpiration),
		time.Minute*time.Duration(conf.Local.CleanupInterval))}, nil
}

type implCacheMemory struct {
	cli *cache.Cache
}

func (i *implCacheMemory) Get(_ context.Context, key string, value any) error {
	v, ok := i.cli.Get(key)
	if !ok {
		return errors.New("no cache data")
	}
	data, ok := v.([]byte)
	if !ok {
		return errors.New("data not []byte type")
	}
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	return nil
}

func (i *implCacheMemory) Set(_ context.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if expiration <= 0 {
		expiration = -1
	}
	i.cli.Set(key, data, expiration)
	return nil
}

func (i *implCacheMemory) Del(_ context.Context, key string) error {
	i.cli.Delete(key)
	return nil
}
