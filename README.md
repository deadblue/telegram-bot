# Telegroid

A _Telegram Bot API_ wrapper for Golang.

Current version is `0.9.0`, which supports Bot API 4.2 without some rarely-used (~~in my opinion~~) functions.

# Example

```Go
import "github.com/deadblue/telegroid"
import "github.com/deadblue/telegroid/arguments"

// Create a bot instance
bot := telegroid.New("your_bot_token")

// Get bot information
me, err := bot.GetMe()

// Send text message
args := new(arguments.SendMessageArgs)
args.ChatId(1234)
args.Text("Hello, world!")
msg, err := bot.SendMessage(args)

```

# Feature:

Telegroid currently supports most of functions, **except**:

* Inline mode
* Telegram Passport
* Payment

These functions will be supported in the plan.

# Reference

* https://core.telegram.org/bots/api
