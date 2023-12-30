//go:build windows

package tmp

import "errors"

var PathTorrentContent string

func Init() (err error) {
	return errors.New("os windows is not supported github.com/anacrolix/torrent")
}

func Close() {
	return
}
