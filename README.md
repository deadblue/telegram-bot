# About me

A Telegram Bot API wrapper

# Basic usage

```Go
import "github.com/deadblue/telegroid"

// Create a bot instance
tgroid := telegroid.NewTelegroid("your_bot_token")

// Get bot information
me, err := tgroid.GetMe()

// Build text message and send
args := telegroid.NewMessageBuilder().ChatId(12345678).
	ForText().Text("Hello, telegram!").Done()
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
| forwardMessage | ✓ | ✓ |
| sendMessage | ✓ | ✓ |
| sendPhoto | ✓ | ✓ |
| sendAudio | ✓ | ✓ |
| sendDocument | ✓ | ✓ |
| sendVideo | ✓ | ✓ |
| sendVoice | ✓ | ✗ |
| sendVideoNote | ✓ | ✗ |
| sendMediaGroup | ✓ | ✗ |
| sendLocation | ✓ | ✗ |
| sendVenue | ✓ | ✗ |
| sendContact | ✓ | ✗ |
| sendChatAction | ✓ | ✓ |
| getUserProfilePhotos | ✓ | ✓ |
| getFile | ✓ | ✓ |
| kickChatMember | ✓ | ✗ |
| unbanChatMember | ✓ | ✗ |
| restrictChatMember | ✓ | ✗ |
| promoteChatMember | ✓ | ✗ |
| exportChatInviteLink | ✓ | ✓ |
| setChatPhoto | ✓ | ✓ |
| deleteChatPhoto | ✓ | ✗ |
| setChatTitle | ✓ | ✓ |
| setChatDescription | ✓ | ✓ |
| pinChatMessage | ✓ | ✗ |
| unpinChatMessage | ✓ | ✗ |
| leaveChat | ✓ | ✓ |
| getChat | ✓ | ✓ |
| getChatAdministrators | ✓ | ✓ |
| getChatMembersCount | ✓ | ✓ |
| getChatMember | ✓ | ✗ |
| setChatStickerSet | ✓ | ✗ |
| deleteChatStickerSet | ✓ | ✗ |
| sendSticker | ✓ | ✓ |
| getStickerSet | ✓ | ✓ |
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
- [x] Add utils for building arguments
- [ ] Support more methods
- [ ] Add error handling
- [ ] emmmmm...