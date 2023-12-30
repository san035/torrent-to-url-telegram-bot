package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// GetInlineButton Создаем Callback кнопку c именем NameButton и данными: *NameButton + "_"+*data
func GetInlineButton(NameButton, data *string) *tgbotapi.InlineKeyboardMarkup {
	dataB := *NameButton + "_" + *data
	inlineBtn := tgbotapi.NewInlineKeyboardButtonData(*NameButton, dataB)
	// Создаем массив кнопок и добавляем в него Callback кнопку
	buttons := []tgbotapi.InlineKeyboardButton{inlineBtn}

	// Создаем объект клавиатуры и добавляем в нее массив кнопок
	b := tgbotapi.NewInlineKeyboardMarkup(buttons)
	return &b
}
