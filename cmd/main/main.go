package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"main.go/internal/download_clients"
	"main.go/internal/download_clients/torrent_anacrolix"
	"main.go/internal/telegram"
	"main.go/internal/web_server"
	"main.go/pkg/config"
	"main.go/pkg/osutils"
	"runtime"
)

type InitPackage struct {
	Name     string
	InitFunc func() error
}

func main() {
	defer funcEnd()
	osutils.CallFuncByInterrupt(funcEnd)

	listInitFunc := []InitPackage{
		{"env", config.Init},
		{"telegram", telegram.Init},
		{"download_client", download_clients.Init},
		{"WebServer", web_server.Init},
	}

	for _, initFunc := range listInitFunc {
		err := initFunc.InitFunc()
		if err != nil {
			log.Fatal().Err(err).Msg(initFunc.Name)
			return
		}
	}
	defer download_clients.DefaultAllClients.Close()

	// torrent client
	Client1, err := torrent_anacrolix.New(download_clients.DefaultAllClients.PathContent)
	if err != nil {
		log.Fatal().Err(err).Msg("torrent_anacrolix.New")
	}
	download_clients.DefaultAllClients.AddClient(Client1)

	log.Info().Str("PATH_TORRENT_CONTENT", download_clients.DefaultAllClients.GetPathContent()).Interface("Names bot", telegram.GetListNameBot()).Msg("Start bots")

	telegram.SendMessageAdmin("Start bots \n" + telegram.GetInfo())

	telegram.Listener()
	log.Info().Msg("End.")
}

func funcEnd() {
	totalText := "Close app host " + web_server.WebServiceDefault.GetRooturl()
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
