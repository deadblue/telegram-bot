package telegroid

import (
	"github.com/deadblue/telegroid/arguments"
	"net/http"
)

// Telegram Bot API wrapper
type Bot struct {
	// http client
	hc *http.Client

	// Get me
	GetMe func() (*User, error)
	// Get updates
	GetUpdates func(args *arguments.GetUpdatesArgs) ([]Update, error)
	// Webhook management
	SetWebhook     func(args *arguments.SetWebhookArgs) (bool, error)
	DeleteWebhook  func() (bool, error)
	GetWebhookInfo func() (*WebhookInfo, error)
	// Send message
	ForwardMessage func(args *arguments.ForwardMessageArgs) (*Message, error)
	SendMessage    func(args *arguments.SendMessageArgs) (*Message, error)
	SendPhoto      func(args *arguments.SendPhotoArgs) (*Message, error)
	SendAudio      func(args *arguments.SendAudioArgs) (*Message, error)
	SendDocument   func(args *arguments.SendDocumentArgs) (*Message, error)
	SendVideo      func(args *arguments.SendVideoArgs) (*Message, error)
	SendAnimation  func(args *arguments.SendAnimationArgs) (*Message, error)
	SendVoice      func(args *arguments.SendVoiceArgs) (*Message, error)
	SendVideoNote  func(args *arguments.SendVideoNoteArgs) (*Message, error)
	SendMediaGroup func(args *arguments.SendMediaGroupArgs) ([]*Message, error)
	SendLocation   func(args *arguments.SendLocationArgs) (*Message, error)
	SendVenue      func(args *arguments.SendVenueArgs) (*Message, error)
	SendContact    func(args *arguments.SendContactArgs) (*Message, error)
	SendPoll       func(args *arguments.SendPollArgs) (*Message, error)
	SendChatAction func(args *arguments.SendChatActionArgs) (bool, error)
	// Inline keyboard callback
	AnswerCallbackQuery func(args *arguments.AnswerCallbackQueryArgs) (bool, error)
	// Delete message
	DeleteMessage func(args *arguments.ChatMessageArgs) (bool, error)
	// Edit own message
	EditMessageText         func(args *arguments.EditMessageTextArgs) (*Message, error)
	EditMessageCaption      func(args *arguments.EditMessageCaptionArgs) (*Message, error)
	EditMessageMedia        func(args *arguments.EditMessageMediaArgs) (*Message, error)
	EditMessageReplyMarkup  func(args *arguments.EditMessageReplyMarkupArgs) (*Message, error)
	EditMessageLiveLocation func(args *arguments.EditMessageLiveLocationArgs) (*Message, error)
	StopMessageLiveLocation func(args *arguments.EditMessageReplyMarkupArgs) (*Message, error)
	StopPoll                func(args *arguments.StopPollArgs) (*Poll, error)
	// Edit other's message
	EditOthersMessageText         func(args *arguments.EditMessageTextArgs) (bool, error)         `method:"editMessageText"`
	EditOthersMessageCaption      func(args *arguments.EditMessageCaptionArgs) (bool, error)      `method:"editMessageCaption"`
	EditOthersMessageMedia        func(args *arguments.EditMessageMediaArgs) (bool, error)        `method:"editMessageMedia"`
	EditOthersMessageReplyMarkup  func(args *arguments.EditMessageReplyMarkupArgs) (bool, error)  `method:"editMessageReplyMarkup"`
	EditOthersMessageLiveLocation func(args *arguments.EditMessageLiveLocationArgs) (bool, error) `method:"editMessageLiveLocation"`
	StopOthersMessageLiveLocation func(args *arguments.EditMessageReplyMarkupArgs) (bool, error)  `method:"stopMessageLiveLocation"`
	// User information
	GetUserProfilePhotos func(args *arguments.GetUserProfilePhotosArgs) (*UserProfilePhotos, error)
	// Chat information
	GetChat               func(args *arguments.ChatArgs) (*Chat, error)
	GetChatAdministrators func(args *arguments.ChatArgs) ([]ChatMember, error)
	GetChatMembersCount   func(args *arguments.ChatArgs) (int, error)
	GetChatMember         func(args *arguments.ChatMemberArgs) (*ChatMember, error)
	// Chat management
	KickChatMember       func(args *arguments.KickChatMemberArgs) (bool, error)
	UnbanChatMember      func(args *arguments.ChatMemberArgs) (bool, error)
	RestrictChatMember   func(args *arguments.RestrictChatMemberArgs) (bool, error)
	PromoteChatMember    func(args *arguments.PromoteChatMemberArgs) (bool, error)
	ExportChatInviteLink func(args *arguments.ChatArgs) (string, error)
	SetChatPhoto         func(args *arguments.SetChatPhotoArgs) (bool, error)
	DeleteChatPhoto      func(args *arguments.ChatArgs) (bool, error)
	SetChatTitle         func(args *arguments.SetChatTitleArgs) (bool, error)
	SetChatDescription   func(args *arguments.SetChatDescriptionArgs) (bool, error)
	PinChatMessage       func(args *arguments.PinChatMessageArgs) (bool, error)
	UnpinChatMessage     func(args *arguments.ChatArgs) (bool, error)
	SetChatStickerSet    func(args *arguments.SetChatStickerSetArgs) (bool, error)
	DeleteChatStickerSet func(args *arguments.ChatArgs) (bool, error)
	LeaveChat            func(args *arguments.ChatArgs) (bool, error)
	// Misc
	GetFile func(args *arguments.GetFileArgs) (*File, error)
	// Sticker
	SendSticker             func(args *arguments.SendStickerArgs) (*Message, error)
	GetStickerSet           func(args *arguments.GetStickerSetArgs) (*StickerSet, error)
	UploadStickerFile       func(args *arguments.UploadStickerFileArgs) (*File, error)
	CreateNewStickerSet     func(args *arguments.CreateNewStickerSetArgs) (bool, error)
	AddStickerToSet         func(args *arguments.AddStickerToSetArgs) (bool, error)
	SetStickerPositionInSet func(args *arguments.SetStickerPositionInSetArgs) (bool, error)
	DeleteStickerFromSet    func(args *arguments.DeleteStickerFromSetArgs) (bool, error)
	// Game
	SendGame          func(args *arguments.SendGameArgs) (*Message, error)
	SetGameScore      func(args *arguments.SetGameScoreArgs) (bool, error)
	GetGameHighScores func(args *arguments.GetGameHighScoresArgs) (*GameHighScore, error)
	// Inline mode
	//AnswerInlineQuery func(args *arguments.AnswerInlineQueryArgs) (bool, error)

}

func New(token string) *Bot {
	bot := new(Bot)
	bindInvoker(bot, token)
	return bot
}
