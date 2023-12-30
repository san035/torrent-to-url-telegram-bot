package osutils

import (
	"log/slog"
	"os"
	"os/signal"
)

func CallFuncByInterrupt(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		f()
		slog.Info("User interrupt")
		os.Exit(0)
	}()
}
