package torrent_client

import (
	"context"
	"os"
	"path/filepath"
)

type TorrentClient interface {
	StartTorrent(ctx *context.Context, urlMagnet *string) (chanStatus *chan StatusTorrent, err error)
	Close()
}

const (
	StatusTorrentStart = iota
	StatusTorrentRun   = iota
	StatusTorrentPause = iota
	StatusTorrentEnd   = iota
)

var (
	DefaultClient      TorrentClient
	PathTorrentContent string
)

func Init() (err error) {
	PathTorrentContent = os.Getenv("PATH_TORRENT_CONTENT")
	if PathTorrentContent == `` {
		PathTorrentContent = filepath.Dir(os.Args[0]) + `/TORRENT_CONTENT` + string(os.PathSeparator)
	}

	err = os.MkdirAll(PathTorrentContent, 0755)
	if err != nil {
		return
	}
	return
}

func GetPathTorrentContent() string {
	return PathTorrentContent
}

type StatusTorrent struct {
	Info        string
	WebFileName *string
	Status      int
}

func RemoveAllContents() error {
	files, err := os.ReadDir(PathTorrentContent)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(PathTorrentContent, file.Name())
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}

//func GetListDownloadContent() (listContent []string, err error) {
//	files, err := os.ReadDir(PathTorrentContent)
//	if err != nil {
//		return
//	}
//
//	for _, file := range files {
//		if file.IsDir() {
//			listContent = append(listContent, file.Name())
//		}
//	}
//
//	return
//
//}
