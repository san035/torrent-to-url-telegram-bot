package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

var (
	Bot       *tgbotapi.BotAPI
	updates   *tgbotapi.UpdatesChannel
	textAbout string
)

func GetInlineButton(NameButton string, data int) *tgbotapi.InlineKeyboardMarkup {
	// Создаем Callback кнопку "Cancel"
	inlineBtn := tgbotapi.NewInlineKeyboardButtonData(NameButton, strconv.Itoa(data))

	// Создаем массив кнопок и добавляем в него Callback кнопку
	buttons := []tgbotapi.InlineKeyboardButton{inlineBtn}

	// Создаем объект клавиатуры и добавляем в нее массив кнопок
	b := tgbotapi.NewInlineKeyboardMarkup(buttons)
	return &b
}

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

	textAbout = os.Getenv("BOT_ABOUT")
	return nil
}
