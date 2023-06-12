package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

var (
	Bot     *tgbotapi.BotAPI
	updates *tgbotapi.UpdatesChannel
)

func Init() (err error) {
	botToken := os.Getenv("BOT_TOKEN")
	Bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates1 := Bot.GetUpdatesChan(u)
	updates = &updates1
	return nil
}
