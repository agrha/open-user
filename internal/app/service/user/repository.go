package user

import (
	"open-user/internal/app/service/dbtx"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Db *gorm.DB
	dbtx.DbTxRepositoryInterface
}

func NewUserRepository(db *gorm.DB, dbTxRepository dbtx.DbTxRepositoryInterface) *UserRepository {
	return &UserRepository{
		Db:                      db,
		DbTxRepositoryInterface: dbTxRepository,
	}
}
