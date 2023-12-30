package core

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"main.go/internal/file_func"
	"main.go/internal/telegram"
	"main.go/pkg/osutils"
	"runtime"
)

func (core *Core) Info(bot *tgbotapi.BotAPI, id int64) {
	pathContent := core.allDownloadClients.GetPathContent()
	textInfo := fmt.Sprintf("%s\nenv.PATH_TORRENT_CONTENT=%s ", osutils.InfoHost(), pathContent)

	size, err := file_func.FolderSize(&pathContent)
	if err != nil {
		slog.Error("telegram.GetInfo.FolderSize-", "error", err)
	}
	textListBots := fmt.Sprint(core.botsTG.GetListNameBot())
	textInfo += fmt.Sprintf("%d Mb Команды:\n/clear_content /show_content /info /start /down\n bots: %s\nOS: %s",
		size/1024/1024, textListBots, runtime.GOOS)

	_, _ = telegram.Send(bot, id, textInfo, nil)
}
