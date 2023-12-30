package core

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"main.go/internal/download_clients"
	"main.go/internal/telegram"
	"main.go/internal/web_server"
)

type Core struct {
	botsTG             *telegram.BotsTelegram
	allDownloadClients *download_clients.AllDownloadClient
	webServer          *web_server.HttpService
}

func New(botsTG *telegram.BotsTelegram, allDownloadClients *download_clients.AllDownloadClient, webServer *web_server.HttpService) *Core {

	core := &Core{
		botsTG:             botsTG,
		allDownloadClients: allDownloadClients,
		webServer:          webServer,
	}

	// доп. команды телеграм которыем исполняются вне модуля telegram
	addCmdTG := telegram.MapCmd{
		"show_content":  {IsAdmin: true, DoFunc: core.ShowContent},
		"clear_content": {IsAdmin: true, DoFunc: core.ClearContent},
		"info":          {IsAdmin: true, DoFunc: core.Info},
	}
	botsTG.AddCommand(addCmdTG)
	botsTG.DoFuncDefault = core.DoFuncDefault

	return core
}

func (core *Core) DoFuncDefault(bot *tgbotapi.BotAPI, chatID int64, testMsg *string) {
	// в сообщении ссылка на скачивание
	// ищем какому клиенту загруки она подходит
	var startWork bool
	for _, clientDownload := range core.allDownloadClients.ListClient {
		if clientDownload.GoodUrl(testMsg) {
			startWork = true
			go core.serveTorrent(bot, chatID, clientDownload, testMsg)
			break
		}
	}

	if !startWork {
		err := errors.New(`unsupport URL`)
		slog.Info(`bot.Send`, "Error", err, "chatId", chatID)
		textMsg := fmt.Sprintf("%v support mask: %v", err, core.allDownloadClients.GetListSuppotUrlPattern())
		_, _ = telegram.Send(bot, chatID, textMsg, nil)
	}
}

func (core *Core) ShowContent(bot *tgbotapi.BotAPI, id int64) {
	//listContent, err := torrent_client.GetListDownloadContent()
	//if err != nil {
	//	_, _ = Send(bot, id, err, nil)
	//	return
	//}
	//
	//var textMsg string
	//for _, fileName := range listContent {
	//	textMsg += web_server.GetUrl(&fileName)
	//}

	_, _ = telegram.Send(bot, id, core.webServer.GetRooturl(), nil)
	return
}
