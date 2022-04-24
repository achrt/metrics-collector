package main

import (
	"log"
	"net"

	"github.com/achrt/metrics-collector/cmd/server/application"
	"github.com/achrt/metrics-collector/cmd/server/handlers"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {

	app := application.New()
	handler := handlers.New(app)
	handler.Router()

	log.Fatal(app.Run(net.JoinHostPort(host, port)))
}
