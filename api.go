package telegroid

// The main api client

type Telegroid struct {
	// Get me
	GetMe func() (*User, error)
	// Get updates
	GetUpdates func(params *GetUpdatesParameters) ([]Update, error)
	// Webhook management
	SetWebhook     func(params *SetWebhookParameters) (bool, error)
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
	DeleteMessage func(params *ChatMessageParameters) (bool, error)
	// Edit own message
	EditMessageText         func(params *EditMessageTextParameters) (*Message, error)
	EditMessageCaption      func(params *EditMessageCaptionParameters) (*Message, error)
	EditMessageMedia        func(params *EditMessageMediaParameters) (*Message, error)
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
	// User information
	GetUserProfilePhotos func(params *GetUserProfilePhotosParameters) (*UserProfilePhotos, error)
	// Chat information
	GetChat               func(params *ChatParameters) (*Chat, error)
	GetChatAdministrators func(params *ChatParameters) ([]ChatMember, error)
	GetChatMembersCount   func(params *ChatParameters) (int, error)
	GetChatMember         func(params *ChatMemberParameters) (*ChatMember, error)
	// Chat management
	KickChatMember       func(params *KickChatMemberParameters) (bool, error)
	UnbanChatMember      func(params *ChatMemberParameters) (bool, error)
	RestrictChatMember   func(params *RestrictChatMemberParameters) (bool, error)
	PromoteChatMember    func(params *PromoteChatMemberParameters) (bool, error)
	ExportChatInviteLink func(params *ChatParameters) (string, error)
	SetChatPhoto         func(params *SetChatPhotoParameters) (bool, error)
	DeleteChatPhoto      func(params *ChatParameters) (bool, error)
	SetChatTitle         func(params *SetChatTitleParameters) (bool, error)
	SetChatDescription   func(params *SetChatDescriptionParameters) (bool, error)
	PinChatMessage       func(params *PinChatMessageParameters) (bool, error)
	UnpinChatMessage     func(params *ChatParameters) (bool, error)
	//SetChatStickerSet    func() (bool, error) // TODO
	//DeleteChatStickerSet func() (bool, error) // TODO
	LeaveChat func(params *ChatParameters) (bool, error)
	// Misc
	GetFile func(params *GetFileParameters) (*File, error)

	// Sticker
	SendSticker             func(params *SendStickerParameters) (*Message, error)
	GetStickerSet           func(params *GetStickerSetParameters) (*StickerSet, error)
	//UploadStickerFile       func(args *deprecated.UploadStickerFileArguments) (*File, error)
	//CreateNewStickerSet     func(args *deprecated.CreateNewStickerSetArguments) (bool, error)
	//AddStickerToSet         func(args *deprecated.AddStickerToSetArguments) (bool, error)
	//SetStickerPositionInSet func(args *deprecated.SetStickerPositionInSetArguments) (bool, error)
	//DeleteStickerFromSet    func(args *deprecated.DeleteStickerFromSetArguments) (bool, error)

	AnswerCallbackQuery func(params *AnswerCallbackQueryParameters) (bool, error)
	//AnswerInlineQuery   func() (bool, error) // TODO
}

func New(token string) *Telegroid {
	tg := new(Telegroid)
	bindInvoker(tg, token)
	return tg
}
