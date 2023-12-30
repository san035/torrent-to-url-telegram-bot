//go:build linux

package torrent_anacrolix

import (
	"github.com/anacrolix/torrent"
	"log/slog"
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
	slog.Info("Start defaultClient torrent_anacrolix+", "pathTorrentContent", pathTorrentContent)

	return torrentAnacrolix, nil
}

func (clientAnacrolix *TorrentAnacrolix) Close() {
	if clientAnacrolix.client != nil {
		clientAnacrolix.client.Close()
	}
}
