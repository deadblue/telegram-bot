package telegram

import (
	"github.com/deadblue/telegram-bot/parameters"
	"github.com/deadblue/telegram-bot/parameters/text"
	"log"
)

const (
	telegramBotToken = "your-bot-token"
)

func ExampleBot_GetMe() {
	// Create bot
	bot := New(telegramBotToken)
	// Call API
	user, err := bot.GetMe()
	if err != nil {
		log.Fatalf("Call getMe error: %s", err)
	} else {
		log.Printf("Bot name: %s", user.Username)
	}
}

func ExampleBot_SendMessage() {
	// Create bot
	bot := New(telegramBotToken)

	// Send text message
	params := new(parameters.SendMessageParams)
	params.ChatId(12345678)
	params.Text(text.MarkdownV2("Hello, world"))
	params.DisableNotification()
	msg, err := bot.SendMessage(params)
	if err != nil {
		log.Fatalf("Call sendMessage error: %s", err)
	} else {
		log.Printf("Sent message: %s", msg.Text)
	}
}
