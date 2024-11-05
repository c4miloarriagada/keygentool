package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func OnCloseHook(hookFunc func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		hookFunc()
		os.Exit(0)
	}()
}
