package application

import (
	"context"
	"log"

	"github.com/achrt/metrics-collector/internal/domain/repositories"
	"github.com/achrt/metrics-collector/internal/storage"
	"github.com/gin-gonic/gin"
)

type App struct {
	Store  repositories.Storage
	Router *gin.Engine

	address string
}

func New(cancel context.CancelFunc) (*App, error) {
	cfg := Config{}
	if err := cfg.loadConfiguration(); err != nil {
		return nil, err
	}

	s, err := storage.New(cfg.StoreFile, cfg.StoreInterval, cancel)
	if err != nil {
		return nil, err
	}

	if cfg.Restore {
		if err = s.Load(); err != nil {
			return nil, err
		}
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	a := &App{
		Store:   s,
		Router:  router,
		address: cfg.Address,
	}
	return a, nil
}

func (a *App) Run() error {
	log.Println("server is up and running in address: ", a.address)
	return a.Router.Run(a.address)
}
