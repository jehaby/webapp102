package storage

import (
	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/entity"

	_ "github.com/lib/pq"
)

type ManufacturerRepository struct {
	db *pg.DB
}

func NewManufacturerRepository(db *pg.DB) *ManufacturerRepository {
	return &ManufacturerRepository{db}
}

func (mr *ManufacturerRepository) GetAll() ([]*entity.Manufacturer, error) {
	var manufacturers []*entity.Manufacturer
	err := mr.db.Model(&manufacturers).Select()
	if err != nil {
		return nil, err
	}
	return manufacturers, nil
}
