package parameters

import (
	"github.com/deadblue/telegram-bot/types"
	"time"
)

// Base parameters for most send methods.
type baseSendParams struct {
	ChatParams
}

func (p *baseSendParams) DisableNotification() {
	p.setBool("disable_notification", true)
}
func (p *baseSendParams) PeplyToMessage(messageId int) {
	p.setInt("reply_to_message_id", messageId)
}
func (p *baseSendParams) ForceReply(selective bool) {
	p.setJson("reply_markup", &types.ForceReply{
		ForceReply: true,
		Selective:  selective,
	})
}
func (p *baseSendParams) RemoveKeyboard(selective bool) {
	p.setJson("reply_markup", &types.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	})
}
func (p *baseSendParams) ReplyKeyboard(markup *types.ReplyKeyboardMarkup) {
	p.setJson("reply_markup", markup)
}
func (p *baseSendParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) {
	p.setJson("reply_markup", markup)
}

// Parameters for sendMessage method.
// Reference: https://core.telegram.org/bots/api#sendmessage
type SendMessageParams struct {
	baseSendParams
}

func (p *SendMessageParams) Text(text FormattedText) {
	p.set("text", text.String())
	if mode := text.Mode(); mode != types.ModePlain {
		p.set("parse_mode", string(mode))
	}
}
func (p *SendMessageParams) DisableWebPagePreview() {
	p.setBool("disable_web_page_preview", true)
}

type baseSendMediaParams struct {
	baseSendParams
}

func (p *baseSendMediaParams) Caption(text FormattedText) {
	p.set("caption", text.String())
	if mode := text.Mode(); mode != types.ModePlain {
		p.set("parse_mode", string(mode))
	}
}

// Parameters for sendPhoto method.
// Reference: https://core.telegram.org/bots/api#sendphoto
type SendPhotoParams struct {
	baseSendMediaParams
}

func (p *SendPhotoParams) Photo(file *InputFile) {
	p.setFile("photo", file)
}

// Parameters for sendAudio method.
// Reference: https://core.telegram.org/bots/api#sendaudio
type SendAudioParams struct {
	baseSendMediaParams
}

func (p *SendAudioParams) Audio(file *InputFile) {
	p.setFile("audio", file)
}
func (p *SendAudioParams) Duration(duration int) {
	p.setInt("duration", duration)
}
func (p *SendAudioParams) Performer(performer string) {
	p.set("performer", performer)
}
func (p *SendAudioParams) Title(title string) {
	p.set("title", title)
}
func (p *SendAudioParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}

// Parameters for sendDocument method.
// Reference: https://core.telegram.org/bots/api#senddocument
type SendDocumentParams struct {
	baseSendMediaParams
}

func (p *SendDocumentParams) Document(file *InputFile) {
	p.setFile("document", file)
}
func (p *SendDocumentParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}

// Parameters for sendVideo method.
// Reference: https://core.telegram.org/bots/api#sendvideo
type SendVideoParams struct {
	baseSendMediaParams
}

func (p *SendVideoParams) Video(file *InputFile) {
	p.setFile("video", file)
}
func (p *SendVideoParams) Duration(duration int) {
	p.setInt("duration", duration)
}
func (p *SendVideoParams) Dimension(width, height int) {
	p.setInt("width", width)
	p.setInt("height", height)
}
func (p *SendVideoParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}
func (p *SendVideoParams) SupportsStreaming() {
	p.setBool("supports_streaming", true)
}

// Parameters for sendAnimation method.
// Reference: https://core.telegram.org/bots/api#sendanimation
type SendAnimationParams struct {
	baseSendMediaParams
}

func (p *SendAnimationParams) Animation(file *InputFile) {
	p.setFile("animation", file)
}
func (p *SendAnimationParams) Duration(duration int) {
	p.setInt("duration", duration)
}
func (p *SendAnimationParams) Size(width, height int) {
	p.setInt("width", width)
	p.setInt("height", height)
}
func (p *SendAnimationParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}

// Parameters for sendVoice method.
// Reference: https://core.telegram.org/bots/api#sendvoice
type SendVoiceParams struct {
	baseSendMediaParams
}

func (p *SendVoiceParams) Voice(file *InputFile) {
	p.setFile("voice", file)
}
func (p *SendVoiceParams) Duration(duration int) {
	p.setInt("duration", duration)
}

// Parameters for sendVideoNote method.
// Reference: https://core.telegram.org/bots/api#sendvideonote
type SendVideoNoteParams struct {
	baseSendMediaParams
}

func (p *SendVideoNoteParams) VideoNote(file *InputFile) {
	p.setFile("video_note", file)
}
func (p *SendVideoNoteParams) Duration(duration int) {
	p.setInt("duration", duration)
}
func (p *SendVideoNoteParams) SideLength(length int) {
	p.setInt("length", length)
}
func (p *SendVideoNoteParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}

// Parameters for sendMediaGroup method method.
// Reference: https://core.telegram.org/bots/api#sendmediagroup
type SendMediaGroupParams struct {
	ChatParams
}

func (p *SendMediaGroupParams) DisableNotification() {
	p.setBool("disable_notification", true)
}
func (p *SendMediaGroupParams) ReplyToMessage(messageId int) {
	p.setInt("reply_to_message_id", messageId)
}
func (p *SendMediaGroupParams) Media(media ...InputMedia) {
	infos, files := make([]interface{}, 0), make(map[string]*InputFile)
	for _, item := range media {
		switch item.(type) {
		case *InputPhoto:
			if photo := item.(*InputPhoto); photo.media != nil {
				info := &types.InputMediaPhoto{
					Type: types.MediaPhoto,
				}
				if caption := photo.caption; caption != nil {
					info.Caption = caption.String()
					info.ParseMode = caption.Mode()
				}
				if photo.media.fileIdOrUrl != "" {
					info.Media = photo.media.fileIdOrUrl
				} else {
					name, uri := randomAttachNameAndUri()
					files[name] = photo.media
					info.Media = uri
				}
				infos = append(infos, info)
			}
		case *InputVideo:
			if video := item.(*InputVideo); video.media != nil {
				info := &types.InputMediaVideo{
					Type:              types.MediaVideo,
					Width:             video.width,
					Height:            video.height,
					Duration:          video.duration,
					SupportsStreaming: video.supportsStreaming,
				}
				if caption := video.caption; caption != nil {
					info.Caption = caption.String()
					info.ParseMode = caption.Mode()
				}
				if video.media.fileIdOrUrl != "" {
					info.Media = video.media.fileIdOrUrl
				} else {
					name, uri := randomAttachNameAndUri()
					files[name] = video.media
					info.Media = uri
				}
				if video.thumb != nil {
					if video.thumb.fileIdOrUrl != "" {
						info.Thumb = video.thumb.fileIdOrUrl
					} else {
						name, uri := randomAttachNameAndUri()
						files[name] = video.thumb
						info.Thumb = uri
					}
				}
				infos = append(infos, info)
			}
		}
	}
	if count := len(infos); count >= 2 && count <= 10 {
		p.setJson("media", infos)
		for name, file := range files {
			p.setFile(name, file)
		}
	}
}

// Parameters for sendLocation method.
// Reference: https://core.telegram.org/bots/api#sendlocation
type SendLocationParams struct {
	baseSendParams
}

func (p *SendLocationParams) Location(latitude, longitude float64) {
	p.setFloat("latitude", latitude)
	p.setFloat("longitude", longitude)
}
func (p *SendLocationParams) LivePeriod(period int) {
	if period >= 60 && period <= 86400 {
		p.setInt("live_period", period)
	}
}

// Parameters for sendVenue method.
// Reference: https://core.telegram.org/bots/api#sendvenue
type SendVenueParams struct {
	baseSendParams
}

func (p *SendVenueParams) Location(latitude, longitude float64) {
	p.setFloat("latitude", latitude)
	p.setFloat("longitude", longitude)
}
func (p *SendVenueParams) Title(title string) {
	p.set("title", title)
}
func (p *SendVenueParams) Address(address string) {
	p.set("address", address)
}
func (p *SendVenueParams) Foursquare(fsId, fsType string) {
	p.set("foursquare_id", fsId)
	p.set("foursquare_type", fsType)
}

// Parameters for sendContact method.
// Reference: https://core.telegram.org/bots/api#sendcontact
type SendContactParams struct {
	baseSendParams
}

func (p *SendContactParams) PhoneNumber(phone string) {
	p.set("phone_number", phone)
}
func (p *SendContactParams) Name(firstName, lastName string) {
	p.set("first_name", firstName)
	if lastName != "" {
		p.set("last_name", lastName)
	}
}
func (p *SendContactParams) Vcard(vcard string) {
	p.set("vcard", vcard)
}

// Parameters for sendPoll method.
// Reference: https://core.telegram.org/bots/api#sendpoll
type SendPollParams struct {
	baseSendParams
}

func (p *SendPollParams) Topic(question string, options ...string) {
	p.set("question", question)
	p.setJson("options", options)
}
func (p *SendPollParams) IsAnonymous(value bool) {
	p.setBool("is_anonymous", value)
}
func (p *SendPollParams) AllowsMultipleAnswers() {
	p.setBool("allows_multiple_answers", true)
}
func (p *SendPollParams) QuizMode(correctOptionId int, explanation FormattedText) {
	p.set("type", "quiz")
	p.setInt("correct_option_id", correctOptionId)
	if explanation != nil {
		p.set("explanation", explanation.String())
		if mode := explanation.Mode(); mode != types.ModePlain {
			p.set("explanation_parse_mode", string(mode))
		}
	}

}
func (p *SendPollParams) OpenPeriod(period int) {
	if period >= 5 && period <= 600 {
		p.setInt("open_period", period)
	}
}
func (p *SendPollParams) CloseDate(date time.Time) {
	p.setInt64("open_period", date.Unix())
}
func (p *SendPollParams) Closed() {
	p.setBool("is_closed", true)
}

// Parameters for sendDice method.
// Reference: https://core.telegram.org/bots/api#senddice
type SendDiceParams struct {
	baseSendParams
}

func (p *SendDiceParams) Dice() {
	p.set("emoji", "🎲")
}
func (p *SendDiceParams) Darts() {
	p.set("emoji", "🎯")
}
func (p *SendDiceParams) Basketball() {
	p.set("emoji", "🏀")
}

// Parameters for forwardMessage method.
// Reference: https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageParams struct {
	ChatParams
}

func (p *ForwardMessageParams) FromChatMessage(chatId int64, messageId int) {
	p.setInt64("from_chat_id", chatId)
	p.setInt("message_id", messageId)
}
func (p *ForwardMessageParams) FromChannelPost(name string, messageId int) {
	p.set("from_chat_id", "@"+name)
	p.setInt("message_id", messageId)
}
func (p *ForwardMessageParams) DisableNotification() *ForwardMessageParams {
	p.setBool("disable_notification", true)
	return p
}
