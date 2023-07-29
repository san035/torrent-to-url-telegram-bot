package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"main.go/internal/telegram"
	"main.go/internal/torrent_client"
	"main.go/internal/torrent_client/torrent_anacrolix"
	"main.go/internal/torrent_client/torrent_mock"
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
		{"torrent_client", torrent_client.Init},
	}

	if runtime.GOOS == "windows" {
		listInitFunc = append(listInitFunc, InitPackage{"torrent_mock", torrent_mock.Init})
	} else {
		listInitFunc = append(listInitFunc, InitPackage{"WebServer", web_server.Init})
		listInitFunc = append(listInitFunc, InitPackage{"torrent_anacrolix", torrent_anacrolix.Init})
	}

	for _, initFunc := range listInitFunc {
		err := initFunc.InitFunc()
		if err != nil {
			log.Fatal().Err(err).Msg(initFunc.Name)
			return
		}
	}
	defer torrent_client.DefaultClient.Close()

	log.Info().Str("PATH_TORRENT_CONTENT", torrent_client.GetPathTorrentContent()).Interface("Names bot", telegram.GetListNameBot()).Msg("Start bots")

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
