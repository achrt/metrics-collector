package main

import (
	"context"

	"github.com/achrt/metrics-collector/cmd/agent/application"
)

const duration = int64(2)
const metricServerAddress = "http://127.0.0.1:8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app := application.New(10, duration, metricServerAddress)
	go app.Run(ctx, cancel)
	<-ctx.Done()
}
