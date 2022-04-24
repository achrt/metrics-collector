package application

import (
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/repositories"
	"github.com/achrt/metrics-collector/internal/storage"
)

type App struct {
	Store repositories.Storage
}

func New() *App {
	return &App{
		Store: storage.New(),
	}
}

func (a *App) Run(address string) error {
	return http.ListenAndServe(address, nil)
}
