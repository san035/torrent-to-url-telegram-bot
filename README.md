# torrent-to-url-telegram-bot
- [Gitlab](https://github.com/san035/torrent-to-url-telegram-bot)


## .env
```
# required parameters
BOT_TOKEN=you_token
PORT=8060

# optional parameters

#default "bot answer url to torrent content"
BOT_ABOUT=

# default http://127.0.0.1
HOST= 

# default `TORRENT_CONTENT/`
PATH_TORRENT_CONTENT= 

# list id telegrem by "," 
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
