/*
A Telegram Bot API wrapper.

Current supports Bot API version is 4.2.

Example:

	import "github.com/deadblue/telegroid"
	import "github.com/deadblue/telegroid/arguments"

	// Create a bot instance
	bot := telegroid.New("your_bot_token")

	// Get bot information
	me, err := bot.GetMe()
	if err != nil {
		panic(err)
	}

	// Send markdown text message with an inline keyboard
	args := new(arguments.SendMessageArgs)
	args.ChatId(1234)
	args.Text("Hello, *world*!").Markdown()
	args.InlineKeyboard().
		AddUrlButton("Github", "https://github.com/deadblue/telegroid").
		AddUrlButton("Author", "tg://resolve?domain=deadbluex").
		AddCallbackButton("Foo", "Bar").
		Layout(2, 1).
		Finish()
	msg, err := bot.SendMessage(args)
	if err != nil {
		panic(err)
	}

*/
package telegroid
