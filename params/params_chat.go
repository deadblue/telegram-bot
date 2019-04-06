package params

import (
	"fmt"
	"io"
)


type ChatParameters struct {
	_BasicParameters
}
func (p *ChatParameters) ChatId(chatId int) {
	p.withInt("chat_id", chatId)
}
func (p *ChatParameters) Channel(channelName string) {
	p.withString("chat_id", fmt.Sprintf("@%s", channelName))
}


type ChatMemberParameters struct {
	ChatParameters
}
func (p *ChatMemberParameters) UserId(userId int) {
	p.withInt("user_id", userId)
}


type KickChatMemberParameters struct {
	ChatMemberParameters
}
func (p *KickChatMemberParameters) UntilDate(timestamp int64) {
	p.withInt64("until_date", timestamp)
}


type _ParameterSwitch struct {
	holder _BasicParameters
	name   string
}
func (p *_ParameterSwitch) On() {
	p.holder.withBool(p.name, true)
}
func (p *_ParameterSwitch) Off() {
	p.holder.withBool(p.name, false)
}


type RestrictChatMemberParameters struct {
	KickChatMemberParameters
}
func (p *RestrictChatMemberParameters) SendMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_send_messages",
	}
}
func (p *RestrictChatMemberParameters) SendMediaMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_send_media_messages",
	}
}
func (p *RestrictChatMemberParameters) SendOtherMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_send_other_messages",
	}
}
func (p *RestrictChatMemberParameters) AddWebPagePreviews() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_add_web_page_previews",
	}
}


type PromoteChatMemberParameters struct {
	ChatMemberParameters
}
func (p *PromoteChatMemberParameters) ChangeInfo() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_change_info",
	}
}
func (p *PromoteChatMemberParameters) PostMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_post_messages",
	}
}
func (p *PromoteChatMemberParameters) EditMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_edit_messages",
	}
}
func (p *PromoteChatMemberParameters) DeleteMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_delete_messages",
	}
}
func (p *PromoteChatMemberParameters) InviteUsers() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_invite_users",
	}
}
func (p *PromoteChatMemberParameters) RestrictMembers() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_restrict_members",
	}
}
func (p *PromoteChatMemberParameters) PinMessages() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_pin_messages",
	}
}
func (p *PromoteChatMemberParameters) PromoteMembers() Switch {
	return &_ParameterSwitch{
		holder: p._BasicParameters,
		name: "can_promote_members",
	}
}


type SetChatPhotoParameters struct {
	ChatParameters
}
func (p *SetChatPhotoParameters) Photo(filename string, filedata io.Reader) {
	p.withFile("photo", filename, filedata)
}


type SetChatTitleParameters struct {
	ChatParameters
}
func (p *SetChatTitleParameters) Title(title string) {
	p.withString("title", title)
}


type SetChatDescriptionParameters struct {
	ChatParameters
}
func (p *SetChatDescriptionParameters) Description(description string) {
	p.withString("description", description)
}


type PinChatMessageParameters struct {
	ChatParameters
}
func (p *PinChatMessageParameters) MessageId(messageId int) {
	p.withInt("message_id", messageId)
}
func (p *PinChatMessageParameters) DisableNotification() {
	p.withBool("disable_notification", true)
}


type SetChatStickerSetParameters struct {
	ChatParameters
}
func (p *SetChatStickerSetParameters) StickerSet(name string) {
	p.withString("sticker_set_name", name)
}
