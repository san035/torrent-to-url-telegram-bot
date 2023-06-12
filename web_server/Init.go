package web_server

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

var listPublicFolder = []string{"public", "images", "css"} // list of public folders

func Init() (err error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		for _, folder := range listPublicFolder {
			filePath := fmt.Sprintf("%s/%s", folder, path)
			if _, err := os.Stat(filePath); err == nil {
				http.ServeFile(w, r, filePath)
				return
			}
		}
		http.NotFound(w, r)
	})

	log.Info().Str("port", port).Msg("Starting server")
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal().Err(err).Str("port", port).Msg("ListenAndServe")
		}
	}()
	return
}
