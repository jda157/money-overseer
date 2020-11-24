package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var anotherKeyboard = tgbotapi.NewInlineKeyboardButtonSwitch("button", "cotton")

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ввести номер учетной записи.", "1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ввести ИНН.", "2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Загрузить логотип.", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ввести название организации.", "4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Прочитать оферту.", "5"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Загрузить файл с товарами.", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Страдать.", "7"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1466677055:AAHC7bUBbessYL0jLSpARJnyqj8bPkqUBaU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	fmt.Print(".")
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Print(update)
			log.Printf("\ngot some event: %d\n", update.CallbackQuery.Data)
			//bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))
			//
			//bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				msg.ReplyMarkup = numericKeyboard
			case "/top":
				msg.ReplyMarkup = tgbotapi.KeyboardButton{
					Text:            "some button",
					RequestLocation: false,
					RequestContact:  false,
				}
			}

			bot.Send(msg)
		}
	}
}
