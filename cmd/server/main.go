package main

import (
	"log"
	"net"

	"github.com/achrt/metrics-collector/cmd/server/application"
	"github.com/achrt/metrics-collector/cmd/server/handlers"
)

const (
	host = "127.0.0.1"
	port = "8080"
)

func main() {

	app := application.New()
	handler := handlers.New(app)
	app.Router.LoadHTMLFiles("./handlers/templates/index.tmpl")
	handler.Router()

	log.Fatal(app.Run(net.JoinHostPort(host, port)))
}
