package controller

import (
	"os"
	"syscall"
	"testing"
)

// TODO: сомнительный тест без проверки поведения, подумать как лучше
func TestRun(t *testing.T) {
	signalChanel := make(chan os.Signal, 1)

	signals := []os.Signal{
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	for _, sig := range signals {
		go run(signalChanel)
		signalChanel <- sig
	}
}
