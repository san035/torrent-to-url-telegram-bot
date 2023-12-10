package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/download_clients"
	"main.go/internal/web_server"
	"main.go/pkg/list_context"
	"strconv"
)

var (
	buttonCansel = `cansel`
	buttonDelete = `delete`
)

func serveTorrent(bot *tgbotapi.BotAPI, chatId int64, clientDownload download_clients.DownloadClient, magnetUrl *string) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "cancel", cancel)

	itemContext := &list_context.DataContext{Context: &ctx, MagnetUrl: magnetUrl}
	idContext, err := listContext.Add(itemContext)
	if err != nil {
		_, _ = Send(bot, chatId, "torrent busy", nil)
		cancel()
		return
	}

	dataMsg := strconv.FormatUint(idContext, 10)
	mapButton := map[int]*tgbotapi.InlineKeyboardMarkup{
		download_clients.StatusTorrentStart: GetInlineButton(&buttonCansel, &dataMsg),
		download_clients.StatusTorrentRun:   GetInlineButton(&buttonCansel, &dataMsg),
		download_clients.StatusTorrentPause: GetInlineButton(&buttonDelete, &dataMsg),
		download_clients.StatusTorrentEnd:   GetInlineButton(&buttonDelete, &dataMsg),
	}

	chanStatus, err := clientDownload.StartDownload(&ctx, magnetUrl)
	if err != nil {
		_, _ = Send(bot, chatId, err, mapButton[download_clients.StatusTorrentEnd])
		cancel()
		return
	}

	firstMessage, err := Send(bot, chatId, "Start", mapButton[download_clients.StatusTorrentStart])
	if err != nil {
		log.Err(err).Int64("chatId", chatId).Msg("Error bot.Send")
		cancel()
		return
	}

	var textMsg string
	// Слушаем сообщения торрента
	for status := range *chanStatus {
		textMsg = status.Info
		if status.WebFileName != nil {
			textMsg += web_server.GetUrl(status.WebFileName)
		}

		log.Debug().Str("bot", bot.Self.String()).Str("user", NikNameById(chatId)).Str("text", textMsg).Msg("Сообщение от торрента")

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
}
