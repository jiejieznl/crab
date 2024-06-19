package sql

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newSQLiteDB() *gorm.DB {
	var err error
	if conf.SQLITE.DbPath == "" {
		conf.SQLITE.DbPath = "./sql.db"
		conf.SQLITE.MaxIdleConns = 10
		conf.SQLITE.MaxOpenConns = 100
	}
	sqliteDB, err := gorm.Open(sqlite.Open(conf.SQLITE.DbPath), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(fmt.Sprintf("gorm open err:%v", err))
	}
	s, err := sqliteDB.DB()
	if err != nil {
		panic(fmt.Sprintf("gorm get sql err:%v", err))
	}
	s.SetMaxIdleConns(conf.SQLITE.MaxIdleConns)
	s.SetMaxOpenConns(conf.SQLITE.MaxOpenConns)
	return sqliteDB
}
