package config

import (
	"crab/pkg/cache"
	"crab/pkg/jwt"
	"crab/pkg/log"
	"crab/pkg/server"
	"crab/pkg/sql"
	"github.com/google/uuid"
)

type Config struct {
	Log    log.Config    `toml:"log"`
	Server server.Config `toml:"server"`
	Cache  cache.Config  `toml:"cache"`
	DB     sql.Config    `toml:"sql"`
	JWT    jwt.Config    `toml:"jwt"`
}

func setDefaultConfig() Config {
	var config Config

	config.Server = server.Config{
		Mode: "debug",
		Port: 8888,
	}

	config.Cache = cache.Config{
		Source: "local",
		Local: cache.Local{
			DefaultExpiration: 1,
			CleanupInterval:   0,
		},
	}

	config.DB = sql.Config{
		Source: "sqlite",
		SQLITE: sql.SQLITE{
			DbPath:       "./sql.db",
			MaxIdleConns: 10,
			MaxOpenConns: 100,
		},
	}

	config.JWT = jwt.Config{SigningKey: uuid.NewString()}
	return config
}
