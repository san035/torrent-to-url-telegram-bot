package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

// Send, text - string or *string
func Send(chatId int64, text interface{}, button *tgbotapi.InlineKeyboardMarkup) (mes tgbotapi.Message, err error) {
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
		log.Error().Err(err).Int64("chatId", chatId).Msg(`telegram.Send`)
		return
	}
	return
}
