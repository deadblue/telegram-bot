/*
Go bindings for Telegram Bot API.

Supports API version 4.9.

Example:

	import "github.com/deadblue/telegram-bot"
	import "github.com/deadblue/telegram-bot/parameters"

	// Create a bot instance
	bot := telegram.New("your_bot_token")

	// Get bot information
	me, err := bot.GetMe()
	if err != nil {
		panic(err)
	}

	// Send markdown text message with an inline keyboard
	params := new(parameters.SendMessageParams)
	params.ChatId(1234)
	params.Text(parameters.MarkdownV2Text("*Hello, world\\!*"))
	msg, err := bot.SendMessage(params)
	if err != nil {
		panic(err)
	}

*/
package telegram
