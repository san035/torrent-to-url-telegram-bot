package torrent_client

import (
	"errors"
	"strings"
)

func CheckUrl(url *string) error {
	if !strings.HasPrefix(*url, "magnet:") {
		return errors.New(`the URL must start with "magnet:"`)
	}
	return nil
}
