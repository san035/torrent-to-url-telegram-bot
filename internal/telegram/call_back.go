package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// runCallBack Обработка кнопок inline
func (botsTelegram *BotsTelegram) runCallBack(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	callbackDataStr := update.CallbackData()
	callbackData := strings.Split(callbackDataStr, "_")
	if len(callbackData) < 2 {
		slog.Error("данные callBack не соответсвуют формату *_*", "callbackData", callbackDataStr, "textMsg", update.CallbackQuery.Message.Text)
		return
	}

	idContext, err := strconv.ParseUint(callbackData[1], 10, 64)
	if err != nil {
		slog.Error("Невозможно конверитровать id в callBack", "error", err, "callbackData", callbackDataStr, "textMsg", update.CallbackQuery.Message.Text)
	}

	// завершаем контекст
	dataContext := botsTelegram.ListContext.GetById(idContext)
	if dataContext != nil {
		select {
		case <-(*dataContext.Context).Done():
		default:
			// Закрываем контекст
			cansel := (*dataContext.Context).Value("cansel")
			if cansel != nil {
				canselFunc, ok := cansel.(func())
				if ok {
					canselFunc()

					// todo лучше дождаться завершения контекста
					time.Sleep(time.Second)
				}
			}
		}
	}

	if callbackData[0] != ButtonDelete {
		return
	}

	// Удаляем данные
	fileOrFolderName := path.Base(update.CallbackQuery.Message.Text)
	err = os.Remove(fileOrFolderName)
	if err != nil {
		slog.Error("Ошибка удаления", "error", err, "file", fileOrFolderName)
		return
	}

	// todo
	deleteMsg := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	_, err = bot.Send(deleteMsg)
	if err != nil {
		slog.Error("Ошибка удаления сообщения", "error", err, "text", update.CallbackQuery.Message.Text)
	}
	return
}
