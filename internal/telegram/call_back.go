package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// runCallBack Обработка кнопок inline
func runCallBack(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	callbackDataStr := update.CallbackData()
	callbackData := strings.Split(callbackDataStr, "_")
	if len(callbackData) < 2 {
		log.Error().Err(err).Str("callbackData", callbackDataStr).Str("textMsg", update.CallbackQuery.Message.Text).Msg("данные callBack не соответсвуют формату *_*")
		return
	}

	idContext, err := strconv.ParseUint(callbackData[1], 10, 64)
	if err != nil {
		log.Error().Err(err).Str("callbackData", callbackDataStr).Str("textMsg", update.CallbackQuery.Message.Text).Msg("Невозможно конверитровать id в callBack")
	}

	// завершаем контекст
	dataContext := listContext.GetById(idContext)
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

	if callbackData[0] != buttonDelete {
		return
	}

	// Удаляем данные
	fileOrFolderName := path.Base(update.CallbackQuery.Message.Text)
	err = os.Remove(fileOrFolderName)
	if err != nil {
		log.Error().Err(err).Str("file", fileOrFolderName).Msg("Ошибка удаления")
		return
	}

	// todo
	deleteMsg := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	_, err = bot.Send(deleteMsg)
	if err != nil {
		log.Error().Err(err).Str("text", update.CallbackQuery.Message.Text).Msg("Ошибка удаления сообщения")
	}
	return
}
