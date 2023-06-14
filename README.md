# torrent-to-url-telegram-bot
- [Gitlab](https://github.com/san035/torrent-to-url-telegram-bot)


## .env
```
BOT_TOKEN=you_token
BOT_ABOUT=bot answer url to torrent content
PORT= # default 8060
HOST= # default http://127.0.0.1
PATH_TORRENT_CONTENT= # default `TORRENT_CONTENT/`
TYPE_ANSWER= # file or url, default file
```

## build
```
git clone git@github.com:san035/torrent-to-url-telegram-bot.git
cd torrent-to-url-telegram-bot
nano .env
go build
```
