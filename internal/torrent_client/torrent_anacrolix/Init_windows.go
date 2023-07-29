//go:build windows

package torrent_anacrolix

import "errors"

var PathTorrentContent string

func Init() (err error) {
	return errors.New("os windows is not supported github.com/anacrolix/torrent")
}

func Close() {
	return
}
