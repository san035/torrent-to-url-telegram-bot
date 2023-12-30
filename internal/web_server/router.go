package web_server

import (
	"log/slog"
	"net/http"
	"os"
)

func (webService *HttpService) staticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	slog.Info("web_server.staticHandler", "path", path)
	filePath := webService.pathContent + path

	if _, err2 := os.Stat(filePath); err2 == nil {
		http.ServeFile(w, r, filePath)
		return
	}
	http.NotFound(w, r)
}
