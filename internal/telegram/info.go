package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/internal/download_clients"
	"main.go/internal/file_func"
	"main.go/pkg/osutils"
	"runtime"
)

func info(bot *tgbotapi.BotAPI, id int64) {
	_, _ = Send(bot, id, GetInfo(), nil)
}

func GetInfo() string {
	folderData := download_clients.DefaultAllClients.GetPathContent()
	textInfo := fmt.Sprintf("%s\nenv.PATH_TORRENT_CONTENT=%s ", osutils.InfoHost(), folderData)

	size, err := file_func.FolderSize(folderData)
	if err != nil {
		log.Error().Err(err).Msg("telegram.GetInfo.FolderSize-")
	}
	textInfo += fmt.Sprintf("%d Mb /clear_content /show_content\n bots: %s\nOS: %s", size/1024/1024, fmt.Sprint(GetListNameBot()), runtime.GOOS)

	return textInfo
}
