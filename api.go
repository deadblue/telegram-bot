package telegroid

// The main api client

type Telegroid struct {
	GetMe                        func() (*User, error)
	GetUpdates                   func(args *GetUpdatesArguments) ([]Update, error)
	SetWebhook                   func(args *SetWebhookArguments) (bool, error)
	DeleteWebhook                func() (bool, error)
	GetWebhookInfo               func() (*WebhookInfo, error)
	ForwardMessage               func(args *ForwardMessageArguments) (*Message, error)
	SendMessage                  func(args *SendMessageArguments) (*Message, error)
	SendPhoto                    func(args *SendPhotoArguments) (*Message, error)
	SendAudio                    func(args *SendAudioArguments) (*Message, error)
	SendDocument                 func(args *SendDocumentArguments) (*Message, error)
	SendVideo                    func(args *SendVideoArguments) (*Message, error)
	SendVoice                    func(args *SendVoiceArguments) (*Message, error)
	SendVideoNote                func(args *SendVideoNoteArguments) (*Message, error)
	SendMediaGroup               func(args *SendMediaGroupArguments) (*Message, error)
	SendLocation                 func(args *SendLocationArguments) (*Message, error)
	SendVenue                    func(args *SendVenueArguments) (*Message, error)
	SendContact                  func(args *SendContentArguments) (*Message, error)
	SendChatAction               func(args *SendChatActionArguments) (bool, error)
	GetUserProfilePhotos         func(args *GetUserProfilePhotosArguments) (*UserProfilePhotos, error)
	GetFile                      func(args *GetFileArguments) (*File, error)
	KickChatMember               func(args *KickChatMemberArguments) (bool, error)
	UnbanChatMember              func(args *ChatMemberArguments) (bool, error)
	RestrictChatMember           func(args *RestrictChatMemberArguments) (bool, error)
	PromoteChatMember            func(args *PromoteChatMemberArguments) (bool, error)
	ExportChatInviteLink         func(args *ChatArguments) (*string, error)
	SetChatPhoto                 func(args *SetChatPhotoArguments) (bool, error)
	DeleteChatPhoto              func(args *ChatArguments) (bool, error)
	SetChatTitle                 func(args *SetChatTitleArguments) (bool, error)
	SetChatDescription           func(args *SetChatDescriptionArguments) (bool, error)
	PinChatMessage               func(args *PinChatMessageArguments) (bool, error)
	UnpinChatMessage             func(args *ChatArguments) (bool, error)
	LeaveChat                    func(args *ChatArguments) (bool, error)
	GetChat                      func(args *ChatArguments) (*Chat, error)
	GetChatAdministrators        func(args *ChatArguments) ([]ChatMember, error)
	GetChatMembersCount          func(args *ChatArguments) (int, error)
	GetChatMember                func(args *ChatMemberArguments) (*ChatMember, error)
	SendSticker                  func(args *SendStickerArguments) (*Message, error)
	GetStickerSet                func(args *GetStickerSetArguments) (*StickerSet, error)
	UploadStickerFile            func(args *UploadStickerFileArguments) (*File, error)
	CreateNewStickerSet          func(args *CreateNewStickerSetArguments) (bool, error)
	AddStickerToSet              func(args *AddStickerToSetArguments) (bool, error)
	SetStickerPositionInSet      func(args *SetStickerPositionInSetArguments) (bool, error)
	DeleteStickerFromSet         func(args *DeleteStickerFromSetArguments) (bool, error)
	DeleteMessage                func(args *DeleteMessageArguments) (bool, error)
	AnswerCallbackQuery          func(args *AnwserCallbackQueryArguments) (bool, error)
	EditMessageText              func(args *EditMessageTextArguments) (*Message, error)
	EditMessageCaption           func(args *EditMessageCaptionArguments) (*Message, error)
	EditMessageReplyMarkup       func(args *EditMessageReplyMarkupArguments) (*Message, error)
	EditOthersMessageText        func(args *EditMessageTextArguments) (bool, error)        `method:"editMessageText"`
	EditOthersMessageCaption     func(args *EditMessageCaptionArguments) (bool, error)     `method:"editMessageCaption"`
	EditOthersMessageReplyMarkup func(args *EditMessageReplyMarkupArguments) (bool, error) `method:"editMessageReplyMarkup"`
}

func New(token string) *Telegroid {
	tg := new(Telegroid)
	bindInvoker(tg, token)
	return tg
}
