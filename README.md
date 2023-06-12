# torrent-to-url-telegram-bot
- [Gitlab](https://github.com/san035/torrent-to-url-telegram-bot)


## .env
```
BOT_TOKEN=you_token
BOT_ABOUT=bot answer url to torrent content
PORT=8080

PATH_TORRENT_CONTENT=/temp/
```

## build
```
git clone git@github.com:san035/torrent-to-url-telegram-bot.git
cd torrent-to-url-telegram-bot
nano .env
go build
```