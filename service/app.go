package service

import (
	"log"

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

	userService := newUserService(db)
	adService := newAdService(db)

	return &App{
		cfg:    cfg,
		User:   userService,
		Ad:     adService,
		Repo:   newRepos(db),
		Logger: getLogger(cfg),
	}
}

type repos struct {
	Manufacturer *storage.ManufacturerRepository
	Ad           *storage.AdRepository
	Category     *storage.CategoryRepository
}

func newRepos(db *sqlx.DB) *repos {
	return &repos{
		Manufacturer: storage.NewManufacturerRepository(db),
		Ad:           storage.NewAdRepository(db),
		Category:     storage.NewCategoryRepository(db),
	}
}

func getLogger(cfg config.C) *zap.SugaredLogger {
	// TODO: prod logging
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
