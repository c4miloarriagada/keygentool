package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func Close(hookFunc func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		hookFunc()
		os.Exit(0)
	}()
}
