package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func creatCmd(bot *tgbotapi.BotAPI, id int64) {

	myOldCmd, err := bot.GetMyCommands()
	if err != nil {
		_, _ = Send(bot, id, err, nil)
	}
	var allOldCmd string
	for _, cmd := range myOldCmd {
		allOldCmd += fmt.Sprintf("%s - %s\n", cmd.Command, cmd.Description)
	}

	var allNewCmd string
	for cmdName, v := range MapCmd {
		if v.Description == "" {
			v.Description = cmdName
		}
		allNewCmd += fmt.Sprintf("%s - %s\n", cmdName, v.Description)
	}

	//todo сравненеие не работает, из за map
	if allOldCmd == allNewCmd {
		_, _ = Send(bot, id, "Команды не изменились", nil)
		return
	}

	_, _ = Send(bot, id, "Настройка нового меню в @BotFather \n"+allNewCmd, nil)
	return
}
