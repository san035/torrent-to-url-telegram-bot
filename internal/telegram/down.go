package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

func (botsTelegram *BotsTelegram) down(bot *tgbotapi.BotAPI, id int64) {
	for _, b := range botsTelegram.listBot {
		u, err := b.GetMe()
		if err != nil {
			slog.Error("telegram.down-", "error", err)
		}

		_, _ = Send(bot, id, "Завершение работы "+u.String(), nil)

		b.StopReceivingUpdates()
	}

	return
}
