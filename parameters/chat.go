package parameters

import "time"

type ChatParams struct {
	implApiParameters
}

func (p *ChatParams) ChatId(chatId int64) {
	p.setInt64("chat_id", chatId)
}
func (p *ChatParams) ChannelName(name string) {
	p.set("chat_id", "@"+name)
}

type ChatMemberParams struct {
	ChatParams
}

func (p *ChatMemberParams) UserId(userId int) {
	p.setInt("user_id", userId)
}

type KickChatMemberParams struct {
	ChatMemberParams
}

func (p *KickChatMemberParams) UntilDate(date time.Time) {
	p.setInt64("until_date", date.Unix())
}

type RestrictChatMemberParams struct {
	KickChatMemberParams
}

// TODO

type PromoteChatMemberParams struct {
	ChatMemberParams
}

// TODO

type SetChatAdministratorCustomTitleParams struct {
	ChatMemberParams
}

// TODO

type SetChatPermissionsParams struct {
	ChatParams
}

// TODO

type SetChatPhotoParams struct {
	ChatParams
}

// TODO

type SetChatTitleParams struct {
	ChatParams
}

func (p *SetChatTitleParams) Title(title string) {
	p.set("title", title)
}

type SetChatDescriptionParams struct {
	ChatParams
}

func (p *SetChatDescriptionParams) Description(description string) {
	p.set("description", description)
}

type PinChatMessageParams struct {
	ChatParams
}

func (p *PinChatMessageParams) MessageId(messageId int) {
	p.setInt("message_id", messageId)
}
func (p *PinChatMessageParams) DisableNotification() {
	p.setBool("disable_notification", true)
}

type SetChatStickerSetParams struct {
	ChatParams
}

func (p *SetChatStickerSetParams) StickerSetName(name string) {
	p.set("sticker_set_name", name)
}
