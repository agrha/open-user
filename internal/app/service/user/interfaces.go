package user

import "open-user/internal/app/service/dbtx"

type UserRepositoryInterface interface {
	dbtx.DbTxRepositoryInterface
}

type UserServiceInterface interface {
}
