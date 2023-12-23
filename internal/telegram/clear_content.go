package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/download_clients"
)

func clearContent(bot *tgbotapi.BotAPI, id int64) {
	counDelete, err := download_clients.DefaultAllClients.RemoveAllContents()
	if err != nil {
		_, _ = Send(bot, id, err, nil)
		return
	}
	_, _ = Send(bot, id, fmt.Sprintf("Удалено %d файлов/каталогов", counDelete), nil)

	//info(bot, id)
	return
}
