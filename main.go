package main

import (
	"github.com/rs/zerolog/log"
	"main.go/bot_tg"
	"main.go/config"
	"main.go/torrent_client"
	"main.go/web_server"
)

func main() {
	listInitFunc := map[string]func() error{
		"env":       config.Init,
		"WebServer": web_server.Init,
		"telegram":  bot_tg.Init,
		"torrent":   torrent_client.Init,
	}
	for namePackage, initFunc := range listInitFunc {
		err := initFunc()
		if err != nil {
			log.Fatal().Err(err).Msg(namePackage)
			return
		}
	}

	log.Info().Str("PATH_TORRENT_CONTENT", torrent_client.PathTorrentContent).Msg("Start bot https://t.me/" + bot_tg.Bot.Self.UserName)
	bot_tg.Run()

	torrent_client.Close()
}
