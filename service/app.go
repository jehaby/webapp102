package service

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
)

type App struct {
	cfg    config.C
	User   *UserService
	Ad     *AdService
	Repo   *repos
	Logger *zap.SugaredLogger
}

func NewApp(cfg config.C) *App {
	db, err := storage.NewDB(cfg.DB)
	if err != nil {
		log.Panicf("couldn't open db: %v", err)
	}

	pgDB, err := storage.NewPGDB(cfg.DB)
	if err != nil {
		log.Fatal("error creating pg: ", err)
	}

	log.Println("pg created")

	userService := newUserService(db)
	adService := newAdService(db, pgDB)

	return &App{
		cfg:    cfg,
		User:   userService,
		Ad:     adService,
		Repo:   newRepos(db, pgDB),
		Logger: getLogger(cfg),
	}
}

type repos struct {
	Ad           *storage.AdRepository
	Category     *storage.CategoryRepository
	Component    *storage.ComponentRepository
	Manufacturer *storage.ManufacturerRepository
}

func newRepos(db *sqlx.DB, pgDB *pg.DB) *repos {
	return &repos{
		Ad:           storage.NewAdRepository(db, pgDB),
		Category:     storage.NewCategoryRepository(db),
		Component:    storage.NewComponentRepository(pgDB),
		Manufacturer: storage.NewManufacturerRepository(pgDB),
	}
}

func getLogger(cfg config.C) *zap.SugaredLogger {
	// TODO: prod logging
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
