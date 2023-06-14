package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/web_server"
	"os"
	"strconv"
	"strings"
)

var (
	Bot            *tgbotapi.BotAPI
	updates        *tgbotapi.UpdatesChannel
	textAbout      string
	adminUsersList []int64
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

	envList := strings.Split(os.Getenv("LIST_ADMIN_ID_TELEGRAM"), ",")
	for _, id := range envList {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error().Err(err).Str("id", id).Msg("Load env LIST_ADMIN_ID_TELEGRAM-")
			continue
		}
		adminUsersList = append(adminUsersList, i)
	}
	SendMessageAdmin("Start host: " + *web_server.HostAndPort)
	return nil
}
