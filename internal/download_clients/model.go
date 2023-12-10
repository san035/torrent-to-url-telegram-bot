package download_clients

import (
	"context"
	"os"
	"path/filepath"
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

func (downloadClient *AllClients) RemoveAllContents() error {
	files, err := os.ReadDir(downloadClient.PathContent)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(downloadClient.PathContent, file.Name())
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}

//func GetListDownloadContent() (listContent []string, err error) {
//	files, err := os.ReadDir(PathContent)
//	if err != nil {
//		return
//	}
//
//	for _, file := range files {
//		if file.IsDir() {
//			listContent = append(listContent, file.Name())
//		}
//	}
//
//	return
//
//}
