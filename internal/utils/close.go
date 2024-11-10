package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func Close() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		onSave()
		os.Exit(0)
	}()
}

func onSave() {

}
