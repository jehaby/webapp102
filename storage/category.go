package storage

import (
	"github.com/jehaby/webapp102/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (cr *CategoryRepository) GetAll() ([]*entity.Category, error) {
	res := make([]*entity.Category, 0, 100)
	err := cr.db.Select(&res, "SELECT id, name, path FROM categories")
	if err != nil {
		return nil, err
	}
	return res, nil
}
