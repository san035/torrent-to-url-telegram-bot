package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (botsTelegram *BotsTelegram) start(bot *tgbotapi.BotAPI, id int64) {
	textAboutWithId := botsTelegram.textAbout + "\n you id: " + strconv.FormatInt(id, 10)
	_, _ = Send(bot, id, &textAboutWithId, nil)

	return
}
