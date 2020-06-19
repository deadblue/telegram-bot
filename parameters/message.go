package parameters

import (
	"github.com/deadblue/telegroid/internal/protocol"
	"github.com/deadblue/telegroid/types"
)

type SendMessageParams struct {
	protocol.BaseApiRequest
}

func (p *SendMessageParams) ChatId(chatId int64) *SendMessageParams {
	p.SetInt64("chat_id", chatId)
	return p
}

func (p *SendMessageParams) ChannelName(name string) *SendMessageParams {
	p.Set("chat_id", "@"+name)
	return p
}
func (p *SendMessageParams) Text(text FormattedText) *SendMessageParams {
	p.Set("text", text.text)
	if text.mode != types.ModePlain {
		p.Set("parse_mode", string(text.mode))
	}
	return p
}
func (p *SendMessageParams) ForceReply(seletive bool) *SendMessageParams {
	p.SetJson("reply_markup", &types.ForceReply{
		ForceReply: true,
		Selective:  seletive,
	})
	return p
}
func (p *SendMessageParams) RemoveKeyboard(seletive bool) *SendMessageParams {
	p.SetJson("reply_markup", &types.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      seletive,
	})
	return p
}
func (p *SendMessageParams) DisableWebPagePreview() *SendMessageParams {
	p.SetBool("disable_web_page_preview", true)
	return p
}
func (p *SendMessageParams) DisableNotification() *SendMessageParams {
	p.SetBool("disable_notification", true)
	return p
}

type ForwardMessageParams struct {
	protocol.BaseApiRequest
}

func (p *ForwardMessageParams) Chat(chatId int64) *ForwardMessageParams {
	p.SetInt64("chat_id", chatId)
	return p
}
func (p *ForwardMessageParams) Channel(name string) *ForwardMessageParams {
	p.Set("chat_id", "@"+name)
	return p
}
func (p *ForwardMessageParams) FromChat(chatId int64) *ForwardMessageParams {
	p.SetInt64("from_chat_id", chatId)
	return p
}
func (p *ForwardMessageParams) FromChannel(name string) *ForwardMessageParams {
	p.Set("from_chat_id", "@"+name)
	return p
}
func (p *ForwardMessageParams) MessageId(messageId int) *ForwardMessageParams {
	p.SetInt("message_id", messageId)
	return p
}
func (p *ForwardMessageParams) DisableNotification() *ForwardMessageParams {
	p.SetBool("disable_notification", true)
	return p
}
