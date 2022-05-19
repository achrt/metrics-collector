package handlers

import (
	"context"
	"log"
	"testing"

	"github.com/achrt/metrics-collector/cmd/server/application"
)

var (
	h *Handler
)

func TestMain(m *testing.M) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := application.New(cancel)
	if err != nil {
		log.Fatal(err)
	}

	h = New(app)
	h.Router()
	m.Run()
}
