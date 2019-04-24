/*
A Telegram Bot API wrapper.

Current supports Bot API version is 4.2.

Example:

	import "github.com/deadblue/telegroid"
	import "github.com/deadblue/telegroid/arguments"

	// create bot instance
	bot := telegroid.New("your-bot-token")
	// Call getMe
	me, err := bot.GetMe()
	if err != nil {
		panic(err)
	}
	// Call sendMessage with a Markdown message
	args := new(arguments.SendMessageArgs)
	args.ChatId(1)
	args.Text("Hello, **world**").Markdown()
	msg, err := bot.SendMessage(args)
	if err != nil {
		panic(err)
	}
*/
package telegroid
