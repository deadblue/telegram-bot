package arguments

import (
	"fmt"
	"io"
)


type ChatArgs struct {
	_BasicArgs
}
func (a *ChatArgs) ChatId(chatId int) {
	a.withInt("chat_id", chatId)
}
func (a *ChatArgs) Channel(channelName string) {
	a.withString("chat_id", fmt.Sprintf("@%s", channelName))
}


type ChatMemberArgs struct {
	ChatArgs
}
func (a *ChatMemberArgs) UserId(userId int) {
	a.withInt("user_id", userId)
}


type KickChatMemberArgs struct {
	ChatMemberArgs
}
func (a *KickChatMemberArgs) UntilDate(timestamp int64) {
	a.withInt64("until_date", timestamp)
}


type RestrictChatMemberArgs struct {
	KickChatMemberArgs
}
func (a *RestrictChatMemberArgs) SendMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_send_messages",
	}
}
func (a *RestrictChatMemberArgs) SendMediaMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_send_media_messages",
	}
}
func (a *RestrictChatMemberArgs) SendOtherMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_send_other_messages",
	}
}
func (a *RestrictChatMemberArgs) AddWebPagePreviews() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_add_web_page_previews",
	}
}


type PromoteChatMemberArgs struct {
	ChatMemberArgs
}
func (a *PromoteChatMemberArgs) ChangeInfo() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_change_info",
	}
}
func (a *PromoteChatMemberArgs) PostMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_post_messages",
	}
}
func (a *PromoteChatMemberArgs) EditMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_edit_messages",
	}
}
func (a *PromoteChatMemberArgs) DeleteMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_delete_messages",
	}
}
func (a *PromoteChatMemberArgs) InviteUsers() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_invite_users",
	}
}
func (a *PromoteChatMemberArgs) RestrictMembers() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_restrict_members",
	}
}
func (a *PromoteChatMemberArgs) PinMessages() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_pin_messages",
	}
}
func (a *PromoteChatMemberArgs) PromoteMembers() Switch {
	return &_ParameterSwitch{
		holder: a._BasicArgs,
		name:   "can_promote_members",
	}
}


type SetChatPhotoArgs struct {
	ChatArgs
}
func (a *SetChatPhotoArgs) Photo(filename string, filedata io.Reader) {
	a.withFile("photo", filename, filedata)
}


type SetChatTitleArgs struct {
	ChatArgs
}
func (a *SetChatTitleArgs) Title(title string) {
	a.withString("title", title)
}


type SetChatDescriptionArgs struct {
	ChatArgs
}
func (a *SetChatDescriptionArgs) Description(description string) {
	a.withString("description", description)
}


type PinChatMessageArgs struct {
	ChatArgs
}
func (a *PinChatMessageArgs) MessageId(messageId int) {
	a.withInt("message_id", messageId)
}
func (a *PinChatMessageArgs) DisableNotification() {
	a.withBool("disable_notification", true)
}


type SetChatStickerSetArgs struct {
	ChatArgs
}
func (a *SetChatStickerSetArgs) StickerSet(name string) {
	a.withString("sticker_set_name", name)
}
