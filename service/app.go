package service

import (
	"log"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
	"go.uber.org/zap"
)

type App struct {
	cfg    config.C
	User   *UserService
	Ad     *AdService
	Logger *zap.SugaredLogger
}

func NewApp(cfg config.C) *App {

	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Panicf("couldn't open db: %v", err)
	}

	userService := newUserService(db)
	adService := newAdService(db)

	return &App{
		cfg:    cfg,
		User:   userService,
		Ad:     adService,
		Logger: getLogger(cfg),
	}
}

func getLogger(cfg config.C) *zap.SugaredLogger {
	// TODO: prod logging
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
