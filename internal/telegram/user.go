package telegram

import (
	"github.com/anacrolix/sync"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegrammUser struct {
	NickName string
}
type MapNikName struct {
	sync.RWMutex
	MapNikName map[int64]*TelegrammUser
}

var mapNikName = MapNikName{MapNikName: map[int64]*TelegrammUser{}}

func SaveUser(u *tgbotapi.Chat) {
	mapNikName.Lock()
	defer mapNikName.Unlock()

	mapNikName.MapNikName[u.ID] = &TelegrammUser{NickName: u.UserName}
}

// nikname пользователя по id
func NikNameById(id int64) string {
	mapNikName.RLock()
	defer mapNikName.RUnlock()

	u, ok := mapNikName.MapNikName[id]
	if !ok {
		return ""
	}

	return u.NickName
}
