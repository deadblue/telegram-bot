package telegroid

// The main api client

type Telegroid struct {
	// Get updates
	GetUpdates func(request *GetUpdatesRequest) ([]Update, error)

	// Webhook management
	SetWebhook     func(request *SetWebhookRequest) (bool, error)
	DeleteWebhook  func() (bool, error)
	GetWebhookInfo func() (*WebhookInfo, error)

	// Send message
	ForwardMessage func(args *ForwardMessageRequest) (*Message, error)
	SendMessage    func(request *SendMessageRequest) (*Message, error)
	SendPhoto      func(request *SendPhotoRequest) (*Message, error)
	SendAudio      func(args *SendAudioRequest) (*Message, error)
	SendDocument   func(args *SendDocumentRequest) (*Message, error)
	SendVideo      func(args *SendVideoRequest) (*Message, error)
	//SendAnimation  func(args *deprecated.SendAnimationArguments) (*Message, error)
	//SendVoice      func(args *deprecated.SendVoiceArguments) (*Message, error)
	//SendVideoNote  func(args *deprecated.SendVideoNoteArguments) (*Message, error)
	//SendMediaGroup func(args *deprecated.SendMediaGroupArguments) (*Message, error)
	//SendLocation   func(args *deprecated.SendLocationArguments) (*Message, error)
	//SendVenue      func(args *deprecated.SendVenueArguments) (*Message, error)
	//SendContact    func(args *deprecated.SendContentArguments) (*Message, error)
	//SendChatAction func(args *deprecated.SendChatActionArguments) (bool, error)

	// Edit/Delete message
	//EditMessageText              func(args *deprecated.EditMessageTextArguments) (*Message, error)
	//EditMessageCaption           func(args *deprecated.EditMessageCaptionArguments) (*Message, error)
	//editMessageMedia             func() (*Message, error) // TODO
	//EditMessageReplyMarkup       func(args *deprecated.EditMessageReplyMarkupArguments) (*Message, error)
	//EditMessageLiveLocation      func() (*Message, error) // TODO
	//StopMessageLiveLocation      func() (*Message, error) // TODO
	//EditOthersMessageText        func(args *deprecated.EditMessageTextArguments) (bool, error)        `method:"editMessageText"`
	//EditOthersMessageCaption     func(args *deprecated.EditMessageCaptionArguments) (bool, error)     `method:"editMessageCaption"`
	//EditOthersMessageReplyMarkup func(args *deprecated.EditMessageReplyMarkupArguments) (bool, error) `method:"editMessageReplyMarkup"`
	//DeleteMessage                func(args *deprecated.DeleteMessageArguments) (bool, error)

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
	GetChat               func(request *ChatRequest) (*Chat, error)
	ExportChatInviteLink  func(request *ChatRequest) (*string, error)
	GetChatAdministrators func(request *ChatRequest) ([]ChatMember, error)
	GetChatMembersCount   func(request *ChatRequest) (int, error)
	//GetChatMember         func(args *deprecated.ChatMemberArguments) (*ChatMember, error)

	// Misc
	GetFile func(request *GetFileRequest) (*File, error)

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
