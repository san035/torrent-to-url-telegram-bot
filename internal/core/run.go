package core

import (
	"fmt"
	"main.go/internal/telegram"
	"os"
)

func (core *Core) Run() {
	// Запуск слушания сообщений всех ботов
	core.botsTG.Listener(core.allDownloadClients)

	v, ok := os.LookupEnv(telegram.SendMessageUpService)
	if !ok || v != "false" {
		textForAdmin := fmt.Sprintf("Start bots %s", core.botsTG.GetListNameBot())
		core.botsTG.SendMessageAdmin(textForAdmin)
	}
}
