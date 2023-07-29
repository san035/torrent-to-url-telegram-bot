// Модуль заглушка
package torrent_mock

import (
	"context"
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client"
)

type TorrentMock struct {
}

func Init() (err error) {
	torrent_client.DefaultClient = &TorrentMock{}
	return
}

func (torClient *TorrentMock) StartTorrent(ctx *context.Context, urlMagnet *string) (chanStatus *chan torrent_client.StatusTorrent, err error) {
	log.Info().Str("urlMagnet", *urlMagnet).Msg("torrent_mock.StartTorrent+")
	return
}

func (torClient *TorrentMock) Close() {
	log.Info().Msg("torrent_mock.Close+")
}
