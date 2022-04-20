package signalController

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exit := make(chan string)
	go func() {
		for {
			s := <-signalChanel
			switch s {
			case syscall.SIGINT:
				exit <- syscall.SIGINT.String()
				return

			case syscall.SIGTERM:
				exit <- syscall.SIGTERM.String()
				return

			case syscall.SIGQUIT:
				exit <- syscall.SIGQUIT.String()
				return

			default:
				log.Println("unknown signal")
			}
		}
	}()

	log.Println(<-exit)
}
