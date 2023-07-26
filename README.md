# torrent-to-url-telegram-bot
- [Gitlab](https://github.com/san035/torrent-to-url-telegram-bot)


## .env
```
# required parameters

# LIST_BOT_TOKEN split ','
LIST_BOT_TOKEN=
PORT=8060

# optional parameters

#default "bot answer url to torrent content"
BOT_ABOUT=

# default http://127.0.0.1
HOST= 

# default `TORRENT_CONTENT/`
PATH_TORRENT_CONTENT= 

# list id telegrem separat by "," 
# all admins must be first in list LIST_BOT_TOKEN
LIST_ADMIN_ID_TELEGRAM=

TYPE_ANSWER= # file or url, default file
```

## build
```
git clone git@github.com:san035/torrent-to-url-telegram-bot.git
cd torrent-to-url-telegram-bot
# edit .env
up.sh
```
