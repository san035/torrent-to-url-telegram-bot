// Модуль заглушка
package torrent_mock

import (
	"context"
	"main.go/internal/download_clients"
)

type TorrentMock struct {
}

//func Init() (err error) {
//	download_clients.ListClient = &TorrentMock{}
//	return
//}

func (torClient *TorrentMock) StartDownload(_ *context.Context, _ *string) (chanStatus *chan download_clients.StatusTorrent, err error) {
	//log.Info().Str("urlMagnet", *urlMagnet).Msg("torrent_mock.StartDownload+")
	return
}

func (torClient *TorrentMock) Close() {
	//log.Info().Msg("torrent_mock.Close+")
}

func (torClient *TorrentMock) GoodUrl(_ *string) bool {
	return true
}

func (torClient *TorrentMock) GetUrlPattern() (nameClient, pattern string) {
	return "mock", "*"
}
