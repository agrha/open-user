package dbtx

import (
	"database/sql"
)

type DbTxRepository struct {
	db *sql.DB
}

func NewDbTxRepository(db *sql.DB) *DbTxRepository {
	return &DbTxRepository{
		db: db,
	}
}
