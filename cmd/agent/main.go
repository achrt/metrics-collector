package main

import (
	"context"
	"log"

	"github.com/achrt/metrics-collector/cmd/agent/application"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := application.New()
	if err != nil {
		log.Fatal(err)
	}

	go app.Run(ctx, cancel)
	<-ctx.Done()
}
