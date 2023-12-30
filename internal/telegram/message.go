package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

// Send, text - string or *string
func Send(Bot *tgbotapi.BotAPI, chatId int64, text interface{}, button *tgbotapi.InlineKeyboardMarkup) (mes tgbotapi.Message, err error) {
	var textSend string

	switch text.(type) {
	case string:
		textSend = text.(string)
	case *string:
		textSend = *text.(*string)
	default:
		textSend = fmt.Sprint(text)
	}

	mesConfig := tgbotapi.NewMessage(chatId, textSend)
	mesConfig.BaseChat.ReplyMarkup = button
	mes, err = Bot.Send(mesConfig)
	if err != nil {
		slog.Error(`telegram.Send`, "error", err, "chatId", chatId)
		return
	}
	return
}

// SendMessageAdmin Отправка админам
func (botsTelegram *BotsTelegram) SendMessageAdmin(text interface{}) {
	for _, chatIdAdmin := range botsTelegram.adminUsersList {
		_, err := Send(botsTelegram.listBot[0], chatIdAdmin, text, nil)
		if err != nil {
			slog.Error("telegram.SendMessageAdmin-", "error", err, "text", text)
			break
		}
	}
}
