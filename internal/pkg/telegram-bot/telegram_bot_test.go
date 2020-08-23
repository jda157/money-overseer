package telegram_bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTelegramBot(t *testing.T) {
	t.Parallel()

	tb := NewTelegramBot()
	assert.NotNil(t, tb, "TelegramBot wasn't initiated.")
	assert.NotNil(t, tb.bot, "bot field in TelegramBot wasn't initiated")

}

func TestTelegramBot_SetToken(t *testing.T) {
	t.Parallel()

	token := "someTelegramToken"

	tb := NewTelegramBot()
	tb.SetToken(token)

	assert.Equal(t, token, tb.bot.Token)
}

func TestTelegramBot_GetToken(t *testing.T) {
	t.Parallel()

	token := "SomeTelegramToken"

	tb := NewTelegramBot()
	tb.bot.Token = token

	assert.Equal(t, token, tb.GetToken())
}
