package download_clients

import (
	"context"
	"os"
	"path/filepath"
)

type AllDownloadClient struct {
	ListClient  []DownloadClient
	PathContent string
}

type DownloadClient interface {
	StartDownload(ctx *context.Context, urlMagnet *string) (chanStatus *chan StatusTorrent, err error)
	Close()
	GoodUrl(url *string) bool
	GetUrlPattern() (nameClient, pattern string)
}

const (
	StatusTorrentStart  = iota
	StatusTorrentRun    = iota
	StatusTorrentPause  = iota
	StatusTorrentEnd    = iota
	CountDownloadClient = 2
)

func New() (downloadClient *AllDownloadClient, err error) {
	pathTorrentContent := os.Getenv("PATH_TORRENT_CONTENT")
	if pathTorrentContent == `` {
		pathTorrentContent = filepath.Dir(os.Args[0]) + `/TORRENT_CONTENT` + string(os.PathSeparator)
	}

	err = os.MkdirAll(pathTorrentContent, 0755)
	if err != nil {
		return
	}

	downloadClient = &AllDownloadClient{
		PathContent: pathTorrentContent,
		ListClient:  make([]DownloadClient, 0, CountDownloadClient),
	}

	return
}

func (allDownloadClient *AllDownloadClient) AddClient(downloadClient DownloadClient) {
	allDownloadClient.ListClient = append(allDownloadClient.ListClient, downloadClient)
}

func (allDownloadClient *AllDownloadClient) GetPathContent() string {
	return allDownloadClient.PathContent
}

func (allDownloadClient *AllDownloadClient) Close() {
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

// GetListSuppotUrlPattern получение словаяр поддерживаемых мосок url
func (allDownloadClient *AllDownloadClient) GetListSuppotUrlPattern() (mapPattern map[string]string) {
	mapPattern = map[string]string{}
	for _, client := range allDownloadClient.ListClient {
		nameClient, pattern := client.GetUrlPattern()
		mapPattern[nameClient] = pattern
	}
	return
}
