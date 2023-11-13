package clean

import (
	"os"
	"os/signal"
	"syscall"
)

var callbacks = make([]func(), 0)

func Clean() {
	for i := len(callbacks) - 1; i >= 0; i-- {
		callbacks[i]()
	}
}

func Add(callback func()) {
	callbacks = append(callbacks, callback)
}

func InterceptInterruption() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	Clean()
	os.Exit(1)
}
