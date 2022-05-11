package handlers

import (
	"testing"

	"github.com/achrt/metrics-collector/cmd/server/application"
)

var (
	h *Handler
)

func TestMain(m *testing.M) {
	app := application.New()
	h = New(app)
	h.Router()
	m.Run()
}
