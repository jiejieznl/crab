package model

import (
	"crab/model/example"
	"gorm.io/gorm"
)

func InstallTable(db *gorm.DB) error {
	return db.AutoMigrate(example.User{})
}
