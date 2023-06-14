package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client"
	"main.go/internal/web_server"
	"main.go/pkg/list_context"
)

// Listener обработка входящих сообщений телеграмма
func Listener() {
	listContext := list_context.NewListContext()
	for update := range *updates {

		if update.Message == nil { // ignore non-message updates
			fmt.Println(update)
			continue
		}

		cmd := update.Message.Command()
		log.Info().Int("id", update.UpdateID).Msg("cmd:" + cmd)
		switch cmd {
		case "start":
			_, _ = Send(update.Message.Chat.ID, &textAbout, nil)
			continue
		default:

			err := torrent_client.CheckUrl(&update.Message.Text)
			if err != nil {
				log.Error().Err(err).Int64("chatId", update.Message.Chat.ID).Msg(`Bot.Send`)
				_, _ = Send(update.Message.Chat.ID, err, nil)
				continue
			}

			go func(chatId int64, magnetUrl *string) {
				ctx, cancel := context.WithCancel(context.Background())

				itemContext := &list_context.DataContext{Context: &ctx, MagnetUrl: magnetUrl}
				idContext, err := listContext.AddContext(itemContext)
				if err != nil {
					_, _ = Send(update.Message.Chat.ID, "torrent busy", nil)
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
					_, _ = Send(update.Message.Chat.ID, err, mapButton[torrent_client.StatusTorrentEnd])
					cancel()
					return
				}

				firstMessage, err := Send(chatId, "Start", mapButton[torrent_client.StatusTorrentStart])
				if err != nil {
					log.Err(err).Int64("chatId", chatId).Msg("Error Bot.Send")
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
					if _, err = Bot.Send(msg); err != nil {
						log.Err(err).Int64("chatId", chatId).Msg("Error Bot.Send")
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
