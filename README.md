# Telegroid

A _Telegram Bot API_ wrapper for Golang.

Currently supports Bot API 4.2 without some rarely-used (~~in my opinion~~) functions.

# Example

```Go
import "github.com/deadblue/telegroid"
import "github.com/deadblue/telegroid/arguments"

// Create a bot instance
bot := telegroid.New("your_bot_token")

// Get bot information
me, err := bot.GetMe()
if err != nil {
    panic(err)
}

// Send markdown text message with inline keyboard
args := new(arguments.SendMessageArgs)
args.ChatId(1234)
args.Text("Hello, *world*!").Markdown()
args.InlineKeyboard().
    UrlButton("Github", "https://github.com/deadblue/telegroid").
    UrlButton("Author", "tg://resolve?domain=deadbluex").
    CallbackButton("Foo", "Bar").
    Layout(2, 1).
    Finish()
msg, err := bot.SendMessage(args)
if err != nil {
    panic(err)
}

```

# Feature:

Telegroid currently supports most of functions, **except**:

* Inline mode
* Telegram Passport
* Payment

These functions will be supported in the plan.

# Reference

* https://core.telegram.org/bots/api
