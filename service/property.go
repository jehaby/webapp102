package service

import (
	"context"

	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/log"
)

type PropertyService struct {
	db  *pg.DB
	log *log.Logger
}

func NewPropertyService(db *pg.DB, log *log.Logger) *PropertyService {
	return &PropertyService{
		db:  db,
		log: log,
	}
}

func (ps *PropertyService) GetByCategory(ctx context.Context, categoryID int64) ([]*entity.Property, error) {

	res := []*entity.Property{}

	// TODO: error if not leaf category?

	err := ps.db.Model(&res).Where("category_id = ?", categoryID).Select()
	if err != nil {
		return nil, err
	}

	return res, nil
}
