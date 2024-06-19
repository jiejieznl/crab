package impl

import (
	"crab/pkg/sql"
	"gorm.io/gorm"
)

func newBaseDB() *baseDB {
	return &baseDB{
		normalDB: sql.GetDB(),
	}
}

type baseDB struct {
	normalDB, txDB *gorm.DB
}

func (_this *baseDB) db() *gorm.DB {
	if _this.txDB != nil {
		return _this.txDB
	}
	return _this.normalDB
}

func (_this *baseDB) SetTxDB(tx *gorm.DB) {
	_this.txDB = tx
}
