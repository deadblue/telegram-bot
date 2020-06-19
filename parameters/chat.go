package parameters

import "github.com/deadblue/telegroid/internal/protocol"

type ChatParams struct {
	protocol.BaseApiRequest
}

func (p *ChatParams) ChatId(chatId int64) *ChatParams {
	p.SetInt64("chat_id", chatId)
	return p
}
func (p *ChatParams) ChannelName(name string) *ChatParams {
	p.Set("chat_id", "@"+name)
	return p
}

type GetChatMemberParams struct {
	protocol.BaseApiRequest
}

func (p *GetChatMemberParams) ChatId(chatId int64) *GetChatMemberParams {
	p.SetInt64("chat_id", chatId)
	return p
}
func (p *GetChatMemberParams) ChannelName(name string) *GetChatMemberParams {
	p.Set("chat_id", "@"+name)
	return p
}
func (p *GetChatMemberParams) UserId(userId int) *GetChatMemberParams {
	p.SetInt("user_id", userId)
	return p
}

type SetChatStickerSetParams struct {
	protocol.BaseApiRequest
}

func (p *SetChatStickerSetParams) ChatId(chatId int64) *SetChatStickerSetParams {
	p.SetInt64("chat_id", chatId)
	return p
}
func (p *SetChatStickerSetParams) ChannelName(name string) *SetChatStickerSetParams {
	p.Set("chat_id", "@"+name)
	return p
}
func (p *SetChatStickerSetParams) StickerSetName(name string) *SetChatStickerSetParams {
	p.Set("sticker_set_name", name)
	return p
}
