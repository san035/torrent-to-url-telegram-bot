package telegram

func (botsTelegram *BotsTelegram) AddCommand(addMapCmd MapCmd) {
	for cmd, funcCmd := range addMapCmd {
		botsTelegram.MapCmd[cmd] = funcCmd
	}
}
