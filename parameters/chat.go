package parameters

type ChatParams struct {
	implApiParameters
}

func (p *ChatParams) ChatId(chatId int64) *ChatParams {
	p.setInt64("chat_id", chatId)
	return p
}
func (p *ChatParams) ChannelName(name string) *ChatParams {
	p.set("chat_id", "@"+name)
	return p
}

type GetChatMemberParams struct {
	implApiParameters
}

func (p *GetChatMemberParams) ChatId(chatId int64) *GetChatMemberParams {
	p.setInt64("chat_id", chatId)
	return p
}
func (p *GetChatMemberParams) ChannelName(name string) *GetChatMemberParams {
	p.set("chat_id", "@"+name)
	return p
}
func (p *GetChatMemberParams) UserId(userId int) *GetChatMemberParams {
	p.setInt("user_id", userId)
	return p
}

type SetChatStickerSetParams struct {
	implApiParameters
}

func (p *SetChatStickerSetParams) ChatId(chatId int64) *SetChatStickerSetParams {
	p.setInt64("chat_id", chatId)
	return p
}
func (p *SetChatStickerSetParams) ChannelName(name string) *SetChatStickerSetParams {
	p.set("chat_id", "@"+name)
	return p
}
func (p *SetChatStickerSetParams) StickerSetName(name string) *SetChatStickerSetParams {
	p.set("sticker_set_name", name)
	return p
}
