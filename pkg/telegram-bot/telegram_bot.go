package telegram_bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(api *tgbotapi.BotAPI) *TelegramBot {
	return &TelegramBot{
		bot: api,
	}

}
func (tb *TelegramBot) SetToken(token string) {
	tb.bot.Token = token
}
