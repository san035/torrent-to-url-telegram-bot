package core

import (
	"fmt"
	"log/slog"
	"main.go/internal/telegram"
	"os"
)

func (core *Core) Run() {
	// Запуск слушания сообщений всех ботов
	slog.Info("Start bots", "PATH_TORRENT_CONTENT", core.allDownloadClients.GetPathContent(), "Names bot", core.botsTG.GetListNameBot())
	core.botsTG.Listener(core.allDownloadClients)

	v, ok := os.LookupEnv(telegram.SendMessageUpService)
	if !ok || v != "false" {
		textForAdmin := fmt.Sprintf("Start bots %s", core.botsTG.GetListNameBot())
		core.botsTG.SendMessageAdmin(textForAdmin)
	}
}
