package web_server

func (webService *HttpService) GetUrl(fileName *string) string {
	return webService.hostAndPort + *fileName
}

func (webService *HttpService) GetRooturl() string {
	return webService.hostAndPort
}
