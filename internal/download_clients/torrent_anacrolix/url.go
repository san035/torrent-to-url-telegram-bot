package torrent_anacrolix

import "strings"

const (
	NameClient = "torrent_anacrolix"
	UrlPattern = "magnet:"
)

func (clientAnacrolix *TorrentAnacrolix) GoodUrl(url *string) bool {
	return strings.HasPrefix(*url, UrlPattern)
}

func (clientAnacrolix *TorrentAnacrolix) GetUrlPattern() (nameClient, pattern string) {
	return NameClient, UrlPattern + "*"
}
