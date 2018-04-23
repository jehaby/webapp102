package storage

import (
	"fmt"

	"github.com/jehaby/webapp102/entity"

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
	db := pg.Connect(&pg.Options{
		User:     cfg.User,
		Database: cfg.Database,
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	})

	err := createSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.Ad)(nil),
		(*entity.User)(nil),
		(*entity.Product)(nil),
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
