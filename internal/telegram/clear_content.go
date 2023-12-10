package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/download_clients"
)

func clearContent(bot *tgbotapi.BotAPI, id int64) {
	err := download_clients.DefaultAllClients.RemoveAllContents()
	if err != nil {
		_, _ = Send(bot, id, err, nil)
		return
	}

	info(bot, id)
	return
}
