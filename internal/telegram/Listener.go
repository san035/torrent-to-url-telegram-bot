package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client"
	"main.go/internal/web_server"
	"main.go/pkg/list_context"
	"main.go/pkg/osutils"
	"runtime"
	"strconv"
)

var (
	listContext = list_context.NewListContext()
)

// Listener обработка входящих сообщений телеграмма
func Listener() {
	SendMessageAdmin("Start bots \nhost: " +
		*web_server.HostAndPort +
		"\n" + osutils.GetFreeHDD() +
		"\n" + osutils.GetFreeMem() +
		"\n Goroutines: " + strconv.Itoa(runtime.NumGoroutine()) +
		"\n bots: " + fmt.Sprint(listNameBot))

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

		if update.Message == nil { // ignore non-message updates
			fmt.Println(update)
			continue
		}

		cmd := update.Message.Command()
		log.Info().Int("id", update.UpdateID).Msg("cmd:" + cmd)
		switch cmd {
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

			go func(chatId int64, magnetUrl *string) {
				ctx, cancel := context.WithCancel(context.Background())

				itemContext := &list_context.DataContext{Context: &ctx, MagnetUrl: magnetUrl}
				idContext, err := listContext.AddContext(itemContext)
				if err != nil {
					_, _ = Send(bot, update.Message.Chat.ID, "torrent busy", nil)
					cancel()
					return
				}

				mapButton := map[int]*tgbotapi.InlineKeyboardMarkup{
					torrent_client.StatusTorrentStart: GetInlineButton(`cansel`, idContext),
					torrent_client.StatusTorrentRun:   GetInlineButton(`cansel`, idContext),
					torrent_client.StatusTorrentPause: GetInlineButton(`delete`, idContext),
					torrent_client.StatusTorrentEnd:   GetInlineButton(`delete`, idContext),
				}

				chanStatus, err := torrent_client.GetChanMessage(&ctx, magnetUrl)
				if err != nil {
					_, _ = Send(bot, update.Message.Chat.ID, err, mapButton[torrent_client.StatusTorrentEnd])
					cancel()
					return
				}

				firstMessage, err := Send(bot, chatId, "Start", mapButton[torrent_client.StatusTorrentStart])
				if err != nil {
					log.Err(err).Int64("chatId", chatId).Msg("Error bot.Send")
					cancel()
					return
				}

				var textMsg string
				for status := range *chanStatus {
					textMsg = status.Info
					if status.WebFileName != nil {
						textMsg += web_server.GetUrl(status.WebFileName)
					}

					msg := tgbotapi.NewEditMessageText(chatId, firstMessage.MessageID, textMsg)
					msg.BaseEdit.ReplyMarkup = mapButton[status.Status]
					if _, err = bot.Send(msg); err != nil {
						log.Err(err).Int64("chatId", chatId).Msg("Error bot.Send")
						cancel()
						return
					}

				}
				cancel()
				listContext.Delete(itemContext)
			}(update.Message.Chat.ID, &update.Message.Text)

		}

	}

}
