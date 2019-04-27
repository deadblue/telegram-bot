package arguments

import (
	"fmt"
)

type _ChatArgs struct {
	_BasicArgs
}

func (a *_ChatArgs) ChatId(chatId int) {
	a.getForm().WithInt("chat_id", chatId)
}
func (a *_ChatArgs) Channel(channelName string) {
	a.getForm().WithString("chat_id", fmt.Sprintf("@%s", channelName))
}

type ChatArgs _ChatArgs

type _ChatMemberArgs struct {
	_ChatArgs
}

func (a *_ChatMemberArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}

type ChatMemberArgs _ChatMemberArgs

type _KickChatMemberArgs struct {
	_ChatMemberArgs
}

func (a *_KickChatMemberArgs) UntilDate(timestamp int64) {
	a.getForm().WithInt64("until_date", timestamp)
}

type KickChatMemberArgs _KickChatMemberArgs

type RestrictChatMemberArgs struct {
	_KickChatMemberArgs
}

func (a *RestrictChatMemberArgs) SendMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_send_messages",
	}
}
func (a *RestrictChatMemberArgs) SendMediaMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_send_media_messages",
	}
}
func (a *RestrictChatMemberArgs) SendOtherMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_send_other_messages",
	}
}
func (a *RestrictChatMemberArgs) AddWebPagePreviews() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_add_web_page_previews",
	}
}

type PromoteChatMemberArgs struct {
	_ChatMemberArgs
}

func (a *PromoteChatMemberArgs) ChangeInfo() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_change_info",
	}
}
func (a *PromoteChatMemberArgs) PostMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_post_messages",
	}
}
func (a *PromoteChatMemberArgs) EditMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_edit_messages",
	}
}
func (a *PromoteChatMemberArgs) DeleteMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_delete_messages",
	}
}
func (a *PromoteChatMemberArgs) InviteUsers() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_invite_users",
	}
}
func (a *PromoteChatMemberArgs) RestrictMembers() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_restrict_members",
	}
}
func (a *PromoteChatMemberArgs) PinMessages() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_pin_messages",
	}
}
func (a *PromoteChatMemberArgs) PromoteMembers() Switch {
	return &implSwitch{
		form: a.getForm(),
		name: "can_promote_members",
	}
}

type SetChatPhotoArgs struct {
	_ChatArgs
}

func (a *SetChatPhotoArgs) Photo(file InputFile) {
	a.getForm().WithFile("photo", file)
}

type SetChatTitleArgs struct {
	_ChatArgs
}

func (a *SetChatTitleArgs) Title(title string) {
	a.getForm().WithString("title", title)
}

type SetChatDescriptionArgs struct {
	_ChatArgs
}

func (a *SetChatDescriptionArgs) Description(description string) {
	a.getForm().WithString("description", description)
}

type PinChatMessageArgs struct {
	_ChatArgs
}

func (a *PinChatMessageArgs) MessageId(messageId int) {
	a.getForm().WithInt("message_id", messageId)
}
func (a *PinChatMessageArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}

type SetChatStickerSetArgs struct {
	_ChatArgs
}

func (a *SetChatStickerSetArgs) StickerSet(name string) {
	a.getForm().WithString("sticker_set_name", name)
}
