//go:build linux

package torrent_client

import (
	"github.com/anacrolix/torrent"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

var (
	client             *torrent.Client
	PathTorrentContent string
)

func Init() (err error) {
	PathTorrentContent = os.Getenv("PATH_TORRENT_CONTENT")
	if PathTorrentContent == `` {
		PathTorrentContent = filepath.Dir(os.Args[0]) + `TORRENT_CONTENT` + string(os.PathSeparator)
	}

	err = os.MkdirAll(PathTorrentContent, 0755)
	if err != nil {
		return
	}

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = PathTorrentContent
	client, err = torrent.NewClient(clientConfig)
	if err != nil {
		return
	}
	log.Info().Str("path", PathTorrentContent).Msg("Start torrent client")
	return
}

func Close() {
	if client != nil {
		client.Close()
	}
}
