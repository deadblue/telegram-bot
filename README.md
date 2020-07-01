# Telegram Bot

![](https://img.shields.io/badge/Release-v0.0.5-brightgreen.svg?style=flat-square)
![](https://img.shields.io/badge/Develop-v0.1.0-orange.svg?style=flat-square)
[![Reference](https://img.shields.io/badge/Go-Reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/deadblue/telegroid)
![](https://img.shields.io/github/license/deadblue/telegram-bot?style=flat-square)

Go bindings for Telegram Bot API.

Supports most Bot APIs in version 4.9, except:

* [Passport](https://core.telegram.org/bots/api#telegram-passport)

## Example

```go
import (
    "github.com/deadblue/telegram-bot"
    "github.com/deadblue/telegram-bot/parameters"
    "github.com/deadblue/telegram-bot/parameters/keyboard"
	"github.com/deadblue/telegram-bot/parameters/text"
)

// Create a bot instance
bot := telegram.New("your_bot_token")

// Get bot information
me, err := bot.GetMe()
if err != nil {
    panic(err)
}

// Prepare text message parameters
params := new(parameters.SendMessageParams)
params.ChatId(1234) 
params.Text(text.MarkdownV2("*Hello, world*"))
params.InlineKeyboard( (&keyboard.InlineBuilder{}).
    AddUrlButton("URL", "https://dead.blue").
    Build() )
// Send message
msg, err := bot.SendMessage(params)
if err != nil {
    panic(err)
}
```

## Reference

* https://core.telegram.org/bots/api

## LICENSE

MIT