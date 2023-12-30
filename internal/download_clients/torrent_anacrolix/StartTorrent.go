//go:build linux

package torrent_anacrolix

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"main.go/internal/download_clients"
	"main.go/internal/file_func"
	"time"
)

const SecondWaitGotInfo = 30

func (clientAnacrolix *TorrentAnacrolix) StartDownload(ctx *context.Context, urlMagnet *string) (chanStatus *chan download_clients.StatusTorrent, err error) {
	ch := make(chan download_clients.StatusTorrent, 2)
	chanStatus = &ch

	go clientAnacrolix.Run(ctx, chanStatus, urlMagnet)
	return
}

// urlMagnet := "magnet:?xt=urn:btih:0EC1A755AEE37ACCD970D3C9662D93ABBCF26BE0&tr=http%3A%2F%2Fbt3.t-ru.org%2Fann%3Fmagnet&dn=Хороший%20доктор%20%2F%20The%20Good%20Doctor%20%2F%20Сезон%3A%201%20%2F%20Серии%3A%201-18%20из%2018%20(Майкл%20Патрик%20Джэнн%2C%20Нестор%20Карбонелл%2C%20Джон%20Дал)%20%5B2017%2C%20США%2C%20драма%2C%20WEB-DLRip%5D%20MVO%20(Lo"
func (clientAnacrolix *TorrentAnacrolix) Run(ctx *context.Context, chanStatus *chan download_clients.StatusTorrent, urlMagnet *string) {
	defer close(*chanStatus)
	statusTorrent := download_clients.StatusTorrent{Info: "Begin download", Status: download_clients.StatusTorrentStart}

	*chanStatus <- statusTorrent

	log.Info().Str("urlMagnet", *urlMagnet).Msg("Start torrent")
	t, err := clientAnacrolix.client.AddMagnet(*urlMagnet)
	if err != nil {
		*chanStatus <- download_clients.StatusTorrent{Info: "Error: " + err.Error()}
		log.Error().Err(err).Str("urlMagnet", *urlMagnet).Msg("error AddMagnet")
		return
	}

	// Проверяем, был ли контекст отменен
	select {
	case <-t.GotInfo():
	case <-time.After(SecondWaitGotInfo * time.Second):
		statusTorrent.Info = fmt.Sprintf("Error: Timeout GotInfo %d sec", SecondWaitGotInfo)
		statusTorrent.Status = 0
		*chanStatus <- statusTorrent
		return
	case <-(*ctx).Done():
		return
	}

	info := t.Info()

	statusTorrent.WebFileName = download_clients.GetWebFileName(&info.Name)
	fullFileName := clientAnacrolix.pathTorrentContent + info.Name
	if file_func.FileExists(&fullFileName) {
		statusTorrent.Info = ""
		statusTorrent.Status = download_clients.StatusTorrentEnd
		log.Info().Str(`file`, info.Name).Msg(statusTorrent.Info)
		*chanStatus <- statusTorrent
		return
	}

	statusTorrent.Info = fmt.Sprintf("Downloading ")
	statusTorrent.Status = download_clients.StatusTorrentRun
	*chanStatus <- statusTorrent

	log.Info().Str(`file`, info.Name).Msg(`Start download`)
	t.DownloadAll()

	rez := clientAnacrolix.client.WaitAll()
	log.Info().Str(`file`, info.Name).Bool("Status", rez).Msg(`End Download`)

	megabytes := fmt.Sprintf("%.1f Mb", float64(t.Length())/(1048576))
	if rez {
		statusTorrent.Info = "download completely " + megabytes + ` `
	} else {
		statusTorrent.Info = "!not fully downloaded! " + megabytes + ` `
	}
	statusTorrent.Status = download_clients.StatusTorrentEnd

	*chanStatus <- statusTorrent
	return
}
