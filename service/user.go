package service

import (
	"github.com/jehaby/webapp102/storage"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	Repo *storage.UserRepository // TODO: make unexported maybe
}

func newUserService (db *sqlx.DB) *UserService {
	return &UserService{
		Repo: storage.NewUserRepository(db),
	}
}