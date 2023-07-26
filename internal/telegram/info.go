package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/file_func"
	"main.go/internal/torrent_client"
	"main.go/pkg/osutils"
)

func info(bot *tgbotapi.BotAPI, id int64) {
	folderData := torrent_client.GetPathTorrentContent()
	size, _ := file_func.FolderSize(folderData)
	text := fmt.Sprintf("%s\nenv.PATH_TORRENT_CONTENT=%s %d Mb /clear_content",
		osutils.InfoHost(), folderData, size/1024/1024)

	_, _ = Send(bot, id, text, nil)
}
