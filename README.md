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
params := new(telegroid.SendMessageParameters)
params.ChatId(1234)
params.Text("Hello, world!")
msg, err := bot.SendMessage(params)



```
