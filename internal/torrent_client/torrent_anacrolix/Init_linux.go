//go:build linux

package torrent_anacrolix

import (
	"github.com/anacrolix/torrent"
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client"
)

type TorrentAnacrolix struct {
	client *torrent.Client
}

var defaultClient *torrent.Client

func Init() (err error) {

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = torrent_client.PathTorrentContent
	defaultClient, err = torrent.NewClient(clientConfig)
	if err != nil {
		return
	}
	log.Info().Str("path", torrent_client.PathTorrentContent).Msg("Start defaultClient torrent_anacrolix+")

	torrent_client.DefaultClient = &TorrentAnacrolix{client: defaultClient}

	return
}

func (сlientAnacrolix *TorrentAnacrolix) Close() {
	if сlientAnacrolix.client != nil {
		сlientAnacrolix.client.Close()
	}
}
