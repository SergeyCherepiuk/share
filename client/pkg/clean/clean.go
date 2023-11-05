package clean

import (
	"os"
	"os/signal"
	"syscall"
)

var callbacks = make([]func(), 0)

func Listen() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	for i := len(callbacks) - 1; i >= 0; i-- {
		callbacks[i]()
	}
	os.Exit(1)
}

func Add(callback func()) {
	callbacks = append(callbacks, callback)
}
