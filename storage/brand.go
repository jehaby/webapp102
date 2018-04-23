package storage

import (
	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/entity"

	_ "github.com/lib/pq"
)

type BrandRepository struct {
	db *pg.DB
}

func NewBrandRepository(db *pg.DB) *BrandRepository {
	return &BrandRepository{db}
}

func (mr *BrandRepository) GetAll() ([]*entity.Brand, error) {
	var brands []*entity.Brand
	err := mr.db.Model(&brands).Select()
	if err != nil {
		return nil, err
	}
	return brands, nil
}
