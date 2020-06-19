package parameters

import (
	"github.com/deadblue/telegroid/internal/protocol"
	"github.com/deadblue/telegroid/types"
)

type EditMessageTextParams struct {
	protocol.BaseApiRequest
}

func (p *EditMessageTextParams) ChatMessage(chatId int64, messageId int) *EditMessageTextParams {
	p.SetInt64("chat_id", chatId)
	p.SetInt("message_id", messageId)
	return p
}
func (p *EditMessageTextParams) ChannelPost(name string, messageId int) *EditMessageTextParams {
	p.Set("chat_id", "@"+name)
	p.SetInt("message_id", messageId)
	return p
}
func (p *EditMessageTextParams) InlineMessage(messageId int) *EditMessageTextParams {
	p.SetInt("inline_message_id", messageId)
	return p
}
func (p *EditMessageTextParams) Text(text FormattedText) *EditMessageTextParams {
	p.Set("text", text.text)
	if text.mode != types.ModePlain {
		p.Set("parse_mode", string(text.mode))
	}
	return p
}
func (p *EditMessageTextParams) DisableWebPagePreview() *EditMessageTextParams {
	p.SetBool("disable_web_page_preview", true)
	return p
}
func (p *EditMessageTextParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *EditMessageTextParams {
	p.SetJson("reply_markup", markup)
	return p
}

type EditMessageCaptionParams struct {
	protocol.BaseApiRequest
}

func (p *EditMessageCaptionParams) ChatMessage(chatId int64, messageId int) *EditMessageCaptionParams {
	p.SetInt64("chat_id", chatId)
	p.SetInt("message_id", messageId)
	return p
}
func (p *EditMessageCaptionParams) ChannelPost(name string, messageId int) *EditMessageCaptionParams {
	p.Set("chat_id", "@"+name)
	p.SetInt("message_id", messageId)
	return p
}
func (p *EditMessageCaptionParams) InlineMessage(messageId int) *EditMessageCaptionParams {
	p.SetInt("inline_message_id", messageId)
	return p
}
func (p *EditMessageCaptionParams) Caption(text FormattedText) *EditMessageCaptionParams {
	p.Set("caption", text.text)
	if text.mode != types.ModePlain {
		p.Set("parse_mode", string(text.mode))
	}
	return p
}
func (p *EditMessageCaptionParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *EditMessageCaptionParams {
	p.SetJson("reply_markup", markup)
	return p
}

type EditMessageReplyMarkupParams struct {
	protocol.BaseApiRequest
}

func (p *EditMessageReplyMarkupParams) ChatMessage(chatId int64, messageId int) *EditMessageReplyMarkupParams {
	p.SetInt64("chat_id", chatId)
	p.SetInt("message_id", messageId)
	return p
}
func (p *EditMessageReplyMarkupParams) ChannelPost(name string, messageId int) *EditMessageReplyMarkupParams {
	p.Set("chat_id", "@"+name)
	p.SetInt("message_id", messageId)
	return p
}
func (p *EditMessageReplyMarkupParams) InlineMessage(messageId int) *EditMessageReplyMarkupParams {
	p.SetInt("inline_message_id", messageId)
	return p
}
func (p *EditMessageReplyMarkupParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *EditMessageReplyMarkupParams {
	p.SetJson("reply_markup", markup)
	return p
}

type StopPollParams struct {
	protocol.BaseApiRequest
}

func (p *StopPollParams) ChatMessage(chatId int64, messageId int) *StopPollParams {
	p.SetInt64("chat_id", chatId)
	p.SetInt("message_id", messageId)
	return p
}
func (p *StopPollParams) ChannelPost(name string, messageId int) *StopPollParams {
	p.Set("chat_id", "@"+name)
	p.SetInt("message_id", messageId)
	return p
}
func (p *StopPollParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *StopPollParams {
	p.SetJson("reply_markup", markup)
	return p
}
