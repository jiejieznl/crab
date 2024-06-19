package internal

import (
	"crab/model"
	"crab/pkg/config"
	"crab/pkg/sql"
	"log"
)

func InstallTables() {
	cfg := config.GetConfig()
	sql.Register(&cfg.DB)
	err := model.InstallTable(sql.GetDB())
	if err != nil {
		log.Fatalf("数据库迁移失败:%v", err)
	}
}
