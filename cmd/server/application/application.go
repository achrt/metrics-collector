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
	return &App{
		Store:  storage.New(),
		Router: gin.Default(),
	}
}

func (a *App) Run(address string) error {
	return a.Router.Run(address)
}
