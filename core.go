package telegroid

import (
	"net/http"
)

type Telegroid struct {
	token  string
	client *http.Client

	GetMe                   func() *User                                                `method:"getMe"`
	GetUpdates              func(request GetUpdatesArguments) *[]Update                 `method:"getUpdates"`
	SetWebhook              func(args SetWebhookArguments) bool                         `method:"setWebhook"`
	DeleteWebhook           func() bool                                                 `method:"deleteWebhook"`
	GetWebhookInfo          func() *WebhookInfo                                         `method:"getWebhookInfo"`
	ForwardMessage          func(args ForwardMessageArguments) *Message                 `method:"forwardMessage"`
	SendMessage             func(args SendMessageArguments) *Message                    `method:"sendMessage"`
	SendPhoto               func(args SendPhotoArguments) *Message                      `method:"sendPhoto"`
	SendAudio               func(args SendAudioArguments) *Message                      `method:"sendAudio"`
	SendDocument            func(args SendDocumentArguments) *Message                   `method:"sendDocument"`
	SendVideo               func(args SendVideoArguments) *Message                      `method:"sendVideo"`
	SendVoice               func(args SendVoiceArguments) *Message                      `method:"sendVoice"`
	SendVideoNote           func(args SendVideoNoteArguments) *Message                  `method:"sendVideoNote"`
	SendMediaGroup          func(args SendMediaGroupArguments) *Message                 `method:"sendMediaGroup"`
	SendLocation            func(args SendLocationArguments) *Message                   `method:"sendLocation"`
	SendVenue               func(args SendVenueArguments) *Message                      `method:"sendVenue"`
	SendContact             func(args SendContentArguments) *Message                    `method:"sendContact"`
	SendChatAction          func(args SendChatActionArguments) bool                     `method:"sendChatAction"`
	GetUserProfilePhotos    func(args GetUserProfilePhotosArguments) *UserProfilePhotos `method:"getUserProfilePhotos"`
	GetFile                 func(args GetFileArguments) *File                           `method:"getFile"`
	KickChatMember          func(args KickChatMemberArguments) bool                     `method:"kickChatMember"`
	UnbanChatMember         func(args ChatMemberArguments) bool                         `method:"unbanChatMember"`
	RestrictChatMember      func(args RestrictChatMemberArguments) bool                 `method:"restrictChatMember"`
	PromoteChatMember       func(args PromoteChatMemberArguments) bool                  `method:"promoteChatMember"`
	ExportChatInviteLink    func(args ChatArguments) *string                            `method:"exportChatInviteLink"`
	SetChatPhoto            func(args SetChatPhotoArguments) bool                       `method:"setChatPhoto"`
	DeleteChatPhoto         func(args ChatArguments) bool                               `method:"deleteChatPhoto"`
	SetChatTitle            func(args SetChatTitleArguments) bool                       `method:"setChatTitle"`
	SetChatDescription      func(args SetChatDescriptionArguments) bool                 `method:"setChatDescription"`
	PinChatMessage          func(args PinChatMessageArguments) bool                     `method:"pinChatMessage"`
	UnpinChatMessage        func(args ChatArguments) bool                               `method:"unpinChatMessage"`
	LeaveChat               func(args ChatArguments) bool                               `method:"leaveChat"`
	GetChat                 func(args ChatArguments) *Chat                              `method:"getChat"`
	GetChatAdministrators   func(args ChatArguments) *[]ChatMember                      `method:"getChatAdministrators"`
	GetChatMembersCount     func(args ChatArguments) int                                `method:"getChatMembersCount"`
	GetChatMember           func(args ChatMemberArguments) *ChatMember                  `method:"getChatMember"`
	SendSticker             func(args SendStickerArguments) *Message                    `method:"sendSticker"`
	GetStickerSet           func(args GetStickerSetArguments) *StickerSet               `method:"getStickerSet"`
	UploadStickerFile       func(args UploadStickerFileArguments) *File                 `method:"uploadStickerFile"`
	CreateNewStickerSet     func(args CreateNewStickerSetArguments) bool                `method:"createNewStickerSet"`
	AddStickerToSet         func(args AddStickerToSetArguments) bool                    `method:"addStickerToSet"`
	SetStickerPositionInSet func(args SetStickerPositionInSetArguments) bool            `method:"setStickerPositionInSet"`
	DeleteStickerFromSet    func(args DeleteStickerFromSetArguments) bool               `method:"deleteStickerFromSet"`
	DeleteMessage           func(args DeleteMessageArguments) bool                      `method:"deleteMessage"`
}

func NewTelegroid(token string) *Telegroid {
	var td = &Telegroid{
		token:  token,
		client: &http.Client{},
	}
	td.bindInvoker()
	return td
}
