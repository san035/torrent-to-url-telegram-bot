package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/download_clients"
	"main.go/pkg/list_context"
	"strconv"
)

var (
	listContext = list_context.NewListContext()
)

// Listener обработка входящих сообщений телеграмма
func Listener() {
	for i := 1; i < len(listBot); i++ {
		go ListenerOneBot(listBot[i])
	}

	ListenerOneBot(listBot[0])
}

func ListenerOneBot(bot *tgbotapi.BotAPI) {

	log.Info().Str("Name", bot.Self.String()).Msg("Start bot")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		SaveUser(update.Message.Chat)

		if update.CallbackQuery != nil {
			runCallBack(bot, &update)
			return
		}

		if update.Message == nil {
			fmt.Println(update)
			continue
		}

		textCmd := update.Message.Command()
		log.Info().Int("id", update.UpdateID).Msg("cmd:" + textCmd)

		cmd, ok := MapCmd[textCmd]
		if ok {
			// Проверка на админа
			if cmd.IsAdmin && !contains(&adminUsersList, update.Message.Chat.ID) {
				_, _ = Send(bot, update.Message.Chat.ID, "You are not admin, "+strconv.FormatInt(update.Message.Chat.ID, 10), nil)
				continue
			}

			cmd.DoFunc(bot, update.Message.Chat.ID)
			continue
		}

		// в сообщении ссылка на скачивание
		var startWork bool
		for _, client := range download_clients.DefaultAllClients.ListClient {
			if client.GoodUrl(&update.Message.Text) {
				startWork = true
				go serveTorrent(bot, update.Message.Chat.ID, client, &update.Message.Text)
				break
			}
		}

		if !startWork {
			err = errors.New(`the URL must start with "magnet:" `)
			log.Error().Err(err).Int64("chatId", update.Message.Chat.ID).Msg(`bot.Send`)
			_, _ = Send(bot, update.Message.Chat.ID, err, nil)
			continue
		}

	}

}

func contains(arr *[]int64, num int64) bool {
	for _, n := range *arr {
		if n == num {
			return true
		}
	}
	return false
}
