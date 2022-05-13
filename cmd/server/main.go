package main

import (
	"log"

	"github.com/achrt/metrics-collector/cmd/server/application"
	"github.com/achrt/metrics-collector/cmd/server/handlers"
)

func main() {

	app, err := application.New()
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.New(app)
	handler.Router()

	log.Fatal(app.Run())
}
