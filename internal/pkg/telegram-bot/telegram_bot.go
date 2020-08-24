package telegram_bot

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot() *TelegramBot {
	return &TelegramBot{
		bot: &tgbotapi.BotAPI{},
	}
}

func (tb *TelegramBot) SetToken(token string) {
	tb.bot.Token = token
}

func (tb *TelegramBot) GetToken() string {
	return tb.bot.Token
}
