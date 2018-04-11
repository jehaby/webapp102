package service

import (
	"github.com/jehaby/webapp102/storage"
	"github.com/jmoiron/sqlx"
)

type AdService struct {
	Repo adServiceRepos
}

type adServiceRepos struct {
	Ad       *storage.AdRepository
	Category *storage.CategoryRepository
}

func newAdService(db *sqlx.DB) *AdService {
	return &AdService{
		Repo: adServiceRepos{
			Ad:       storage.NewAdRepository(db),
			Category: storage.NewCategoryRepository(db),
		},
	}
}
