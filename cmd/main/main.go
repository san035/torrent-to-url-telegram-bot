package main

import (
	"fmt"
	"log/slog"
	"main.go/internal/config"
	"main.go/internal/core"
	"main.go/internal/download_clients"
	"main.go/internal/download_clients/torrent_anacrolix"
	"main.go/internal/telegram"
	"main.go/internal/web_server"
	"main.go/pkg/osutils"
	"os"
	"runtime"
)

func main() {
	// Установка уровня логирования
	setupLogger(slog.LevelInfo.String())
	slog.Info("Start", "app", os.Args[0])

	// Чтение конфиг
	err := config.Load()
	if err != nil {
		slog.Error("config.Load", "error", err)
		panic(1)
		return
	}

	// Cозданние сущности всех клиенетов загрузки
	allDownloadClient, err := download_clients.New()
	if err != nil {
		slog.Error("Load download_clients", "err", err)
	}
	defer allDownloadClient.Close()

	// Cоздание телеграм ботов
	botsTG, err := telegram.New()
	if err != nil {
		slog.Error("Load telegram", "err", err)
		panic(1)
	}

	// Настройка прерываения приложения
	defer funcEnd(botsTG)
	osutils.CallFuncByInterrupt(func() {
		funcEnd(botsTG)
	})

	// torrent client
	Client1, err := torrent_anacrolix.New(allDownloadClient.PathContent)
	if err != nil {
		slog.Error("torrent_anacrolix.New", "err", err)
		panic(1)
	}
	allDownloadClient.AddClient(Client1)

	// web сревис
	webServer := web_server.New(allDownloadClient.GetPathContent())

	// ядро
	coreService := core.New(botsTG, allDownloadClient, webServer)

	// Запуск работы ядра
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
