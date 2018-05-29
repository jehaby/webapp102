package storage

import (
	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/entity"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

func NewAdRepository(db *sqlx.DB, pgdb *pg.DB) *AdRepository {
	return &AdRepository{db, pgdb}
}

type AdRepository struct {
	db *sqlx.DB // TODO: remove

	pgdb *pg.DB
}

func (ar *AdRepository) GetByUUID(uuid uuid.UUID) (*entity.Ad, error) {
	ad := &entity.Ad{}
	err := ar.pgdb.Model(ad).
		Relation("Category").
		Relation("Locality").
		Relation("Brand").
		Relation("User").
		Where("ad.uuid = ?", uuid).
		Select()

	if err != nil {
		// TODO: maybe no need to wrap?
		return nil, errors.Wrapf(err, "AdRepository.GetByUUID: '%s'", uuid.String())
	}

	return ad, nil
}
