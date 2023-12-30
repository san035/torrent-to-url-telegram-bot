package youtuber_client

import (
	"context"
	"github.com/kkdai/youtube/v2"
	"main.go/internal/download_clients"
	"strings"
)

const (
	NameClient = "youtube"
	UrlPattern = "https://www.youtube.com/watch?v="
)

type youtubeClient struct {
	client youtube.Client
}

func (yClient *youtubeClient) GoodUrl(url *string) bool {
	return strings.HasPrefix(*url, UrlPattern)
}

func (yClient *youtubeClient) StartDownload(ctx *context.Context, urlMagnet *string) (chanStatus *chan download_clients.StatusTorrent, err error) {
	return
}

func (yClient *youtubeClient) Close() {

}

func (yClient *youtubeClient) GetUrlPattern() (nameClient, pattern string) {
	return NameClient, UrlPattern + "*"
}
