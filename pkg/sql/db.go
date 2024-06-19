package sql

import (
	"gorm.io/gorm"
	"log"
)

var (
	conf      *Config
	runtimeDB *gorm.DB
)

func Register(cfg *Config) {
	conf = cfg
	switch conf.Source {
	case "mysql":
		runtimeDB = newMySQLDB()
	case "postgresql":
		runtimeDB = newPostgresqlDB()
	default:
		runtimeDB = newSQLiteDB()
	}
}

func GetDB() *gorm.DB {
	if runtimeDB == nil {
		log.Fatalf("%s not init", conf.Source)
	}
	return runtimeDB
}
