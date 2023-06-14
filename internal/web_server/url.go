package web_server

func GetUrl(fileName *string) string {
	return *hostAndPort + *fileName
}
