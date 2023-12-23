package download_clients

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
)

type AllClients struct {
	ListClient  []DownloadClient
	PathContent string
}

type DownloadClient interface {
	StartDownload(ctx *context.Context, urlMagnet *string) (chanStatus *chan StatusTorrent, err error)
	Close()
	GoodUrl(url *string) bool
}

const (
	StatusTorrentStart  = iota
	StatusTorrentRun    = iota
	StatusTorrentPause  = iota
	StatusTorrentEnd    = iota
	CountDownloadClient = 2
)

var DefaultAllClients *AllClients

func New() (downloadClient *AllClients, err error) {
	pathTorrentContent := os.Getenv("PATH_TORRENT_CONTENT")
	if pathTorrentContent == `` {
		pathTorrentContent = filepath.Dir(os.Args[0]) + `/TORRENT_CONTENT` + string(os.PathSeparator)
	}

	err = os.MkdirAll(pathTorrentContent, 0755)
	if err != nil {
		return
	}

	downloadClient = &AllClients{
		PathContent: pathTorrentContent,
		ListClient:  make([]DownloadClient, 0, CountDownloadClient),
	}

	return
}

func Init() (err error) {
	DefaultAllClients, err = New()
	return
}

func (allDownloadClient *AllClients) AddClient(downloadClient DownloadClient) {
	allDownloadClient.ListClient = append(allDownloadClient.ListClient, downloadClient)
}

func (allDownloadClient *AllClients) GetPathContent() string {
	return allDownloadClient.PathContent
}

func (allDownloadClient *AllClients) Close() {
	for _, client := range allDownloadClient.ListClient {
		client.Close()
	}
	return
}

type StatusTorrent struct {
	Info        string
	WebFileName *string
	Status      int
}

func (downloadClient *AllClients) RemoveAllContents() (counDelete int, err error) {

	// Открыть каталог
	dir, err := os.Open(downloadClient.PathContent)
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

		filePath := filepath.Join(downloadClient.PathContent, file.Name())
		err = os.Remove(filePath)
		if err != nil {
			return
		}
		counDelete++
	}

	return
}
