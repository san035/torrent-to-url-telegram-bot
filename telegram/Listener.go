package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/torrent_client"
	"main.go/web_server"
	"os"
)

// Listener обработка входящих сообщений телеграмма
func Listener() {
	for update := range *updates {

		if update.Message == nil { // ignore non-message updates
			continue
		}

		cmd := update.Message.Command()
		log.Info().Int("id", update.UpdateID).Msg("cmd:" + cmd)
		switch cmd {
		case "start":
			_, err := Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, os.Getenv("BOT_ABOUT")))
			if err != nil {
				log.Error().Err(err).Int64("chatId", update.Message.Chat.ID).Msg(`Bot.Send`)
			}
			continue
		default:

			ctx, cancel := context.WithCancel(context.Background())

			go func(ctx *context.Context, cancel context.CancelFunc, chatId int64) {
				chanStatus, err := torrent_client.GetChanMessage(ctx, &update.Message.Text)
				if err != nil {
					_, err2 := Bot.Send(tgbotapi.NewMessage(chatId, err.Error()))
					if err2 != nil {
						log.Err(err).Err(err2).Int64("chatId", chatId).Msg("Error Bot.Send")
					}
					cancel()
					return
				}
				firstMessage, err := Bot.Send(tgbotapi.NewMessage(chatId, "Start"))
				if err != nil {
					log.Err(err).Int64("chatId", chatId).Msg("Error Bot.Send")
					cancel()
					return
				}

				var textMsg string
				for status := range *chanStatus {
					textMsg = status.Info
					if status.FileName != nil {
						url := web_server.GetUrl(status.FileName)
						textMsg += "\n" + url
					}
					msg := tgbotapi.NewEditMessageText(chatId, firstMessage.MessageID, textMsg)
					if _, err = Bot.Send(msg); err != nil {
						log.Err(err).Int64("chatId", chatId).Msg("Error Bot.Send")
						cancel()
						return
					}
				}
			}(&ctx, cancel, update.Message.Chat.ID)

		}

	}

}
