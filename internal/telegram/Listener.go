package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"main.go/internal/download_clients"
	"slices"
	"strconv"
)

// Listener обработка входящих сообщений телеграмма
func (botsTelegram *BotsTelegram) Listener(allDownloadClient *download_clients.AllDownloadClient) {
	for i := 1; i < len(botsTelegram.listBot); i++ {
		go botsTelegram.ListenerOneBot(botsTelegram.listBot[i])
	}

	botsTelegram.ListenerOneBot(botsTelegram.listBot[0])
}

func (botsTelegram *BotsTelegram) ListenerOneBot(bot *tgbotapi.BotAPI) {

	slog.Info("Start bot", "Name", bot.Self.String())
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		SaveUser(update.Message.Chat)

		if update.CallbackQuery != nil {
			botsTelegram.runCallBack(bot, &update)
			return
		}

		if update.Message == nil {
			fmt.Println(update)
			continue
		}

		textCmd := update.Message.Command()
		slog.Info("cmd:"+textCmd, "id", update.UpdateID)

		cmd, ok := botsTelegram.MapCmd[textCmd]
		if ok {
			// Проверка на админа
			if cmd.IsAdmin && slices.Contains(botsTelegram.adminUsersList, update.Message.Chat.ID) {
				_, _ = Send(bot, update.Message.Chat.ID, "You are not admin, "+strconv.FormatInt(update.Message.Chat.ID, 10), nil)
				continue
			}

			cmd.DoFunc(bot, update.Message.Chat.ID)
			continue
		}

		// обработка сообщений по умолчанию
		botsTelegram.DoFuncDefault(bot, update.Message.Chat.ID, &update.Message.Text)

	}
	slog.Info("End bot", "Name", bot.Self.String())
}
