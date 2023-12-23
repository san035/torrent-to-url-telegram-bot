package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/web_server"
)

func showContent(bot *tgbotapi.BotAPI, id int64) {
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

	_, _ = Send(bot, id, web_server.WebServiceDefault.GetRooturl(), nil)
	return
}
