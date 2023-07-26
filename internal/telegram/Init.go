package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/pkg/config"
	"os"
	"strconv"
	"strings"
)

type cmdTg struct {
	IsAdmin     bool
	Description string
	DoFunc      func(*tgbotapi.BotAPI, int64)
}

type mapCmd map[string]cmdTg

var (
	listBot []*tgbotapi.BotAPI
	//updates        *tgbotapi.UpdatesChannel
	textAbout      string
	adminUsersList []int64
	err            error
	listNameBot    []string

	MapCmd = mapCmd{
		"info":          {IsAdmin: true, DoFunc: info},
		"down":          {IsAdmin: true, DoFunc: down},
		"start":         {IsAdmin: false, DoFunc: start},
		"clear_content": {IsAdmin: true, DoFunc: clearContent},
		// create_cmd - ниже
	}
)

func GetListNameBot() []string {
	return listNameBot
}

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

func init() {
	err = config.Init()
	if err != nil {
		return
	}

	adminUsersList = SplitStringToInt64(os.Getenv("LIST_ADMIN_ID_TELEGRAM"))

	listToken := strings.Split(os.Getenv("LIST_BOT_TOKEN"), ",")
	listBot = make([]*tgbotapi.BotAPI, 0, len(listToken))
	listNameBot = make([]string, 0, len(listToken))
	for i := 0; i < len(listToken); i++ {
		token := listToken[i]
		bot, err2 := tgbotapi.NewBotAPI(token)
		if err2 != nil {
			listNameBot = append(listNameBot, token+"-"+err.Error())
			log.Error().Err(err2).Str("token", token).Msg("bad token")
		}
		listBot = append(listBot, bot)
		listNameBot = append(listNameBot, bot.Self.String())
	}
	if len(listBot) == 0 {
		err = errors.New("Not work token in env:LIST_BOT_TOKEN")
		return
	}

	textAbout = os.Getenv("BOT_ABOUT")

	MapCmd["create_cmd"] = cmdTg{IsAdmin: true, DoFunc: creatCmd, Description: "Создать список команд для @BotFather"}

	return
}

func Init() error {
	return err
}

func SplitStringToInt64(envList string) (adminUsersList []int64) {
	adminUsersListStr := strings.Split(envList, ",")
	adminUsersList = make([]int64, 0, len(adminUsersListStr))
	for _, id := range adminUsersListStr {
		idAdmin, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error().Err(err).Str("id", id).Msg("Load env LIST_ADMIN_ID_TELEGRAM-")
			continue
		}
		adminUsersList = append(adminUsersList, idAdmin)
	}
	return
}
