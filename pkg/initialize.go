package pkg

import (
	"crab/pkg/cache"
	"crab/pkg/config"
	"crab/pkg/log"
	"crab/pkg/server"
	"crab/pkg/sql"
)

func Initialize() {
	cfg := config.GetConfig()
	log.Register(&cfg.Log)
	cache.Register(&cfg.Cache)
	sql.Register(&cfg.DB)
	server.Register(&cfg.Server)
}
