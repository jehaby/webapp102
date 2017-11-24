package service

import (
	"github.com/jehaby/webapp102/storage"
	"github.com/jmoiron/sqlx"
)

type AdService struct {
	Repo *storage.AdRepository // TODO: make unexported maybe
}

func newAdService(db *sqlx.DB) *AdService {
	return &AdService{
		Repo: storage.NewAdRepository(db),
	}
}
