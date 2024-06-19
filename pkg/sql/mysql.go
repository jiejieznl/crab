package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newMySQLDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", conf.MYSQL.Username, conf.MYSQL.Password, conf.MYSQL.Addr, conf.MYSQL.Port, conf.MYSQL.DbName, conf.MYSQL.Config)
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(fmt.Sprintf("gorm open err:%v", err))
	}
	s, err := mysqlDB.DB()
	if err != nil {
		panic(fmt.Sprintf("gorm get sql err:%v", err))
	}
	s.SetMaxIdleConns(conf.MYSQL.MaxIdleConns)
	s.SetMaxOpenConns(conf.MYSQL.MaxOpenConns)
	return mysqlDB
}
