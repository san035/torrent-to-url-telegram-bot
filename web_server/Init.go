package web_server

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func Init() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8060" // default port
	}

	http.HandleFunc("/", staticHandler)

	log.Info().Str("port", port).Msg("Starting server")
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal().Err(err).Str("port", port).Msg("ListenAndServe")
		}
	}()

	return nil
}
