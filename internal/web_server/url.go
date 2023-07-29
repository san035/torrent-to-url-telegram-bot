package web_server

func GetUrl(fileName *string) string {
	return *HostAndPort + *fileName
}

func GetRooturl() string {
	return *HostAndPort
}
