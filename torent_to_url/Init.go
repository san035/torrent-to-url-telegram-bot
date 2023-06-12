package torent_to_url

import "time"

func GetChanMessage(urlMagnet *string) (chanStatus *chan string, err error) {
	ch := make(chan string, 2)
	chanStatus = &ch

	go Run(chanStatus)
	return
}

func Run(chanStatus *chan string) {
	*chanStatus <- "Начало закачки "
	time.Sleep(time.Second * 3)
	*chanStatus <- "Загрузка "
	time.Sleep(time.Second * 3)
	*chanStatus <- "Загружено"
	return
}
