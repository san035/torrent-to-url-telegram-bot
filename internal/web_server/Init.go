package web_server

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

const (
	HostDefault = "http://127.0.0.1"
	PortDefault = "8060"
)

type HttpService struct {
	port        string
	hostAndPort string
}

var WebServiceDefault *HttpService

func Init() error {
	WebServiceDefault = NewHttpService(os.Getenv("HOST"), os.Getenv("PORT"))
	go WebServiceDefault.Serve()
	return nil
}

func NewHttpService(host, port string) *HttpService {
	if port == "" {
		port = PortDefault
	}

	if host == "" {
		host = HostDefault
	}

	webService := &HttpService{
		port:        port,
		hostAndPort: host + `:` + port + `/`,
	}

	return webService
}

func (webService *HttpService) Serve() {
	http.HandleFunc(`/`, staticHandler)

	log.Info().Str("HostAndPort", webService.GetRooturl()).Msg("Starting server")
	err := http.ListenAndServe(`:`+webService.port, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("ListenAndServe")
	}

}
