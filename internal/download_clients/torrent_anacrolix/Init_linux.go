//go:build linux

package torrent_anacrolix

import (
	"github.com/anacrolix/torrent"
	"github.com/rs/zerolog/log"
)

type TorrentAnacrolix struct {
	client             *torrent.Client
	pathTorrentContent string
}

func New(pathTorrentContent string) (*TorrentAnacrolix, error) {

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = pathTorrentContent
	client, err := torrent.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}
	torrentAnacrolix := &TorrentAnacrolix{
		client:             client,
		pathTorrentContent: pathTorrentContent,
	}
	log.Info().Str("path", pathTorrentContent).Msg("Start defaultClient torrent_anacrolix+")

	return torrentAnacrolix, nil
}

func (clientAnacrolix *TorrentAnacrolix) Close() {
	if clientAnacrolix.client != nil {
		clientAnacrolix.client.Close()
	}
}
