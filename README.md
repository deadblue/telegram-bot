# Telegroid

A Telegram Bot API wrapper for Go.

Supported Bot API version 4.9.

# Example

```Go
import "github.com/deadblue/telegroid"
import "github.com/deadblue/telegroid/parameters"

// Create a bot instance
bot := telegroid.New("your_bot_token")

// Get bot information
me, err := bot.GetMe()
if err != nil {
    panic(err)
}

// Send markdown text message with inline keyboard
params := new(parameters.SendMessageParams)
args.ChatId(1234)
args.Text(parameters.MarkdownV2Text("**Hello, world**"))
msg, err := bot.SendMessage(args)
if err != nil {
    panic(err)
}
```

# Reference

* https://core.telegram.org/bots/api

# LICENSE

MIT