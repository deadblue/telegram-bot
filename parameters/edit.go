package parameters

import (
	"github.com/deadblue/telegroid/types"
)

type EditMessageTextParams struct {
	implApiParameters
}

func (p *EditMessageTextParams) ChatMessage(chatId int64, messageId int) *EditMessageTextParams {
	p.setInt64("chat_id", chatId)
	p.setInt("message_id", messageId)
	return p
}
func (p *EditMessageTextParams) ChannelPost(name string, messageId int) *EditMessageTextParams {
	p.set("chat_id", "@"+name)
	p.setInt("message_id", messageId)
	return p
}
func (p *EditMessageTextParams) InlineMessage(messageId int) *EditMessageTextParams {
	p.setInt("inline_message_id", messageId)
	return p
}
func (p *EditMessageTextParams) Text(text FormattedText) *EditMessageTextParams {
	p.set("text", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
	return p
}
func (p *EditMessageTextParams) DisableWebPagePreview() *EditMessageTextParams {
	p.setBool("disable_web_page_preview", true)
	return p
}
func (p *EditMessageTextParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *EditMessageTextParams {
	p.setJson("reply_markup", markup)
	return p
}

type EditMessageCaptionParams struct {
	implApiParameters
}

func (p *EditMessageCaptionParams) ChatMessage(chatId int64, messageId int) *EditMessageCaptionParams {
	p.setInt64("chat_id", chatId)
	p.setInt("message_id", messageId)
	return p
}
func (p *EditMessageCaptionParams) ChannelPost(name string, messageId int) *EditMessageCaptionParams {
	p.set("chat_id", "@"+name)
	p.setInt("message_id", messageId)
	return p
}
func (p *EditMessageCaptionParams) InlineMessage(messageId int) *EditMessageCaptionParams {
	p.setInt("inline_message_id", messageId)
	return p
}
func (p *EditMessageCaptionParams) Caption(text FormattedText) *EditMessageCaptionParams {
	p.set("caption", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
	return p
}
func (p *EditMessageCaptionParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *EditMessageCaptionParams {
	p.setJson("reply_markup", markup)
	return p
}

type EditMessageReplyMarkupParams struct {
	implApiParameters
}

func (p *EditMessageReplyMarkupParams) ChatMessage(chatId int64, messageId int) *EditMessageReplyMarkupParams {
	p.setInt64("chat_id", chatId)
	p.setInt("message_id", messageId)
	return p
}
func (p *EditMessageReplyMarkupParams) ChannelPost(name string, messageId int) *EditMessageReplyMarkupParams {
	p.set("chat_id", "@"+name)
	p.setInt("message_id", messageId)
	return p
}
func (p *EditMessageReplyMarkupParams) InlineMessage(messageId int) *EditMessageReplyMarkupParams {
	p.setInt("inline_message_id", messageId)
	return p
}
func (p *EditMessageReplyMarkupParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *EditMessageReplyMarkupParams {
	p.setJson("reply_markup", markup)
	return p
}

type StopPollParams struct {
	implApiParameters
}

func (p *StopPollParams) ChatMessage(chatId int64, messageId int) *StopPollParams {
	p.setInt64("chat_id", chatId)
	p.setInt("message_id", messageId)
	return p
}
func (p *StopPollParams) ChannelPost(name string, messageId int) *StopPollParams {
	p.set("chat_id", "@"+name)
	p.setInt("message_id", messageId)
	return p
}
func (p *StopPollParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) *StopPollParams {
	p.setJson("reply_markup", markup)
	return p
}
