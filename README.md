# Telegroid

A Telegram Bot API wrapper in golang

Version: `1.0` (WIP)

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

# Features

The goal is to support Telegram Bot API `4.2`.

But following APIs are not supported yet:

* Payments API
* Passport API

# Reference

* https://core.telegram.org/bots/api
