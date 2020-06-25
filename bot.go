package telegram

import (
	"github.com/deadblue/telegram-bot/parameters"
	"github.com/deadblue/telegram-bot/types"
)

// Telegram Bot API wrapper
type Bot struct {

	// Get basic information about the bot.
	GetMe func() (*types.User, error)
	// Change the list of the bot's commands.
	SetMyCommands func(params *parameters.SetMyCommandsParams) (bool, error)
	// Get the current list of the bot's commands.
	GetMyCommands func() ([]*types.BotCommand, error)

	// Receive incoming updates using long polling.
	GetUpdate func(params *parameters.GetUpdateParams) ([]*types.Update, error)
	// Specify a url and receive incoming updates via an outgoing webhook.
	SetWebhook func(params *parameters.SetWebhookParams) (bool, error)
	// Remove webhook integration.
	DeleteWebhook func() (bool, error)
	// Get current webhook status.
	GetWebhookInfo func() (*types.WebhookInfo, error)

	// Send text message.
	SendMessage func(params *parameters.SendMessageParams) (*types.Message, error)
	// Send photo.
	SendPhoto func(params *parameters.SendPhotoParams) (*types.Message, error)
	// Send audio file, must be in the mp3 or m4a format, up to 50 MB in size.
	SendAudio func(params *parameters.SendAudioParams) (*types.Message, error)
	// Send general file, up to 50 MB in size.
	SendDocument func(params *parameters.SendDocumentParams) (*types.Message, error)
	// Send video file, must be mp4 format, up to 50 MB in size.
	SendVideo func(params *parameters.SendVideoParams) (*types.Message, error)
	// Send animation file (GIF or H.264 video without sound), up to 50 MB in size.
	SendAnimation func(params *parameters.SendAnimationParams) (*types.Message, error)
	// Send voice message in ogg format, up to 50 MB in size.
	SendVoice func(params *parameters.SendVoiceParams) (*types.Message, error)
	// Send rounded square mp4 video message, up to 1 minute long.
	SendVideoNote func(params *parameters.SendVideoNoteParams) (*types.Message, error)
	// Send a group of photos or videos as an album.
	SendMediaGroup func(params *parameters.SendMediaGroupParams) ([]*types.Message, error)
	// Send point on the map.
	SendLocation func(params *parameters.SendLocationParams) (*types.Message, error)
	// Send information about a venue.
	SendVenue func(params *parameters.SendVenueParams) (*types.Message, error)
	// Send phone contacts.
	SendContact func(params *parameters.SendContactParams) (*types.Message, error)
	// Send a native poll.
	SendPoll func(params *parameters.SendPollParams) (*types.Message, error)
	// Send an animated emoji that will display a random value.
	SendDice func(params *parameters.SendDiceParams) (*types.Message, error)
	// Forward messages of any kind.
	ForwardMessage func(params *parameters.ForwardMessageParams) (*types.Message, error)
	// Tell the user that something is happening on the bot's side.
	SendChatAction func(params *parameters.SendChatActionParams) (bool, error)

	// Get a list of profile pictures for a user.
	GetUserProfilePhotos func(params *parameters.GetUserProfilePhotosParams) (*types.UserProfilePhotos, error)
	// Get basic info about a file and prepare it for downloading.
	GetFile func(params *parameters.FileParams) (*types.File, error)

	// Edit text and game messages.
	EditMessageText func(params *parameters.EditMessageTextParams) (*types.EditMessageResult, error)
	// Edit captions of messages.
	EditMessageCaption func(params *parameters.EditMessageCaptionParams) (*types.EditMessageResult, error)
	// Edit animation, audio, document, photo, or video messages.
	EditMessageMedia func(params *parameters.EditMessageMediaParams) (*types.EditMessageResult, error)
	// Edit only the reply markup of messages.
	EditMessageReplyMarkup func(params *parameters.EditMessageReplyMarkupParams) (*types.EditMessageResult, error)
	// Stop a poll which was sent by the bot.
	StopPoll func(params *parameters.StopPollParams) (*types.Poll, error)
	// Delete a message, including service messages.
	DeleteMessage func(params *parameters.DeleteMessageParams) (bool, error)

	// Edit live location message.
	EditMessageLiveLocation func(params *parameters.EditMessageLiveLocation) (*types.EditMessageResult, error)
	// Stop updating a live location message before it expires.
	StopMessageLiveLocation func(params *parameters.StopMessageLiveLocation) (*types.EditMessageResult, error)

	// Kick a user from a group, a supergroup or a channel.
	KickChatMember func(params *parameters.KickChatMemberParams) (bool, error)
	// Unban a previously kicked user in a supergroup or channel.
	UnbanChatMember func(params *parameters.ChatMemberParams) (bool, error)
	// Restrict a user in a supergroup.
	RestrictChatMember func(params *parameters.RestrictChatMemberParams) (bool, error)
	// Promote or demote a user in a supergroup or a channel.
	PromoteChatMember func(params *parameters.PromoteChatMemberParams) (bool, error)
	// Set a custom title for an administrator in a supergroup promoted by the bot.
	SetChatAdministratorCustomTitle func(params *parameters.SetChatAdministratorCustomTitleParams) (bool, error)
	// Set default chat permissions for all members.
	SetChatPermissions func(params *parameters.SetChatPermissionsParams) (bool, error)
	// Generate a new invite link for a chat, any previously generated link is revoked.
	ExportChatInviteLink func(params *parameters.ChatParams) (string, error)
	// Set a new profile photo for the chat.
	SetChatPhoto func(params *parameters.SetChatPhotoParams) (bool, error)
	// Delete a chat photo.
	DeleteChatPhoto func(params *parameters.ChatParams) (bool, error)
	// Change the title of a chat.
	SetChatTitle func(params *parameters.SetChatTitleParams) (bool, error)
	// Change the description of a group, a supergroup or a channel.
	SetChatDescription func(params *parameters.SetChatDescriptionParams) (bool, error)
	// Pin a message in a group, a supergroup, or a channel.
	PinChatMessage func(params *parameters.PinChatMessageParams) (bool, error)
	// Unpin a message in a group, a supergroup, or a channel.
	UnpinChatMessage func(params *parameters.ChatParams) (bool, error)
	// Leave a group, supergroup or channel.
	LeaveChat func(params *parameters.ChatParams) (bool, error)
	// Get up to date information about the chat.
	GetChat func(params *parameters.ChatParams) (*types.Chat, error)
	// Get a list of administrators in a chat.
	GetChatAdministrators func(params *parameters.ChatParams) ([]*types.ChatMember, error)
	// Get the number of members in a chat.
	GetChatMembersCount func(params *parameters.ChatParams) (int, error)
	// Get information about a member of a chat.
	GetChatMember func(params *parameters.ChatMemberParams) (*types.ChatMember, error)
	// Set a new group sticker set for a supergroup.
	SetChatStickerSet func(params *parameters.SetChatStickerSetParams) (bool, error)
	// Delete a group sticker set from a supergroup.
	DeleteChatStickerSet func(params *parameters.ChatParams) (bool, error)

	// Send answers to callback queries sent from inline keyboards.
	AnswerCallbackQuery func(params *parameters.AnswerCallbackQueryParams) (bool, error)

	// Send static .WEBP or animated .TGS stickers.
	SendSticker func(params *parameters.SendStickerParams) (*types.Message, error)
	// Get a sticker set.
	GetStickerSet func(params *parameters.GetStickerSetParams) (*types.StickerSet, error)
	// Upload a .PNG file with a sticker for later use.
	UploadStickerFile func(params *parameters.UploadStickerFileParams) (*types.File, error)
	// Create a new sticker set owned by a user.
	CreateNewStickerSet func(params *parameters.CreateNewStickerSetParams) (bool, error)
	// Add a new sticker to a set created by the bot.
	AddStickerToSet func(params *parameters.AddStickerToSetParams) (bool, error)
	// Move a sticker in a set created by the bot to a specific position.
	SetStickerPositionInSet func(params *parameters.SetStickerPositionInSetParams) (bool, error)
	// Delete a sticker from a set created by the bot.
	DeleteStickerFromSet func(params *parameters.DeleteStickerFromSetParams) (bool, error)
	// Set the thumbnail of a sticker set.
	SetStickerSetThumb func(params *parameters.SetStickerSetThumbParams) (bool, error)

	// Send a game
	SendGame func(params *parameters.SendGameParams) (*types.Message, error)
	// Set the score of the specified user in a game.
	SetGameScore func(params *parameters.SetGameScoreParams) (*types.EditMessageResult, error)
	// Get data for high score tables.
	GetGameHighScores func(params *parameters.GetGameHighScoresParams) ([]*types.GameHighScore, error)

	// Send answers to an inline query.
	AnswerInlineQuery func(params *parameters.AnswerInlineQueryParams) (bool, error)

	// Send invoice.
	SendInvoice func(params *parameters.SendInvoiceParams) (*types.Message, error)
	// Answer the shipping_query from update.
	AnswerShippingQuery func(params *parameters.AnswerShippingQueryParams) (bool, error)
	// Answer the final confirmation from update.
	AnswerPreCheckoutQuery func(params *parameters.AnswerPreCheckoutQueryParams) (bool, error)
}

/*
Create a Bot with token.

There is some reflect operations during creating a bot, it is recommended
to create only one bot for each bot account, the bot is safe to use in
multiple goroutines.
*/
func New(token string) *Bot {
	bot := new(Bot)
	upUpDownDownLeftRightLeftRightBA(bot, token)
	return bot
}
