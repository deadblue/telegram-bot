package parameters

import (
	"github.com/deadblue/telegroid/types"
)

type DeleteMessageParams struct {
	implApiParameters
}

func (p *DeleteMessageParams) ChatMessage(chatId int64, messageId int) {
	p.setInt64("chat_id", chatId)
	p.setInt("message_id", messageId)
}
func (p *DeleteMessageParams) ChannelPost(channelName string, messageId int) {
	p.set("chat_id", "@"+channelName)
	p.setInt("message_id", messageId)
}

type StopPollParams struct {
	DeleteMessageParams
}

func (p *StopPollParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) {
	p.setJson("reply_markup", markup)
}

type EditMessageReplyMarkupParams struct {
	DeleteMessageParams
}

func (p *EditMessageReplyMarkupParams) InlineMessage(inlineMessageId string) {
	p.set("inline_message_id", inlineMessageId)
}

type EditMessageTextParams struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageTextParams) Text(text FormattedText) {
	p.set("text", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
}
func (p *EditMessageTextParams) DisableWebPagePreview() {
	p.setBool("disable_web_page_preview", true)
}

type EditMessageCaptionParams struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageCaptionParams) Caption(text FormattedText) {
	p.set("caption", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
}

type EditMessageMediaParams struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageMediaParams) Animation(media *types.InputMediaAnimation) {
	if media.Type == "" {
		media.Type = types.MediaAnimation
	}
	p.setJson("media", media)
}
func (p *EditMessageMediaParams) Audio(media *types.InputMediaAudio) {
	if media.Type == "" {
		media.Type = types.MediaAudio
	}
	p.setJson("media", media)
}
func (p *EditMessageMediaParams) Document(media *types.InputMediaDocument) {
	if media.Type == "" {
		media.Type = types.MediaDocument
	}
	p.setJson("media", media)
}
func (p *EditMessageMediaParams) Photo(media *types.InputMediaPhoto) {
	if media.Type == "" {
		media.Type = types.MediaPhoto
	}
	p.setJson("media", media)
}
func (p *EditMessageMediaParams) Video(media *types.InputMediaVideo) {
	if media.Type == "" {
		media.Type = types.MediaVideo
	}
	p.setJson("media", media)
}

type EditMessageLiveLocation struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageLiveLocation) Location(latitude, longitude float64) {
	p.setFloat("latitude", latitude)
	p.setFloat("longitude", longitude)
}

type StopMessageLiveLocation struct {
	EditMessageReplyMarkupParams
}
