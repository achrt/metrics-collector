package application

import (
	"context"
	"log"
	"testing"
	"time"
)

var app *App

func TestMain(m *testing.M) {
	var err error
	if app, err = New(); err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestRunContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go app.Run(ctx, cancel)

	<-time.After(time.Duration(3) * time.Second)
	cancel()
}
