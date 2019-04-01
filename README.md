# Telegroid

A Telegram Bot API wrapper in golang

Version: 1.0 (WIP)

# Example

```Go
import "github.com/deadblue/telegroid"

// Create a bot instance
bot := telegroid.New("your_bot_token")

// Get bot information
me, err := bot.GetMe()

// Send text message
args := (&SendRequestBuilder{}).
	WithChatId(1234).
	DisableNotification().
	    ForText().WithText("Hello, world!")
msg, err := bot.SendMessage(args)
```
