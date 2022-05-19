package main

import (
	"context"
	"log"

	"github.com/achrt/metrics-collector/cmd/server/application"
	"github.com/achrt/metrics-collector/cmd/server/handlers"

	sc "github.com/achrt/metrics-collector/internal/controller"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go sc.Run(ctx, cancel)

	app, err := application.New()
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.New(app)
	handler.Router()

	go run(app, ctx, cancel)
	<-ctx.Done()
}

func run(app *application.App, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	log.Fatal(app.Run())
}
