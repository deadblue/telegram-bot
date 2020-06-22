package parameters

import (
	"github.com/deadblue/telegroid/types"
)

type SendMessageParams struct {
	implApiParameters
}

func (p *SendMessageParams) ChatId(chatId int64) *SendMessageParams {
	p.setInt64("chat_id", chatId)
	return p
}

func (p *SendMessageParams) ChannelName(name string) *SendMessageParams {
	p.set("chat_id", "@"+name)
	return p
}
func (p *SendMessageParams) Text(text FormattedText) *SendMessageParams {
	p.set("text", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
	return p
}
func (p *SendMessageParams) ForceReply(seletive bool) *SendMessageParams {
	p.setJson("reply_markup", &types.ForceReply{
		ForceReply: true,
		Selective:  seletive,
	})
	return p
}
func (p *SendMessageParams) RemoveKeyboard(seletive bool) *SendMessageParams {
	p.setJson("reply_markup", &types.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      seletive,
	})
	return p
}
func (p *SendMessageParams) DisableWebPagePreview() *SendMessageParams {
	p.setBool("disable_web_page_preview", true)
	return p
}
func (p *SendMessageParams) DisableNotification() *SendMessageParams {
	p.setBool("disable_notification", true)
	return p
}

type ForwardMessageParams struct {
	implApiParameters
}

func (p *ForwardMessageParams) Chat(chatId int64) *ForwardMessageParams {
	p.setInt64("chat_id", chatId)
	return p
}
func (p *ForwardMessageParams) Channel(name string) *ForwardMessageParams {
	p.set("chat_id", "@"+name)
	return p
}
func (p *ForwardMessageParams) FromChat(chatId int64) *ForwardMessageParams {
	p.setInt64("from_chat_id", chatId)
	return p
}
func (p *ForwardMessageParams) FromChannel(name string) *ForwardMessageParams {
	p.set("from_chat_id", "@"+name)
	return p
}
func (p *ForwardMessageParams) MessageId(messageId int) *ForwardMessageParams {
	p.setInt("message_id", messageId)
	return p
}
func (p *ForwardMessageParams) DisableNotification() *ForwardMessageParams {
	p.setBool("disable_notification", true)
	return p
}
