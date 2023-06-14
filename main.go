package main

import (
	"github.com/rs/zerolog/log"
	"main.go/internal/telegram"
	"main.go/internal/torrent_client"
	"main.go/internal/web_server"
	"main.go/pkg/config"
	"main.go/pkg/osutils"
)

func main() {
	osutils.CallFuncByInterrupt(func() {
		telegram.SendMessageAdmin("Close app host " + *web_server.HostAndPort)
	})

	listInitFunc := map[string]func() error{
		"env":       config.Init,
		"WebServer": web_server.Init,
		"telegram":  telegram.Init,
		"torrent":   torrent_client.Init,
	}
	for namePackage, initFunc := range listInitFunc {
		err := initFunc()
		if err != nil {
			log.Fatal().Err(err).Msg(namePackage)
			return
		}
	}

	log.Info().Str("PATH_TORRENT_CONTENT", torrent_client.PathTorrentContent).Msg("Start bot https://t.me/" + telegram.Bot.Self.UserName)
	telegram.Listener()

	torrent_client.Close()
}
