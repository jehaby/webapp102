package storage

import (
	"time"

	"github.com/davecgh/go-spew/spew"
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
		Relation("Component").
		Relation("User").
		Where("ad.uuid = ?", uuid).
		Select()

	spew.Dump("getbyuuid", ad)

	if err != nil {
		// TODO: maybe no need to wrap?
		return nil, errors.Wrapf(err, "AdRepository.GetByUUID: '%s'", uuid.String())
	}

	return ad, nil
}

const createAdQuery = `INSERT INTO  ads (uuid, name, description, user_uuid, category_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

func (ar *AdRepository) Create(ad entity.Ad) (*entity.Ad, error) {
	ad.UUID = uuid.NewV4()
	ad.CreatedAt = time.Now()

	res, err := ar.db.Exec(
		createAdQuery,
		ad.UUID,
		ad.Name,
		ad.Description,
		// ad.User.UUID,
		ad.CategoryID,
		ad.CreatedAt,
	)

	if err != nil {
		return nil, errors.Wrapf(err, "creating ad: %+v", ad)
	}
	if ra, err := res.RowsAffected(); err != nil || ra != 1 {
		return nil, errors.Wrapf(err, "creating ad: %+v, rows affected: %d", ad, ra)
	}
	return &ad, nil
}

const updateAdQuery = `
UPDATE ads SET
  name = $1,
  description = $2,
  updated_at = $3
WHERE uuid = $4
`

func (ar *AdRepository) Update(ad entity.Ad) error {
	res, err := ar.db.Exec(
		updateAdQuery,
		ad.Name,
		ad.Description,
		time.Now(),
		ad.UUID,
	)

	if err != nil {
		return errors.Wrapf(err, "updating ad: %+v", ad)
	}
	if ra, err := res.RowsAffected(); err != nil || ra != 1 {
		return errors.Wrapf(err, "updating ad: %+v, rows affected: %d", ad, ra)
	}

	return nil
}

const deleteAdQuery = `UPDATE ads SET deleted_ad = $1 WHERE uuid = $2;`

func (ar *AdRepository) Delete(uid uuid.UUID) error {
	res, err := ar.db.Exec(
		deleteAdQuery,
		time.Now(),
		uid,
	)

	if err != nil {
		return errors.Wrapf(err, "deleting ad: %+v", uid)
	}
	if ra, err := res.RowsAffected(); err != nil || ra != 1 {
		return errors.Wrapf(err, "deleting ad: %+v, rows affected: %d", uid, ra)
	}

	return nil
}
