package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func clearContent(bot *tgbotapi.BotAPI, id int64) {
	// todo
	//err := removeAllContents(torrent_client.GetPathTorrentContent())
	//if err != nil {
	//	_, _ = Send(bot, id, err, nil)
	//	return
	//}

	info(bot, id)
	return
}

func removeAllContents(folderPath string) error {
	err := os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
