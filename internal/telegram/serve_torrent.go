package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client"
	"main.go/internal/web_server"
	"main.go/pkg/list_context"
	"strconv"
)

var (
	buttonCansel = `cansel`
	buttonDelete = `delete`
)

func serveTorrent(bot *tgbotapi.BotAPI, chatId int64, magnetUrl *string) {
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
		torrent_client.StatusTorrentStart: GetInlineButton(&buttonCansel, &dataMsg),
		torrent_client.StatusTorrentRun:   GetInlineButton(&buttonCansel, &dataMsg),
		torrent_client.StatusTorrentPause: GetInlineButton(&buttonDelete, &dataMsg),
		torrent_client.StatusTorrentEnd:   GetInlineButton(&buttonDelete, &dataMsg),
	}

	chanStatus, err := torrent_client.DefaultClient.StartTorrent(&ctx, magnetUrl)
	if err != nil {
		_, _ = Send(bot, chatId, err, mapButton[torrent_client.StatusTorrentEnd])
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
	// Слушаем сообщения торрента
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
}
