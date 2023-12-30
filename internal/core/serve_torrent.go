package core

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"main.go/internal/download_clients"
	"main.go/internal/telegram"
	"main.go/pkg/list_context"
	"strconv"
)

func (core *Core) serveTorrent(bot *tgbotapi.BotAPI, chatId int64, clientDownload download_clients.DownloadClient, magnetUrl *string) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "cancel", cancel)

	itemContext := &list_context.DataContext{Context: &ctx, MagnetUrl: magnetUrl}
	idContext, err := core.botsTG.ListContext.Add(itemContext)
	if err != nil {
		_, _ = telegram.Send(bot, chatId, "torrent busy", nil)
		cancel()
		return
	}

	dataMsg := strconv.FormatUint(idContext, 10)
	mapButton := map[int]*tgbotapi.InlineKeyboardMarkup{
		download_clients.StatusTorrentStart: telegram.GetInlineButton(&telegram.ButtonCansel, &dataMsg),
		download_clients.StatusTorrentRun:   telegram.GetInlineButton(&telegram.ButtonCansel, &dataMsg),
		download_clients.StatusTorrentPause: telegram.GetInlineButton(&telegram.ButtonDelete, &dataMsg),
		download_clients.StatusTorrentEnd:   telegram.GetInlineButton(&telegram.ButtonDelete, &dataMsg),
	}

	chanStatus, err := clientDownload.StartDownload(&ctx, magnetUrl)
	if err != nil {
		_, _ = telegram.Send(bot, chatId, err, mapButton[download_clients.StatusTorrentEnd])
		cancel()
		return
	}

	firstMessage, err := telegram.Send(bot, chatId, "Start", mapButton[download_clients.StatusTorrentStart])
	if err != nil {
		slog.Error("bot.Send", "error", err, "chatId", chatId)
		cancel()
		return
	}

	var textMsg string
	// Слушаем сообщения торрента
	for status := range *chanStatus {
		textMsg = status.Info
		if status.WebFileName != nil {
			textMsg += core.webServer.GetUrl(status.WebFileName)
		}

		slog.Debug("Сообщение от торрента", "bot", bot.Self.String(), "user", telegram.NikNameById(chatId), "text", textMsg)

		msg := tgbotapi.NewEditMessageText(chatId, firstMessage.MessageID, textMsg)
		msg.BaseEdit.ReplyMarkup = mapButton[status.Status]
		if _, err = bot.Send(msg); err != nil {
			slog.Error("Error bot.Send", "error", err, "chatId", chatId)
			cancel()
			return
		}

	}
	cancel()
	core.botsTG.ListContext.Delete(itemContext)
}
