package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func start(bot *tgbotapi.BotAPI, id int64) {
	_, _ = Send(bot, id, &textAbout, nil)

	return
}
