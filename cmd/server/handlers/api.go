package handlers

import (
	"net/http"

	"github.com/achrt/metrics-collector/cmd/server/application"
	"github.com/achrt/metrics-collector/internal/domain/repositories"
)

type Handler struct {
	store repositories.Storage
}

func New(app *application.App) *Handler {
	return &Handler{
		store: app.Store,
	}
}

func (h *Handler) Router() {
	http.HandleFunc("/health", h.Health)
	http.HandleFunc("/update/", h.Update)
}
