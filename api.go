//
package telegroid

import "github.com/deadblue/telegroid/params"

// The bot API wrapper
type Bot struct {
	// Get me
	GetMe func() (*User, error)
	// Get updates
	GetUpdates func(params *params.GetUpdatesParameters) ([]Update, error)
	// Webhook management
	SetWebhook     func(params *params.SetWebhookParameters) (bool, error)
	DeleteWebhook  func() (bool, error)
	GetWebhookInfo func() (*WebhookInfo, error)
	// Send message
	ForwardMessage func(params *params.ForwardMessageParameters) (*Message, error)
	SendMessage    func(params *params.SendMessageParameters) (*Message, error)
	SendPhoto      func(params *params.SendPhotoParameters) (*Message, error)
	SendAudio      func(params *params.SendAudioParameters) (*Message, error)
	SendDocument   func(params *params.SendDocumentParameters) (*Message, error)
	SendVideo      func(params *params.SendVideoParameters) (*Message, error)
	SendAnimation  func(params *params.SendAnimationParameters) (*Message, error)
	SendVoice      func(params *params.SendVoiceParameters) (*Message, error)
	SendVideoNote  func(params *params.SendVideoNoteParameters) (*Message, error)
	SendMediaGroup func(params *params.SendMediaGroupParameters) (*Message, error)
	SendLocation   func(params *params.SendLocationParameters) (*Message, error)
	SendVenue      func(params *params.SendVenueParameters) (*Message, error)
	SendContact    func(params *params.SendContactParameters) (*Message, error)
	SendChatAction func(params *params.SendChatActionParameters) (bool, error)
	// Inline Callback
	AnswerCallbackQuery func(params *params.AnswerCallbackQueryParameters) (bool, error)
	// Delete message
	DeleteMessage func(params *params.ChatMessageParameters) (bool, error)
	// Edit own message
	EditMessageText         func(params *params.EditMessageTextParameters) (*Message, error)
	EditMessageCaption      func(params *params.EditMessageCaptionParameters) (*Message, error)
	EditMessageMedia        func(params *params.EditMessageMediaParameters) (*Message, error)
	EditMessageReplyMarkup  func(params *params.EditMessageReplyMarkupParameters) (*Message, error)
	EditMessageLiveLocation func(params *params.EditMessageLiveLocationParameters) (*Message, error)
	StopMessageLiveLocation func(params *params.EditMessageReplyMarkupParameters) (*Message, error)
	// Edit other's message
	EditOthersMessageText         func(params *params.EditMessageTextParameters) (bool, error)         `method:"editMessageText"`
	EditOthersMessageCaption      func(params *params.EditMessageCaptionParameters) (bool, error)      `method:"editMessageCaption"`
	EditOthersMessageMedia        func(params *params.EditMessageMediaParameters) (bool, error)        `method:"editMessageMedia"`
	EditOthersMessageReplyMarkup  func(params *params.EditMessageReplyMarkupParameters) (bool, error)  `method:"editMessageReplyMarkup"`
	EditOthersMessageLiveLocation func(params *params.EditMessageLiveLocationParameters) (bool, error) `method:"editMessageLiveLocation"`
	StopOthersMessageLiveLocation func(params *params.EditMessageReplyMarkupParameters) (bool, error)  `method:"stopMessageLiveLocation"`
	// User information
	GetUserProfilePhotos func(params *params.GetUserProfilePhotosParameters) (*UserProfilePhotos, error)
	// Chat information
	GetChat               func(params *params.ChatParameters) (*Chat, error)
	GetChatAdministrators func(params *params.ChatParameters) ([]ChatMember, error)
	GetChatMembersCount   func(params *params.ChatParameters) (int, error)
	GetChatMember         func(params *params.ChatMemberParameters) (*ChatMember, error)
	// Chat management
	KickChatMember       func(params *params.KickChatMemberParameters) (bool, error)
	UnbanChatMember      func(params *params.ChatMemberParameters) (bool, error)
	RestrictChatMember   func(params *params.RestrictChatMemberParameters) (bool, error)
	PromoteChatMember    func(params *params.PromoteChatMemberParameters) (bool, error)
	ExportChatInviteLink func(params *params.ChatParameters) (string, error)
	SetChatPhoto         func(params *params.SetChatPhotoParameters) (bool, error)
	DeleteChatPhoto      func(params *params.ChatParameters) (bool, error)
	SetChatTitle         func(params *params.SetChatTitleParameters) (bool, error)
	SetChatDescription   func(params *params.SetChatDescriptionParameters) (bool, error)
	PinChatMessage       func(params *params.PinChatMessageParameters) (bool, error)
	UnpinChatMessage     func(params *params.ChatParameters) (bool, error)
	SetChatStickerSet    func(params *params.SetChatStickerSetParameters) (bool, error)
	DeleteChatStickerSet func(params *params.ChatParameters) (bool, error)
	LeaveChat            func(params *params.ChatParameters) (bool, error)
	// Misc
	GetFile func(params *params.GetFileParameters) (*File, error)
	// Sticker
	SendSticker             func(params *params.SendStickerParameters) (*Message, error)
	GetStickerSet           func(params *params.GetStickerSetParameters) (*StickerSet, error)
	UploadStickerFile       func(params *params.UploadStickerFileParameters) (*File, error)
	CreateNewStickerSet     func(params *params.CreateNewStickerSetParameters) (bool, error)
	AddStickerToSet         func(params *params.AddStickerToSetParameters) (bool, error)
	SetStickerPositionInSet func(params *params.SetStickerPositionInSetParameters) (bool, error)
	DeleteStickerFromSet    func(params *params.DeleteStickerFromSetParameters) (bool, error)
	// Inline mode
	AnswerInlineQuery func(params *params.AnswerInlineQueryParameters) (bool, error)
	// Game
	SendGame          func(params *params.SendGameParameters) (*Message, error)
	SetGameScore      func(params *params.SetGameScoreParameters) (bool, error)
	GetGameHighScores func(params *params.GetGameHighScoresParameters) (*GameHighScore, error)
}

func New(token string) *Bot {
	tg := new(Bot)
	bindInvoker(tg, token)
	return tg
}
