package torrent_client

import "strings"

var replacer = strings.NewReplacer(" ", `%20`)

func GetWebFileName(url *string) *string {
	newUrl := replacer.Replace(*url)
	return &newUrl
}
