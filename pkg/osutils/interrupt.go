package osutils

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func CallFuncByInterrupt(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		f()
		log.Error().Msg("User interrupt")
		os.Exit(0)
	}()
}
