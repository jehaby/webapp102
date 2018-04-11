package storage

import (
	"github.com/jehaby/webapp102/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ManufacturerRepository struct {
	db *sqlx.DB
}

func NewManufacturerRepository(db *sqlx.DB) *ManufacturerRepository {
	return &ManufacturerRepository{db}
}

func (mr *ManufacturerRepository) GetAll() ([]*entity.Manufacturer, error) {
	res := make([]*entity.Manufacturer, 0, 100)
	err := mr.db.Select(&res, "SELECT id, name FROM manufacturers")
	if err != nil {
		return nil, err
	}
	return res, nil
}
