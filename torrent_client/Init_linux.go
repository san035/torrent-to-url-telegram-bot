//go:build linux

package torrent_client

import (
	"github.com/anacrolix/torrent"
	"os"
)

var client *torrent.Client
var PathTorrentContent string

func Init() (err error) {
	PathTorrentContent = os.Getenv("PATH_TORRENT_CONTENT")
	err = os.MkdirAll(PathTorrentContent, 0755)
	if err != nil {
		return
	}

	clientConfig := torrent.NewDefaultClientConfig()
	client, err = torrent.NewClient(clientConfig)
	return
}

func Close() {
	if client != nil {
		client.Close()
	}
}
