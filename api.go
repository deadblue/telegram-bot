package telegroid

// The main api client

type Telegroid struct {
	// Get updates
	GetUpdates func(params *GetUpdatesRequest) ([]Update, error)
	// Webhook management
	SetWebhook     func(params *SetWebhookRequest) (bool, error)
	DeleteWebhook  func() (bool, error)
	GetWebhookInfo func() (*WebhookInfo, error)
	// Send message
	ForwardMessage func(params *ForwardMessageParameters) (*Message, error)
	SendMessage    func(params *SendMessageParameters) (*Message, error)
	SendPhoto      func(params *SendPhotoParameters) (*Message, error)
	SendAudio      func(params *SendAudioParameters) (*Message, error)
	SendDocument   func(params *SendDocumentParameters) (*Message, error)
	SendVideo      func(params *SendVideoParameters) (*Message, error)
	SendAnimation  func(params *SendAnimationParameters) (*Message, error)
	SendVoice      func(params *SendVoiceParameters) (*Message, error)
	SendVideoNote  func(params *SendVideoNoteParameters) (*Message, error)
	//SendMediaGroup func(args *deprecated.SendMediaGroupArguments) (*Message, error)
	SendLocation   func(params *SendLocationParameters) (*Message, error)
	SendVenue      func(params *SendVenueParameters) (*Message, error)
	SendContact    func(params *SendContactParameters) (*Message, error)
	SendChatAction func(params *SendChatActionParameters) (bool, error)

	// Delete message
	DeleteMessage func(params *DeleteMessageParameters) (bool, error)
	// Edit own message
	EditMessageText         func(params *EditMessageTextParameters) (*Message, error)
	EditMessageCaption      func(params *EditMessageCaptionParameters) (*Message, error)
	editMessageMedia        func(params *EditMessageMediaParameters) (*Message, error)
	EditMessageReplyMarkup  func(params *EditMessageReplyMarkupParameters) (*Message, error)
	EditMessageLiveLocation func(params *EditMessageLiveLocationParameters) (*Message, error)
	StopMessageLiveLocation func(params *EditMessageReplyMarkupParameters) (*Message, error)
	// Edit other's message
	EditOthersMessageText         func(params *EditMessageTextParameters) (bool, error)         `method:"editMessageText"`
	EditOthersMessageCaption      func(params *EditMessageCaptionParameters) (bool, error)      `method:"editMessageCaption"`
	EditOthersMessageMedia        func(params *EditMessageMediaParameters) (bool, error)        `method:"editMessageMedia"`
	EditOthersMessageReplyMarkup  func(params *EditMessageReplyMarkupParameters) (bool, error)  `method:"editMessageReplyMarkup"`
	EditOthersMessageLiveLocation func(params *EditMessageLiveLocationParameters) (bool, error) `method:"editMessageLiveLocation"`
	StopOthersMessageLiveLocation func(params *EditMessageReplyMarkupParameters) (bool, error)  `method:"stopMessageLiveLocation"`

	// Chat management
	//KickChatMember       func(args *deprecated.KickChatMemberArguments) (bool, error)
	//UnbanChatMember      func(args *deprecated.ChatMemberArguments) (bool, error)
	//RestrictChatMember   func(args *deprecated.RestrictChatMemberArguments) (bool, error)
	//PromoteChatMember    func(args *deprecated.PromoteChatMemberArguments) (bool, error)
	//SetChatPhoto         func(args *deprecated.SetChatPhotoArguments) (bool, error)
	//DeleteChatPhoto      func(args *deprecated.ChatArguments) (bool, error)
	//SetChatTitle         func(args *deprecated.SetChatTitleArguments) (bool, error)
	//SetChatDescription   func(args *deprecated.SetChatDescriptionArguments) (bool, error)
	//PinChatMessage       func(args *deprecated.PinChatMessageArguments) (bool, error)
	//UnpinChatMessage     func(args *deprecated.ChatArguments) (bool, error)
	//SetChatStickerSet    func() (bool, error) // TODO
	//DeleteChatStickerSet func() (bool, error) // TODO
	//LeaveChat            func(args *deprecated.ChatArguments) (bool, error)

	// User information
	GetMe func() (*User, error)
	//GetUserProfilePhotos func(args *deprecated.GetUserProfilePhotosArguments) (*UserProfilePhotos, error)

	// Chat information
	GetChat               func(params *ChatParameters) (*Chat, error)
	ExportChatInviteLink  func(params *ChatParameters) (*string, error)
	GetChatAdministrators func(params *ChatParameters) ([]ChatMember, error)
	GetChatMembersCount   func(params *ChatParameters) (int, error)
	//GetChatMember         func(args *deprecated.ChatMemberArguments) (*ChatMember, error)

	// Misc
	GetFile func(params *GetFileRequest) (*File, error)

	// Sticker
	//SendSticker             func(args *deprecated.SendStickerArguments) (*Message, error)
	//GetStickerSet           func(args *deprecated.GetStickerSetArguments) (*StickerSet, error)
	//UploadStickerFile       func(args *deprecated.UploadStickerFileArguments) (*File, error)
	//CreateNewStickerSet     func(args *deprecated.CreateNewStickerSetArguments) (bool, error)
	//AddStickerToSet         func(args *deprecated.AddStickerToSetArguments) (bool, error)
	//SetStickerPositionInSet func(args *deprecated.SetStickerPositionInSetArguments) (bool, error)
	//DeleteStickerFromSet    func(args *deprecated.DeleteStickerFromSetArguments) (bool, error)

	//AnswerCallbackQuery func(args *deprecated.AnwserCallbackQueryArguments) (bool, error)
	//AnswerInlineQuery   func() (bool, error) // TODO
}

func New(token string) *Telegroid {
	tg := new(Telegroid)
	bindInvoker(tg, token)
	return tg
}
