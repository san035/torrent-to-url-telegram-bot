package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"log/slog"
	"main.go/pkg/list_context"
	"os"
	"strconv"
	"strings"
)

// Используемые имена env
const (
	ListAdminIdTelegram  = "TELEGRAM_LIST_ADMIN_ID"
	ListBotToken         = "TELEGRAM_LIST_BOT_TOKEN"
	BotAbout             = "TELEGRAM_BOT_ABOUT"
	SendMessageUpService = "TELEGRAM_SEND_MESSAGE_UP_SERVICE"
)

type DoFunc func(*tgbotapi.BotAPI, int64)
type DoFuncDefault func(*tgbotapi.BotAPI, int64, *string)

type CmdTg struct {
	IsAdmin     bool
	Description string
	DoFunc      DoFunc
}

type MapCmd map[string]CmdTg

type BotsTelegram struct {
	listBot        []*tgbotapi.BotAPI
	adminUsersList []int64
	textAbout      string
	//pathContent    string
	listNameBot   []string
	MapCmd        MapCmd
	DoFuncDefault DoFuncDefault // функция для обработки сообщений не попадающих в MapCmd
	ListContext   *list_context.ListContext
}

var (
	ButtonCansel = `cansel`
	ButtonDelete = `delete`
)

func New() (botsTelegram *BotsTelegram, err error) {

	listToken := strings.Split(os.Getenv(ListBotToken), ",")

	listBot := make([]*tgbotapi.BotAPI, 0, len(listToken))
	listNameBot := make([]string, 0, len(listToken))
	for i := 0; i < len(listToken); i++ {
		token := listToken[i]
		bot, err2 := tgbotapi.NewBotAPI(token)
		if err2 != nil {
			listNameBot = append(listNameBot, token+"-"+err.Error())
			slog.Error("bad token", "error", err2, "token", token)
			continue
		}
		listBot = append(listBot, bot)
		listNameBot = append(listNameBot, "@"+bot.Self.String())
	}
	if len(listBot) == 0 {
		err = errors.New("Not work token in env: " + ListBotToken)
		return
	}

	botsTelegram = &BotsTelegram{
		listBot:        listBot,
		listNameBot:    listNameBot,
		adminUsersList: SplitStringToInt64(os.Getenv(ListAdminIdTelegram)),
		textAbout:      os.Getenv(BotAbout),
		//pathContent:    pathContent,
		ListContext: list_context.NewListContext(),
	}

	botsTelegram.MapCmd = MapCmd{
		"down":       {IsAdmin: true, DoFunc: botsTelegram.down},
		"start":      {IsAdmin: false, DoFunc: botsTelegram.start},
		"create_cmd": {IsAdmin: true, DoFunc: botsTelegram.creatCmd, Description: "Создать список команд для @BotFather"},
	}
	return
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

func (botsTelegram *BotsTelegram) GetListNameBot() []string {
	return botsTelegram.listNameBot
}
