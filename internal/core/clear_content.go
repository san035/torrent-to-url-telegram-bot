package core

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/telegram"
	"os"
	"path"
	"path/filepath"
	"slices"
)

func (core *Core) ClearContent(bot *tgbotapi.BotAPI, id int64) {
	counDelete, err := core.clearFolderContent()
	if err != nil {
		_, _ = telegram.Send(bot, id, err, nil)
		return
	}
	_, _ = telegram.Send(bot, id, fmt.Sprintf("Удалено %d файлов/каталогов", counDelete), nil)
	return
}

func (core *Core) clearFolderContent() (counDelete int, err error) {

	pathContent := core.allDownloadClients.GetPathContent()

	// Открыть каталог
	dir, err := os.Open(pathContent)
	if err != nil {
		err = fmt.Errorf("Ошибка открытия каталога: %s", err)
		return
	}
	defer dir.Close()

	// Получить информацию о содержимом каталога
	fileInfo, err := dir.Readdir(0) // 0 означает получить список всех файлов и папок
	if err != nil {
		err = fmt.Errorf("Ошибка чтения содержимого каталога: %s", err)
		return
	}

	skipFiles := []string{".torrent.db", ".torrent.db-wal", ".torrent.db-shm"}
	for _, file := range fileInfo {
		if slices.Contains(skipFiles, path.Base(file.Name())) {
			continue
		}

		filePath := filepath.Join(pathContent, file.Name())
		err = os.Remove(filePath)
		if err != nil {
			return
		}
		counDelete++
	}

	return
}
