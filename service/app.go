package service

import (
	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
)

type App struct {
	cfg config.C
	UR  *storage.UserRepository
}

func NewApp(cfg config.C, ur *storage.UserRepository) *App {
	return &App{
		cfg,
		ur,
	}
}
