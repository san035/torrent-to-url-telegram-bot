package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"log/slog"
	"main.go/internal/core"
	"main.go/internal/download_clients"
	"main.go/internal/download_clients/torrent_anacrolix"
	"main.go/internal/telegram"
	"main.go/internal/web_server"
	"main.go/pkg/config"
	"main.go/pkg/osutils"
	"os"
	"runtime"
)

type InitPackage struct {
	Name     string
	InitFunc func() error
}

func main() {
	// Установка уровня логирования
	setupLogger(slog.LevelInfo.String())
	slog.Info("Start", "app", os.Args[0])

	listInitFunc := []InitPackage{
		{"env", config.Init},
	}

	for _, initFunc := range listInitFunc {
		err := initFunc.InitFunc()
		if err != nil {
			log.Fatal().Err(err).Msg(initFunc.Name)
			return
		}
	}

	// Cозданние сущности всех клиенетов загрузки
	allDownloadClient, err := download_clients.New()
	if err != nil {
		slog.Error("Init download_clients", "err", err)
	}
	defer allDownloadClient.Close()

	// Cоздание телеграм ботов
	botsTG, err := telegram.New()
	if err != nil {
		slog.Error("Init telegram", "err", err)
		panic(1)
	}

	defer funcEnd(botsTG)
	osutils.CallFuncByInterrupt(func() {
		funcEnd(botsTG)
	})

	// torrent client
	Client1, err := torrent_anacrolix.New(allDownloadClient.PathContent)
	if err != nil {
		log.Fatal().Err(err).Msg("torrent_anacrolix.New")
	}
	allDownloadClient.AddClient(Client1)

	webServer := web_server.New(allDownloadClient.GetPathContent())

	coreService := core.New(botsTG, allDownloadClient, webServer)

	slog.Info("Start bots", "PATH_TORRENT_CONTENT", allDownloadClient.GetPathContent(), "Names bot", botsTG.GetListNameBot())

	coreService.Run()

	slog.Info("End.")
}

func funcEnd(botsTelegram *telegram.BotsTelegram) {
	totalText := "Close app"
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
	slog.Info("Close app", "recover", totalText)
	if botsTelegram != nil {
		botsTelegram.SendMessageAdmin(totalText)
	}
}

func setupLogger(logLevel string) {
	var programLevel = new(slog.Level)
	err := programLevel.UnmarshalText([]byte(logLevel))
	if err != nil {
		slog.Error("Error set log_level", "log_level", logLevel)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	slog.SetDefault(logger)
	slog.Debug("", "log_level", logLevel)
}
