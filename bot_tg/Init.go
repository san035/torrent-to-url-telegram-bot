package bot_tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"main.go/torent_to_url"
	"os"
)

var (
	Bot     *tgbotapi.BotAPI
	updates *tgbotapi.UpdatesChannel
)

func Init() (err error) {
	Bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates1 := Bot.GetUpdatesChan(u)
	updates = &updates1
	return nil
}

func Run() {
	for update := range *updates {

		if update.Message == nil { // ignore non-message updates
			continue
		}

		cmd := update.Message.Command()
		log.Info().Int("id", update.UpdateID).Msg("cmd:" + cmd)
		switch cmd {
		case "start":
			Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, os.Getenv("BOT_ABOUT")))
			continue
		default:

			url, err := torent_to_url.ServUrlByTorrent(&update.Message.Text)
			if err != nil {
				Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				continue
			}

			Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, url))

		}

		//if update.Message.Document == nil {
		//	Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Send me a torrent file and I will download it for you."))
		//	continue
		//}
		//if !file_func.IsTorrentFile(&update.Message.Document.FileName) {
		//	Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "File is not Torrent"))
		//	continue
		//
		//}
		//
		//FileNameTorent := pathTorrentContent + update.Message.Document.FileName
		//
		//if !file_func.fileExists(FileNameTorent) {
		//
		//	fileURL, err := Bot.GetFileDirectURL(update.Message.Document.FileID)
		//	if err != nil {
		//		log.Printf("Error getting file URL: %v", err)
		//		log.Error().Err(err).Str("FileID", update.Message.Document.FileID).Msg("Error getting file URL")
		//		continue
		//	}
		//
		//	resp, err := http.Get(fileURL)
		//	if err != nil {
		//		log.Error().Err(err).Msg("Error downloading file")
		//		continue
		//	}
		//
		//	err = file_func.saveBodyToFile(resp.Body, &FileNameTorent)
		//	if err != nil {
		//		log.Error().Err(err).Msg("Error by FileNameTorent")
		//		Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Error by saveBodyToFile"))
		//	}
		//
		//	log.Info().Str("file", FileNameTorent).Msg("Save torrent")
		//}

		//	t, err := client.AddTorrentFromFile(tmpfile.Name())
		//	if err != nil {
		//		log.Printf("Error adding torrent: %v", err)
		//		continue
		//	}
		//
		//	<-t.GotInfo()
		//
		//	fmt.Printf("Downloading %d files...\n", len(t.Files()))
		//
		//	for _, file := range t.Files() {
		//		file.Download()
		//		//if err != nil {
		//		//	log.Printf("Error downloading %s: %v", file.Path(), err)
		//		//} else {
		//		//}
		//		fmt.Printf("Downloaded %s\n", file.Path())
		//	}
		//
		//	//publicLink, err := bot.UploadFile(t.Files()[0].Path())
		//	//if err != nil {
		//	//	log.Printf("Error uploading file: %v", err)
		//	//	continue
		//	//}
		//
		//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "publicLink")
		//	bot.Send(msg)
	}

}
