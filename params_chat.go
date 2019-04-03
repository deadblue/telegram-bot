package telegroid

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


type SwitchParameter struct {
	holder _BasicParameters
	name   string
}
func (p *SwitchParameter) On() {
	p.holder.withBool(p.name, true)
}
func (p *SwitchParameter) Off() {
	p.holder.withBool(p.name, false)
}


type RestrictChatMemberParameters struct {
	KickChatMemberParameters
}
func (p *RestrictChatMemberParameters) SendMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_send_messages",
	}
}
func (p *RestrictChatMemberParameters) SendMediaMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_send_media_messages",
	}
}
func (p *RestrictChatMemberParameters) SendOtherMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_send_other_messages",
	}
}
func (p *RestrictChatMemberParameters) AddWebPagePreviews() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_add_web_page_previews",
	}
}


type PromoteChatMemberParameters struct {
	ChatMemberParameters
}
func (p *PromoteChatMemberParameters) ChangeInfo() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_change_info",
	}
}
func (p *PromoteChatMemberParameters) PostMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_post_messages",
	}
}
func (p *PromoteChatMemberParameters) EditMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_edit_messages",
	}
}
func (p *PromoteChatMemberParameters) DeleteMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_delete_messages",
	}
}
func (p *PromoteChatMemberParameters) InviteUsers() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_invite_users",
	}
}
func (p *PromoteChatMemberParameters) RestrictMembers() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_restrict_members",
	}
}
func (p *PromoteChatMemberParameters) PinMessages() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "can_pin_messages",
	}
}
func (p *PromoteChatMemberParameters) PromoteMembers() *SwitchParameter {
	return &SwitchParameter{
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
