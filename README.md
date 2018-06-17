# About me
A Telegram Bot API wrapper

# Basic usage
```
// Create a bot instance
tgroid := telegroid.NewTelegroid("your_bot_token")

// Get bot information
me := tgroid.GetMe()

// Send message
args := telegroid.SendMessageArguments{}
args.ChatId = "@deadblue9"
args.Text = "Hello, world!"
tgroid.SendMessage(args)
```

# Supported Methods

| Method | Supported | Tested |
|--------|-----------|--------|
| getMe | Yes | Yes |
| getUpdates | Yes | Yes |
| setWebhook | Yes | No |
| deleteWebhook | Yes | No |
| getWebhookInfo | Yes | No |
| forwardMessage | Yes | No |
| sendMessage | Yes | Yes |
| sendPhoto | Yes | Yes |
| sendAudio | Yes | No |
| sendDocument | Yes | No |
| sendVideo | Yes | No |
| sendVoice | Yes | No |
| sendVideoNote | Yes | No |
| sendMediaGroup | Yes | No |
| sendLocation | Yes | No |
| sendVenue | Yes | No |
| sendContact | Yes | No |
| sendChatAction | Yes | No |


# TODO List
* Support all methods
* Add error handling
* Add utils for building arguments
* emmm...