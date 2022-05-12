package application

import (
	"github.com/achrt/metrics-collector/internal/domain/repositories"
	"github.com/achrt/metrics-collector/internal/storage"
	"github.com/gin-gonic/gin"
)

type App struct {
	Store  repositories.Storage
	Router *gin.Engine

	address string
}

func New() (*App, error) {
	cfg, err := loadConfiguration()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	return &App{
		Store:   storage.New(),
		Router:  router,
		address: cfg.Address,
	}, nil
}

func (a *App) Run() error {
	return a.Router.Run(a.address)
}
