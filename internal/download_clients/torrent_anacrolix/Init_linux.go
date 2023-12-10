//go:build linux

package torrent_anacrolix

import (
	"github.com/anacrolix/torrent"
	"github.com/rs/zerolog/log"
)

type TorrentAnacrolix struct {
	client *torrent.Client
}

func New(PathTorrentContent string) (*TorrentAnacrolix, error) {

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = PathTorrentContent
	client, err := torrent.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}
	torrentAnacrolix := &TorrentAnacrolix{client: client}
	log.Info().Str("path", PathTorrentContent).Msg("Start defaultClient torrent_anacrolix+")

	return torrentAnacrolix, nil
}

func (cl *TorrentAnacrolix) Close() {
	if cl.client != nil {
		cl.client.Close()
	}
}
