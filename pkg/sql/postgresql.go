package sql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgresqlDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.PGSQL.Addr, conf.PGSQL.Username, conf.PGSQL.Password, conf.PGSQL.DbName, conf.PGSQL.Port)
	postgresqlDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 可选设置
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("gorm open err:%v", err))
	}
	s, err := postgresqlDB.DB()
	if err != nil {
		panic(fmt.Sprintf("gorm get sql err:%v", err))
	}
	s.SetMaxIdleConns(conf.PGSQL.MaxIdleConns)
	s.SetMaxOpenConns(conf.PGSQL.MaxOpenConns)
	return postgresqlDB
}
