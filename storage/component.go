package storage

import (
	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/entity"
	"github.com/pkg/errors"
)

type ComponentRepository struct {
	db *pg.DB
}

func NewComponentRepository(db *pg.DB) *ComponentRepository {
	return &ComponentRepository{db}
}

func (cr *ComponentRepository) GetByID(id uint32) (*entity.Component, error) {
	c := &entity.Component{}
	err := cr.db.Model(c).
		Relation("Manufacturer").
		Relation("Category").
		Where("component.id = ?", id).
		First()

	if err != nil {
		return nil, errors.Wrapf(err, "ComponentRepository.GetByID: %d", id)
	}
	return c, nil
}

func (cr *ComponentRepository) GetAll() ([]*entity.Component, error) {
	// TODO: pagination!
	var Components []*entity.Component
	err := cr.db.Model(&Components).Select()
	if err != nil {
		return nil, err
	}
	return Components, nil
}
