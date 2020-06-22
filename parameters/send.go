package parameters

import (
	"github.com/deadblue/telegroid/types"
	"time"
)

// Base struct for all send operation parameters.
type baseSendParams struct {
	ChatParams
}

func (p *baseSendParams) DisableNotification() {
	p.setBool("disable_notification", true)
}
func (p *baseSendParams) PeplyToMessage(messageId int) {
	p.setInt("reply_to_message_id", messageId)
}
func (p *baseSendParams) ForceReply(seletive bool) {
	p.setJson("reply_markup", &types.ForceReply{
		ForceReply: true,
		Selective:  seletive,
	})
}
func (p *baseSendParams) RemoveKeyboard(seletive bool) {
	p.setJson("reply_markup", &types.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      seletive,
	})
}

// Parameters holder for sending text message.
type SendMessageParams struct {
	baseSendParams
}

func (p *SendMessageParams) Text(text FormattedText) {
	p.set("text", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
}
func (p *SendMessageParams) DisableWebPagePreview() {
	p.setBool("disable_web_page_preview", true)
}

// Base parameters holder for sending media message.
type baseSendMediaParams struct {
	baseSendParams
}

func (p *baseSendMediaParams) Caption(text FormattedText) {
	p.set("caption", text.text)
	if text.mode != types.ModePlain {
		p.set("parse_mode", string(text.mode))
	}
}

type SendPhotoParams struct {
	baseSendMediaParams
}

func (p *SendPhotoParams) Photo(file *InputFile) {
	p.setFile("photo", file)
}

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

type SendDocumentParams struct {
	baseSendMediaParams
}

func (p *SendDocumentParams) Document(file *InputFile) {
	p.setFile("document", file)
}
func (p *SendDocumentParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}

type SendVideoParams struct {
	baseSendMediaParams
}

func (p *SendVideoParams) Video(file *InputFile) {
	p.setFile("video", file)
}
func (p *SendVideoParams) Duration(duration int) {
	p.setInt("duration", duration)
}
func (p *SendVideoParams) Size(width, height int) {
	p.setInt("width", width)
	p.setInt("height", height)
}
func (p *SendVideoParams) Thumb(file *InputFile) {
	p.setFile("thumb", file)
}
func (p *SendVideoParams) SupportsStreaming() {
	p.setBool("supports_streaming", true)
}

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

type SendVoiceParams struct {
	baseSendMediaParams
}

func (p *SendVoiceParams) Voice(file *InputFile) {
	p.setFile("voice", file)
}
func (p *SendVoiceParams) Duration(duration int) {
	p.setInt("duration", duration)
}

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

type SendContactParams struct {
	baseSendParams
}

func (p *SendContactParams) PhoneNumber(phone string) {
	p.set("phone_number", phone)
}
func (p *SendContactParams) Name(firstName, lastName string) {
	p.set("first_name", firstName)
	p.set("last_name", lastName)
}
func (p *SendContactParams) Vcard(vcard string) {
	p.set("vcard", vcard)
}

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
	p.set("explanation", explanation.text)
	if explanation.mode != types.ModePlain {
		p.set("explanation_parse_mode", string(explanation.mode))
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
