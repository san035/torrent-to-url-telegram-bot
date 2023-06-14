package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/web_server"
	"main.go/pkg/config"
	"main.go/pkg/osutils"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	Bot            *tgbotapi.BotAPI
	updates        *tgbotapi.UpdatesChannel
	textAbout      string
	adminUsersList []int64
	err            error
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

func init() {
	err = config.Init()
	if err != nil {
		return
	}

	botToken := os.Getenv("BOT_TOKEN")
	Bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates1 := Bot.GetUpdatesChan(u)
	updates = &updates1

	textAbout = os.Getenv("BOT_ABOUT")

	envList := strings.Split(os.Getenv("LIST_ADMIN_ID_TELEGRAM"), ",")
	for _, id := range envList {
		idAdmin, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error().Err(err).Str("id", id).Msg("Load env LIST_ADMIN_ID_TELEGRAM-")
			continue
		}
		adminUsersList = append(adminUsersList, idAdmin)
	}
	SendMessageAdmin("Start host: " +
		*web_server.HostAndPort +
		"\n" + osutils.GetFreeHDD() +
		"\n" + osutils.GetFreeMem() +
		"\n Goroutines: " + strconv.Itoa(runtime.NumGoroutine()))
	return
}

func Init() error {
	return err
}
