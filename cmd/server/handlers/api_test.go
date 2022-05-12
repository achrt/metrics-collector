package handlers

import (
	"log"
	"testing"

	"github.com/achrt/metrics-collector/cmd/server/application"
)

var (
	h *Handler
)

func TestMain(m *testing.M) {
	app, err := application.New()
	if err != nil {
		log.Fatal(err)
	}

	h = New(app)
	h.Router()
	m.Run()
}
