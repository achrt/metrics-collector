package controller

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChanel
}
