package parameters

import (
	"github.com/deadblue/telegroid/types"
)

// Parameters for deleteMessage.
// Reference: https://core.telegram.org/bots/api#deletemessage
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

// Parameters for stopPoll.
// Reference: https://core.telegram.org/bots/api#stoppoll
type StopPollParams struct {
	DeleteMessageParams
}

func (p *StopPollParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) {
	p.setJson("reply_markup", markup)
}

// Parameters for editMessageReplyMarkup.
// Reference: https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkupParams struct {
	StopPollParams
}

func (p *EditMessageReplyMarkupParams) InlineMessage(messageId string) {
	p.set("inline_message_id", messageId)
}

// Parameters for editMessageText.
// Reference: https://core.telegram.org/bots/api#editmessagetext
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

// Parameters for editMessageCaption.
// Reference: https://core.telegram.org/bots/api#editmessagecaption
type EditMessageCaptionParams struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageCaptionParams) Caption(text FormattedText) {
	p.set("caption", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
}

// Parameters for editMessageMedia.
// Reference: https://core.telegram.org/bots/api#editmessagemedia
type EditMessageMediaParams struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageMediaParams) Photo(photo *InputPhoto) {
	if photo == nil || photo.media == nil {
		return
	}
	info := &types.InputMediaPhoto{
		Type:      types.MediaPhoto,
		Caption:   photo.caption.text,
		ParseMode: photo.caption.mode,
	}
	if photo.media.fileIdOrUrl != "" {
		info.Media = photo.media.fileIdOrUrl
	} else {
		info.Media = randomAttachName()
		p.setFile(info.Media, photo.media)
	}
	p.setJson("media", info)
}

func (p *EditMessageMediaParams) Video(video *InputVideo) {
	if video == nil || video.media == nil {
		return
	}
	info := &types.InputMediaVideo{
		Type:              types.MediaVideo,
		Caption:           video.caption.text,
		ParseMode:         video.caption.mode,
		Width:             video.width,
		Height:            video.height,
		Duration:          video.duration,
		SupportsStreaming: video.supportsStreaming,
	}
	if video.media.fileIdOrUrl != "" {
		info.Media = video.media.fileIdOrUrl
	} else {
		info.Media = randomAttachName()
		p.setFile(info.Media, video.media)
	}
	if video.thumb != nil {
		if video.thumb.fileIdOrUrl != "" {
			info.Thumb = video.thumb.fileIdOrUrl
		} else {
			info.Thumb = randomAttachName()
			p.setFile(info.Thumb, video.thumb)
		}
	}
	p.setJson("media", info)
}

func (p *EditMessageMediaParams) Animation(animation *InputAnimation) {
	if animation == nil || animation.media == nil {
		return
	}
	info := &types.InputMediaAnimation{
		Type:      types.MediaAnimation,
		Caption:   animation.caption.text,
		ParseMode: animation.caption.mode,
		Width:     animation.width,
		Height:    animation.height,
		Duration:  animation.duration,
	}
	if animation.media.fileIdOrUrl != "" {
		info.Media = animation.media.fileIdOrUrl
	} else {
		info.Media = randomAttachName()
		p.setFile(info.Media, animation.media)
	}
	if animation.thumb != nil {
		if animation.thumb.fileIdOrUrl != "" {
			info.Thumb = animation.thumb.fileIdOrUrl
		} else {
			info.Thumb = randomAttachName()
			p.setFile(info.Thumb, animation.thumb)
		}
	}
	p.setJson("media", info)
}

func (p *EditMessageMediaParams) Audio(audio *InputAudio) {
	if audio == nil || audio.media == nil {
		return
	}
	info := &types.InputMediaAudio{
		Type:      types.MediaAudio,
		Caption:   audio.caption.text,
		ParseMode: audio.caption.mode,
		Duration:  audio.duration,
		Performer: audio.performer,
		Title:     audio.title,
	}
	if audio.media.fileIdOrUrl != "" {
		info.Media = audio.media.fileIdOrUrl
	} else {
		info.Media = randomAttachName()
		p.setFile(info.Media, audio.media)
	}
	if audio.thumb != nil {
		if audio.thumb.fileIdOrUrl != "" {
			info.Thumb = audio.thumb.fileIdOrUrl
		} else {
			info.Thumb = randomAttachName()
			p.setFile(info.Thumb, audio.thumb)
		}
	}
	p.setJson("media", info)
}

func (p *EditMessageMediaParams) Document(document *InputDocument) {
	if document == nil || document.media == nil {
		return
	}
	info := &types.InputMediaDocument{
		Type:      types.MediaDocument,
		Caption:   document.caption.text,
		ParseMode: document.caption.mode,
	}
	if document.media.fileIdOrUrl != "" {
		info.Media = document.media.fileIdOrUrl
	} else {
		info.Media = randomAttachName()
		p.setFile(info.Media, document.media)
	}
	if document.thumb != nil {
		if document.thumb.fileIdOrUrl != "" {
			info.Thumb = document.thumb.fileIdOrUrl
		} else {
			info.Thumb = randomAttachName()
			p.setFile(info.Thumb, document.thumb)
		}
	}
	p.setJson("media", info)
}

// Parameters for editMessageLiveLocation.
// Reference: https://core.telegram.org/bots/api#editmessagelivelocation
type EditMessageLiveLocation struct {
	EditMessageReplyMarkupParams
}

func (p *EditMessageLiveLocation) Location(latitude, longitude float64) {
	p.setFloat("latitude", latitude)
	p.setFloat("longitude", longitude)
}

// Parameters for stopMessageLiveLocation, which is as same as EditMessageReplyMarkupParams.
// Reference: https://core.telegram.org/bots/api#stopmessagelivelocation
type StopMessageLiveLocation struct {
	EditMessageReplyMarkupParams
}
