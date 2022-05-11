package exit_signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Wait() chan os.Signal {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	return exitSignal
}
