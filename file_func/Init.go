package file_func

import (
	"os"
)

func FileExists(filename *string) bool {
	_, err := os.Stat(*filename)
	return !os.IsNotExist(err)
}

//func SaveBodyToFile(body io.ReadCloser, filename *string) error {
//	defer body.Close()
//	file, err := os.Create(*filename)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	_, err = io.Copy(file, body)
//	return err
//}

//func IsTorrentFile(filename *string) bool {
//	return strings.HasSuffix(*filename, ".torrent")
//}
