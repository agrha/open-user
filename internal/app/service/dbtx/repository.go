package dbtx

import (
	"github.com/jinzhu/gorm"
)

type DbTxRepository struct {
	db *gorm.DB
}

func NewDbTxRepository(db *gorm.DB) *DbTxRepository {
	return &DbTxRepository{
		db: db,
	}
}
