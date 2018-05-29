package storage

import (
	"github.com/jehaby/webapp102/entity"
	"github.com/pkg/errors"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jehaby/webapp102/config"
)

func NewDB(cfg config.DB) (*sqlx.DB, error) {
	return sqlx.Open("postgres", cfg.Conn)
}

func NewPGDB(cfg config.DB) (*pg.DB, error) {
	options, err := pg.ParseURL(cfg.URL)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't parse url (%s)", cfg.URL)
	}
	db := pg.Connect(options)
	err = createSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.Ad)(nil),
		(*entity.User)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			// Temp: true,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
