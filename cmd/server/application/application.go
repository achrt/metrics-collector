package application

import (
	"github.com/achrt/metrics-collector/internal/domain/repositories"
	"github.com/achrt/metrics-collector/internal/storage"
	"github.com/gin-gonic/gin"
)

type App struct {
	Store  repositories.Storage
	Router *gin.Engine
}

func New() *App {
	gin.SetMode(gin.ReleaseMode)
	return &App{
		Store:  storage.New(),
		Router: gin.New(),
	}
}

func (a *App) Run(address string) error {
	return a.Router.Run(address)
}
