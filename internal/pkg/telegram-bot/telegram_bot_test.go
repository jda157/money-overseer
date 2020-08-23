package telegram_bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTelegramBot(t *testing.T) {
	tb := NewTelegramBot()
	assert.NotNil(t, tb)
	assert.NotNil(t, tb.bot)

}
