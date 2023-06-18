package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client"
	"main.go/pkg/list_context"
	"main.go/pkg/osutils"
)

var listContext = list_context.NewListContext()

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

		if update.CallbackQuery != nil {
			runCallBack(bot, &update)
			return
		}

		if update.Message == nil {
			fmt.Println(update)
			continue
		}

		cmd := update.Message.Command()
		log.Info().Int("id", update.UpdateID).Msg("cmd:" + cmd)
		switch cmd {
		case "info":
			if contains(&adminUsersList, update.Message.Chat.ID) {
				_, _ = Send(bot, update.Message.Chat.ID, osutils.InfoHost(), nil)
				continue
			}
			fallthrough
		case "about":
			fallthrough
		case "start":
			_, _ = Send(bot, update.Message.Chat.ID, &textAbout, nil)
			continue
		default:

			err := torrent_client.CheckUrl(&update.Message.Text)
			if err != nil {
				log.Error().Err(err).Int64("chatId", update.Message.Chat.ID).Msg(`bot.Send`)
				_, _ = Send(bot, update.Message.Chat.ID, err, nil)
				continue
			}

			go serveTorrent(bot, update.Message.Chat.ID, &update.Message.Text)
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
