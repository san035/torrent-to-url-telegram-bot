package web_server

import (
	"github.com/rs/zerolog/log"
	"main.go/internal/torrent_client/torrent_anacrolix"
	"net/http"
	"os"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Info().Msg(path)
	filePath := torrent_anacrolix.PathTorrentContent + path

	if _, err2 := os.Stat(filePath); err2 == nil {
		http.ServeFile(w, r, filePath)
		return
	}
	http.NotFound(w, r)
}
