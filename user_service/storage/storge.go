package storage

import (
	"github.com/jackc/pgx/v5"
	"gitlab.com/e-market724/back-end/user_service/storage/postgres"
	"gitlab.com/e-market724/back-end/user_service/storage/repoi"
)

type StorageRepoI interface {
	GetUserRepo() repoi.UserRepoI
}

type storageRepo struct {
	userRepo repoi.UserRepoI
}

func NewUserRepo(db *pgx.Conn) StorageRepoI {

	return &storageRepo{userRepo: postgres.NewUserRepo(db)}
}

func (s *storageRepo) GetUserRepo() repoi.UserRepoI {

	return s.userRepo
}
