package application

import (
	"context"
	"testing"
	"time"
)

const (
	duration       = 1
	reportInterval = 2
	msaddress      = ""
)

var app *App

func TestMain(m *testing.M) {
	app = New(reportInterval, duration, msaddress)
	m.Run()
}

func TestRunContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go app.Run(ctx, cancel)

	<-time.After(time.Duration(3) * time.Second)
	cancel()
}
