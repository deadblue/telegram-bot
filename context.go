package telegram

import (
	"context"
	"github.com/deadblue/telegram-bot/internal/core"
	"github.com/deadblue/telegram-bot/parameters"
	"github.com/deadblue/telegram-bot/types"
)

// ContextBot wraps almost all Telegram Bot API with context.
type ContextBot struct {

	// Get basic information about the bot.
	GetMe func(ctx context.Context) (*types.User, error)
	// Change the list of the bot's commands.
	SetMyCommands func(ctx context.Context, params *parameters.SetMyCommandsParams) (bool, error)
	// Get the current list of the bot's commands.
	GetMyCommands func(ctx context.Context) ([]*types.BotCommand, error)

	// Receive incoming updates using long polling.
	GetUpdate func(ctx context.Context, params *parameters.GetUpdateParams) ([]*types.Update, error)
	// Specify a url and receive incoming updates via an outgoing webhook.
	SetWebhook func(ctx context.Context, params *parameters.SetWebhookParams) (bool, error)
	// Remove webhook integration.
	DeleteWebhook func(ctx context.Context) (bool, error)
	// Get current webhook status.
	GetWebhookInfo func(ctx context.Context) (*types.WebhookInfo, error)

	// Send text message.
	SendMessage func(ctx context.Context, params *parameters.SendMessageParams) (*types.Message, error)
	// Send photo.
	SendPhoto func(ctx context.Context, params *parameters.SendPhotoParams) (*types.Message, error)
	// Send audio file, must be in the mp3 or m4a format, up to 50 MB in size.
	SendAudio func(ctx context.Context, params *parameters.SendAudioParams) (*types.Message, error)
	// Send general file, up to 50 MB in size.
	SendDocument func(ctx context.Context, params *parameters.SendDocumentParams) (*types.Message, error)
	// Send video file, must be mp4 format, up to 50 MB in size.
	SendVideo func(ctx context.Context, params *parameters.SendVideoParams) (*types.Message, error)
	// Send animation file (GIF or H.264 video without sound), up to 50 MB in size.
	SendAnimation func(ctx context.Context, params *parameters.SendAnimationParams) (*types.Message, error)
	// Send voice message in ogg format, up to 50 MB in size.
	SendVoice func(ctx context.Context, params *parameters.SendVoiceParams) (*types.Message, error)
	// Send rounded square mp4 video message, up to 1 minute long.
	SendVideoNote func(ctx context.Context, params *parameters.SendVideoNoteParams) (*types.Message, error)
	// Send a group of photos or videos as an album.
	SendMediaGroup func(ctx context.Context, params *parameters.SendMediaGroupParams) ([]*types.Message, error)
	// Send point on the map.
	SendLocation func(ctx context.Context, params *parameters.SendLocationParams) (*types.Message, error)
	// Send information about a venue.
	SendVenue func(ctx context.Context, params *parameters.SendVenueParams) (*types.Message, error)
	// Send phone contacts.
	SendContact func(ctx context.Context, params *parameters.SendContactParams) (*types.Message, error)
	// Send a native poll.
	SendPoll func(ctx context.Context, params *parameters.SendPollParams) (*types.Message, error)
	// Send an animated emoji that will display a random value.
	SendDice func(ctx context.Context, params *parameters.SendDiceParams) (*types.Message, error)
	// Forward messages of any kind.
	ForwardMessage func(ctx context.Context, params *parameters.ForwardMessageParams) (*types.Message, error)
	// Tell the user that something is happening on the bot's side.
	SendChatAction func(ctx context.Context, params *parameters.SendChatActionParams) (bool, error)

	// Get a list of profile pictures for a user.
	GetUserProfilePhotos func(ctx context.Context, params *parameters.GetUserProfilePhotosParams) (*types.UserProfilePhotos, error)
	// Get basic info about a file and prepare it for downloading.
	GetFile func(ctx context.Context, params *parameters.FileParams) (*types.File, error)

	// Edit text and game messages.
	EditMessageText func(ctx context.Context, params *parameters.EditMessageTextParams) (*types.EditMessageResult, error)
	// Edit captions of messages.
	EditMessageCaption func(ctx context.Context, params *parameters.EditMessageCaptionParams) (*types.EditMessageResult, error)
	// Edit animation, audio, document, photo, or video messages.
	EditMessageMedia func(ctx context.Context, params *parameters.EditMessageMediaParams) (*types.EditMessageResult, error)
	// Edit only the reply markup of messages.
	EditMessageReplyMarkup func(ctx context.Context, params *parameters.EditMessageReplyMarkupParams) (*types.EditMessageResult, error)
	// Stop a poll which was sent by the bot.
	StopPoll func(ctx context.Context, params *parameters.StopPollParams) (*types.Poll, error)
	// Delete a message, including service messages.
	DeleteMessage func(ctx context.Context, params *parameters.DeleteMessageParams) (bool, error)

	// Edit live location message.
	EditMessageLiveLocation func(ctx context.Context, params *parameters.EditMessageLiveLocation) (*types.EditMessageResult, error)
	// Stop updating a live location message before it expires.
	StopMessageLiveLocation func(ctx context.Context, params *parameters.StopMessageLiveLocation) (*types.EditMessageResult, error)

	// Kick a user from a group, a supergroup or a channel.
	KickChatMember func(ctx context.Context, params *parameters.KickChatMemberParams) (bool, error)
	// Unban a previously kicked user in a supergroup or channel.
	UnbanChatMember func(ctx context.Context, params *parameters.ChatMemberParams) (bool, error)
	// Restrict a user in a supergroup.
	RestrictChatMember func(ctx context.Context, params *parameters.RestrictChatMemberParams) (bool, error)
	// Promote or demote a user in a supergroup or a channel.
	PromoteChatMember func(ctx context.Context, params *parameters.PromoteChatMemberParams) (bool, error)
	// Set a custom title for an administrator in a supergroup promoted by the bot.
	SetChatAdministratorCustomTitle func(ctx context.Context, params *parameters.SetChatAdministratorCustomTitleParams) (bool, error)
	// Set default chat permissions for all members.
	SetChatPermissions func(ctx context.Context, params *parameters.SetChatPermissionsParams) (bool, error)
	// Generate a new invite link for a chat, any previously generated link is revoked.
	ExportChatInviteLink func(ctx context.Context, params *parameters.ChatParams) (string, error)
	// Set a new profile photo for the chat.
	SetChatPhoto func(ctx context.Context, params *parameters.SetChatPhotoParams) (bool, error)
	// Delete a chat photo.
	DeleteChatPhoto func(ctx context.Context, params *parameters.ChatParams) (bool, error)
	// Change the title of a chat.
	SetChatTitle func(ctx context.Context, params *parameters.SetChatTitleParams) (bool, error)
	// Change the description of a group, a supergroup or a channel.
	SetChatDescription func(ctx context.Context, params *parameters.SetChatDescriptionParams) (bool, error)
	// Pin a message in a group, a supergroup, or a channel.
	PinChatMessage func(ctx context.Context, params *parameters.PinChatMessageParams) (bool, error)
	// Unpin a message in a group, a supergroup, or a channel.
	UnpinChatMessage func(ctx context.Context, params *parameters.ChatParams) (bool, error)
	// Leave a group, supergroup or channel.
	LeaveChat func(ctx context.Context, params *parameters.ChatParams) (bool, error)
	// Get up to date information about the chat.
	GetChat func(ctx context.Context, params *parameters.ChatParams) (*types.Chat, error)
	// Get a list of administrators in a chat.
	GetChatAdministrators func(ctx context.Context, params *parameters.ChatParams) ([]*types.ChatMember, error)
	// Get the number of members in a chat.
	GetChatMembersCount func(ctx context.Context, params *parameters.ChatParams) (int, error)
	// Get information about a member of a chat.
	GetChatMember func(ctx context.Context, params *parameters.ChatMemberParams) (*types.ChatMember, error)
	// Set a new group sticker set for a supergroup.
	SetChatStickerSet func(ctx context.Context, params *parameters.SetChatStickerSetParams) (bool, error)
	// Delete a group sticker set from a supergroup.
	DeleteChatStickerSet func(ctx context.Context, params *parameters.ChatParams) (bool, error)

	// Send answers to callback queries sent from inline keyboards.
	AnswerCallbackQuery func(ctx context.Context, params *parameters.AnswerCallbackQueryParams) (bool, error)

	// Send static .WEBP or animated .TGS stickers.
	SendSticker func(ctx context.Context, params *parameters.SendStickerParams) (*types.Message, error)
	// Get a sticker set.
	GetStickerSet func(ctx context.Context, params *parameters.GetStickerSetParams) (*types.StickerSet, error)
	// Upload a .PNG file with a sticker for later use.
	UploadStickerFile func(ctx context.Context, params *parameters.UploadStickerFileParams) (*types.File, error)
	// Create a new sticker set owned by a user.
	CreateNewStickerSet func(ctx context.Context, params *parameters.CreateNewStickerSetParams) (bool, error)
	// Add a new sticker to a set created by the bot.
	AddStickerToSet func(ctx context.Context, params *parameters.AddStickerToSetParams) (bool, error)
	// Move a sticker in a set created by the bot to a specific position.
	SetStickerPositionInSet func(ctx context.Context, params *parameters.SetStickerPositionInSetParams) (bool, error)
	// Delete a sticker from a set created by the bot.
	DeleteStickerFromSet func(ctx context.Context, params *parameters.DeleteStickerFromSetParams) (bool, error)
	// Set the thumbnail of a sticker set.
	SetStickerSetThumb func(ctx context.Context, params *parameters.SetStickerSetThumbParams) (bool, error)

	// Send a game
	SendGame func(ctx context.Context, params *parameters.SendGameParams) (*types.Message, error)
	// Set the score of the specified user in a game.
	SetGameScore func(ctx context.Context, params *parameters.SetGameScoreParams) (*types.EditMessageResult, error)
	// Get data for high score tables.
	GetGameHighScores func(ctx context.Context, params *parameters.GetGameHighScoresParams) ([]*types.GameHighScore, error)

	// Send answers to an inline query.
	AnswerInlineQuery func(ctx context.Context, params *parameters.AnswerInlineQueryParams) (bool, error)

	// Send invoice.
	SendInvoice func(ctx context.Context, params *parameters.SendInvoiceParams) (*types.Message, error)
	// Answer the shipping_query from update.
	AnswerShippingQuery func(ctx context.Context, params *parameters.AnswerShippingQueryParams) (bool, error)
	// Answer the final confirmation from update.
	AnswerPreCheckoutQuery func(ctx context.Context, params *parameters.AnswerPreCheckoutQueryParams) (bool, error)
}

/*
NewWithContext creates a ContextBot with token.

There is some reflect operations during creating a bot, it is recommended
to create only one bot for each bot account, the bot is safe to use in
multiple goroutines.
*/
func NewWithContext(token string) *ContextBot {
	bot := new(ContextBot)
	core.Bind(bot, true, token)
	return bot
}
