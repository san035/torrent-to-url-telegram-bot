package web_server

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

const (
	ENV_NAME_WEB_HOST = "WEB_HOST"
	ENV_NAME_WEB_PORT = "WEB_PORT"
	HostDefault       = "http://127.0.0.1"
	PortDefault       = "8060"
)

type HttpService struct {
	port        string
	hostAndPort string
	pathContent string
}

func New(pathContent string) *HttpService {
	port, ok := os.LookupEnv(ENV_NAME_WEB_PORT)
	if !ok || port == "" {
		port = PortDefault
	}

	host, ok := os.LookupEnv(ENV_NAME_WEB_HOST)
	if !ok || host == "" {
		host = HostDefault
	}

	webService := &HttpService{
		port:        port,
		hostAndPort: host + `:` + port + `/`,
		pathContent: pathContent,
	}

	go webService.Serve()

	return webService
}

func (webService *HttpService) Serve() {
	http.HandleFunc(`/`, webService.staticHandler)

	log.Info().Str("HostAndPort", webService.GetRooturl()).Msg("Starting server")
	err := http.ListenAndServe(`:`+webService.port, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("ListenAndServe")
	}

}
