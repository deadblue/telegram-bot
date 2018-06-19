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

# Methods

| Method | Supported | Tested |
|--------|-----------|--------|
| getMe | ✓ | ✓ |
| getUpdates | ✓ | ✓ |
| setWebhook | ✓ | ✗ |
| deleteWebhook | ✓ | ✗ |
| getWebhookInfo | ✓ | ✗ |
| forwardMessage | ✓ | ✗ |
| sendMessage | ✓ | ✓ |
| sendPhoto | ✓ | ✓ |
| sendAudio | ✓ | ✗ |
| sendDocument | ✓ | ✗ |
| sendVideo | ✓ | ✗ |
| sendVoice | ✓ | ✗ |
| sendVideoNote | ✓ | ✗ |
| sendMediaGroup | ✓ | ✗ |
| sendLocation | ✓ | ✗ |
| sendVenue | ✓ | ✗ |
| sendContact | ✓ | ✗ |
| sendChatAction | ✓ | ✗ |
| getUserProfilePhotos | ✓ | ✗ |
| getFile | ✓ | ✗ |
| kickChatMember | ✓ | ✗ |
| unbanChatMember | ✓ | ✗ |
| restrictChatMember | ✓ | ✗ |
| promoteChatMember | ✓ | ✗ |
| exportChatInviteLink | ✓ | ✗ |
| setChatPhoto | ✓ | ✗ |
| deleteChatPhoto | ✓ | ✗ |
| setChatTitle | ✓ | ✗ |
| setChatDescription | ✓ | ✗ |
| pinChatMessage | ✓ | ✗ |
| unpinChatMessage | ✓ | ✗ |
| leaveChat | ✓ | ✗ |
| getChat | ✓ | ✗ |
| getChatAdministrators | ✓ | ✗ |
| getChatMembersCount | ✓ | ✗ |
| getChatMember | ✓ | ✗ |
| setChatStickerSet | ✓ | ✗ |
| deleteChatStickerSet | ✓ | ✗ |
| sendSticker | ✓ | ✗ |
| getStickerSet | ✓ | ✗ |
| uploadStickerFile | ✓ | ✗ |
| createNewStickerSet | ✓ | ✗ |
| addStickerToSet | ✓ | ✗ |
| setStickerPositionInSet | ✓ | ✗ |
| deleteStickerFromSet | ✓ | ✗ |
| answerCallbackQuery | ✗ | ✗ |
| answerInlineQuery | ✗ | ✗ |
| editMessageText | ✗ | ✗ |
| editMessageCaption | ✗ | ✗ |
| editMessageReplyMarkup | ✗ | ✗ |
| deleteMessage | ✓ | ✗ |
| editMessageLiveLocation | ✗ | ✗ |
| stopMessageLiveLocation | ✗ | ✗ |
| Payment methods | ✗ | ✗ |
| Game methods | ✗ | ✗ |

# TODO list
* Support more methods
* Add error handling
* Add utils for building arguments
* emmm...