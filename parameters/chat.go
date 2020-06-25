package parameters

import (
	"github.com/deadblue/telegram-bot/types"
	"time"
)

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

func (p *RestrictChatMemberParams) Permissions(permission *types.ChatPermissions) {
	p.setJson("permissions", permission)
}

type PromoteChatMemberParams struct {
	ChatMemberParams
}

func (p *PromoteChatMemberParams) CanChangeInfo() {
	p.setBool("can_change_info", true)
}
func (p *PromoteChatMemberParams) CanPostMessage() {
	p.setBool("can_post_messages", true)
}
func (p *PromoteChatMemberParams) CanEditMessages() {
	p.setBool("can_edit_messages", true)
}
func (p *PromoteChatMemberParams) CanDeleteMessages() {
	p.setBool("can_delete_messages", true)
}
func (p *PromoteChatMemberParams) CanInviteUsers() {
	p.setBool("can_invite_users", true)
}
func (p *PromoteChatMemberParams) CanRestrictMembers() {
	p.setBool("can_restrict_members", true)
}
func (p *PromoteChatMemberParams) CanPinMessages() {
	p.setBool("can_pin_messages", true)
}
func (p *PromoteChatMemberParams) CanPromoteMembers() {
	p.setBool("can_promote_members", true)
}

type SetChatAdministratorCustomTitleParams struct {
	ChatMemberParams
}

func (p *SetChatAdministratorCustomTitleParams) CustomTitle(title string) {
	p.set("custom_title", title)
}

type SetChatPermissionsParams struct {
	ChatParams
}

func (p *SetChatPermissionsParams) Permissions(permission *types.ChatPermissions) {
	p.setJson("permissions", permission)
}

type SetChatPhotoParams struct {
	ChatParams
}

func (p *SetChatPhotoParams) Photo(file *InputFile) {
	p.setFile("photo", file)
}

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
