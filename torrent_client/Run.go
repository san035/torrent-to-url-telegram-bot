package torrent_client

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"main.go/file_func"
	"time"
)

const SecondWaitGotInfo = 30

type StatusTorrent struct {
	Info     string
	FileName *string
}

func GetChanMessage(ctx *context.Context, urlMagnet *string) (chanStatus *chan StatusTorrent, err error) {
	ch := make(chan StatusTorrent, 2)
	chanStatus = &ch

	go Run(ctx, chanStatus, urlMagnet)
	return
}

// urlMagnet := "magnet:?xt=urn:btih:0EC1A755AEE37ACCD970D3C9662D93ABBCF26BE0&tr=http%3A%2F%2Fbt3.t-ru.org%2Fann%3Fmagnet&dn=Хороший%20доктор%20%2F%20The%20Good%20Doctor%20%2F%20Сезон%3A%201%20%2F%20Серии%3A%201-18%20из%2018%20(Майкл%20Патрик%20Джэнн%2C%20Нестор%20Карбонелл%2C%20Джон%20Дал)%20%5B2017%2C%20США%2C%20драма%2C%20WEB-DLRip%5D%20MVO%20(Lo"
func Run(ctx *context.Context, chanStatus *chan StatusTorrent, urlMagnet *string) {
	defer close(*chanStatus)
	statusTorrent := StatusTorrent{Info: "Начало закачки"}

	*chanStatus <- statusTorrent

	log.Info().Str("urlMagnet", *urlMagnet).Msg("Start torrent")
	t, err := client.AddMagnet(*urlMagnet)
	if err != nil {
		*chanStatus <- StatusTorrent{Info: "Error: " + err.Error()}
		log.Error().Err(err).Str("urlMagnet", *urlMagnet).Msg("error AddMagnet")
		return
	}

	// Проверяем, был ли контекст отменен
	select {
	case <-t.GotInfo():
	case <-time.After(SecondWaitGotInfo * time.Second):
		statusTorrent.Info = fmt.Sprintf("Error: Timeout GotInfo %d sec", SecondWaitGotInfo)
		*chanStatus <- statusTorrent
		return
	case <-(*ctx).Done():
		return
	}

	info := t.Info()

	statusTorrent.FileName = &info.Name
	fileName := PathTorrentContent + info.Name
	if file_func.FileExists(&fileName) {
		statusTorrent.Info = "Already downloaded"
		log.Info().Str(`file`, info.Name).Msg(statusTorrent.Info)
		*chanStatus <- statusTorrent
		return
	}

	statusTorrent.Info = fmt.Sprintf("Downloading %d files", len(info.Files))
	*chanStatus <- statusTorrent

	log.Info().Str(`file`, info.Name).Msg(`Start DownloadAll`)
	t.DownloadAll()
	client.WaitAll()
	log.Info().Str(`file`, info.Name).Msg(`End Download`)

	statusTorrent.Info = "End download"
	*chanStatus <- statusTorrent
	return
}
