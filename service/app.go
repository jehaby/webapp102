package service

import (
	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
	"go.uber.org/zap"
)

type App struct {
	cfg    config.C
	UR     *storage.UserRepository
	Logger *zap.SugaredLogger
}

func NewApp(cfg config.C, ur *storage.UserRepository) *App {
	return &App{
		cfg:    cfg,
		UR:     ur,
		Logger: getLogger(cfg),
	}
}

func getLogger(cfg config.C) *zap.SugaredLogger {
	// TODO: prod logging
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
