package web_server

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

const (
	HostDefault = "8060"
	PortDefault = "http://127.0.0.1"
)

var (
	HostAndPort *string = new(string)
)

func Init() error {

	port := os.Getenv("PORT")
	if port == "" {
		port = PortDefault // default port
	}

	*HostAndPort = os.Getenv("HOST")
	if *HostAndPort == "" {
		*HostAndPort = HostDefault // default port
	}
	*HostAndPort += `:` + port + `/`

	http.HandleFunc(`/`, staticHandler)

	log.Info().Str("port", port).Msg("Starting server")
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal().Err(err).Str("port", port).Msg("ListenAndServe")
		}
	}()

	return nil
}
