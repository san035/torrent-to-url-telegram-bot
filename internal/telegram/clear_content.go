package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/torrent_client"
)

func clearContent(bot *tgbotapi.BotAPI, id int64) {
	err := torrent_client.RemoveAllContents()
	if err != nil {
		_, _ = Send(bot, id, err, nil)
		return
	}

	info(bot, id)
	return
}
