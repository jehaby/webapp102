package service

import (
	"gopkg.in/go-playground/validator.v9"

	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/pkg/log"
	"github.com/jehaby/webapp102/storage"
)

type App struct {
	cfg  config.C
	User *UserService
	Ad   *AdService

	Service services

	Repo   *repos
	Logger *log.Logger
}

func NewApp(cfg config.C) *App {

	logger := getLogger(cfg)

	db, err := storage.NewDB(cfg.DB)
	if err != nil {
		logger.Fatalf("couldn't open db: %v", err)
	}

	pgDB, err := storage.NewPGDB(cfg.DB)
	if err != nil {
		logger.Fatal("error creating pg: ", err)
	}

	val := validator.New()

	// TODO: clean up
	services := newServices(db, pgDB, val, logger)
	return &App{
		cfg:     cfg,
		User:    services.User,
		Ad:      services.Ad,
		Repo:    newRepos(db, pgDB),
		Service: newServices(db, pgDB, val, logger),
		Logger:  logger,
	}
}

type services struct {
	Ad       *AdService
	Category *CategoryService
	User     *UserService
	Property *PropertyService
}

func newServices(
	db *sqlx.DB,
	pgDB *pg.DB,
	val *validator.Validate,
	log *log.Logger,

) services {
	categoryService := NewCategoryService(pgDB, log)
	propertyService := NewPropertyService(pgDB, log)

	return services{
		Ad:       NewAdService(pgDB, val, categoryService, propertyService, log),
		Category: categoryService,
		User:     newUserService(db),
		Property: propertyService,
	}
}

type repos struct {
	Ad    *storage.AdRepository
	Brand *storage.BrandRepository
}

func newRepos(db *sqlx.DB, pgDB *pg.DB) *repos {
	return &repos{
		Ad:    storage.NewAdRepository(db, pgDB),
		Brand: storage.NewBrandRepository(pgDB),
	}
}

func getLogger(cfg config.C) *log.Logger {
	// TODO: prod logging
	logger, _ := zap.NewDevelopment()
	// logger, _ := zap.NewProduction()
	return &log.Logger{logger.Sugar()}
}
