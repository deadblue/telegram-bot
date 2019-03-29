# About me

A Telegram Bot API wrapper

This project is an experiment of golang's reflection, although 
it works, but not better.

The new version is designing, and may be drop the reflection.

# Basic usage

```Go
import "github.com/deadblue/telegroid"
import "github.com/deadblue/telegroid/builder"

// Create a bot instance
tgroid := telegroid.New("your_bot_token")

// Get bot information
me, err := tgroid.GetMe()

// Build text message and send
args := builder.NewMessageBuilder().ChatId(12345678).
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
| answerCallbackQuery | ✓ | ✗ |
| answerInlineQuery | ✗ | ✗ |
| editMessageText | ✓ | ✗ |
| editMessageCaption | ✓ | ✗ |
| editMessageReplyMarkup | ✓ | ✗ |
| deleteMessage | ✓ | ✗ |
| editMessageLiveLocation | ✗ | ✗ |
| stopMessageLiveLocation | ✗ | ✗ |
| Payment methods | ✗ | ✗ |
| Game methods | ✗ | ✗ |

# TODO list
- [ ] Support more methods
- [ ] Add error handling
- [ ] Re-define the API to avoid reflection
