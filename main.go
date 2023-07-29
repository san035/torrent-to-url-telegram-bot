package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"main.go/internal/telegram"
	"main.go/internal/torrent_client"
	"main.go/internal/web_server"
	"main.go/pkg/config"
	"main.go/pkg/osutils"
	"runtime"
)

func main() {
	defer funcEnd()
	osutils.CallFuncByInterrupt(funcEnd)

	listInitFunc := map[string]func() error{
		"env":       config.Init,
		"WebServer": web_server.Init,
		"telegram":  telegram.Init,
		"torrent":   torrent_client.Init,
	}
	defer torrent_client.Close()

	for namePackage, initFunc := range listInitFunc {
		err := initFunc()
		if err != nil {
			log.Fatal().Err(err).Msg(namePackage)
			return
		}
	}

	log.Info().Str("PATH_TORRENT_CONTENT", torrent_client.PathTorrentContent).Interface("Names bot", telegram.GetListNameBot()).Msg("Start bots")

	telegram.SendMessageAdmin("Start bots \n" + telegram.GetInfo())

	telegram.Listener()
	log.Info().Msg("End.")
}

func funcEnd() {
	totalText := "Close app host " + *web_server.HostAndPort
	r := recover()
	textRecover := fmt.Sprint("Recovered:", r)
	if r != nil {
		totalText += "\n" + textRecover
		totalText += fmt.Sprintf("\nPanic occurred at")
		for i := 0; i < 6; i++ {
			_, file, line, _ := runtime.Caller(i)
			totalText += fmt.Sprintf("\n %s:%d", file, line)
		}
	}
	log.Info().Str("recover", textRecover).Msg("Close app")
	telegram.SendMessageAdmin(totalText)
}
