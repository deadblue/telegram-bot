package params

import (
	"fmt"
	"io"
)

// The parameters for `forwardMessage`
type ForwardMessageParameters struct {
	ChatParameters
}
func (p *ForwardMessageParameters) FromChatId(chatId int) {
	p.withInt("from_chat_id", chatId)
}
func (p *ForwardMessageParameters) FromChannel(channelName string) {
	p.withString("from_chat_id", fmt.Sprintf("@%s", channelName))
}
func (p *ForwardMessageParameters) MessageId(messageId int) {
	p.withInt("message_id", messageId)
}
func (p *ForwardMessageParameters) DisableNotification() {
	p.withBool("disable_notification", true)
}

// The common parameters for all send methods
type _CommonSendParameters struct {
	ChatParameters
}
func (p *_CommonSendParameters) ReplyToMessageId(messageId int) {
	p.withInt("reply_to_message_id", messageId)
}
func (p *_CommonSendParameters) DisableNotification() {
	p.withBool("disable_notification", true)
}
func (p *_CommonSendParameters) ForceReply(selective bool) {
	markup := map[string]bool{
		"force_reply": true, "selective": selective,
	}
	p.withJson("reply_markup", markup)
}
func (p *_CommonSendParameters) RemoveKeyboard(selective bool) {
	markup := map[string]bool{
		"remove_keyboard": true, "selective": selective,
	}
	p.withJson("reply_markup", markup)
}
func (p *_CommonSendParameters) ReplyKeyboard() *ReplyKeyboardBuilder {
	return &ReplyKeyboardBuilder{
		holder: p._BasicParameters,
		name: "reply_markup",
	}
}
func (p *_CommonSendParameters) InlineKeyboard() *InlineKeyboardBuilder {
	return &InlineKeyboardBuilder{
		holder: p._BasicParameters,
		name: "reply_markup",
	}
}

// The parameters for `sendMessage`
type SendMessageParameters struct {
	_CommonSendParameters
}
func (p *SendMessageParameters) Text(text string) {
	p.withString("text", text)
}
func (p *SendMessageParameters) Markdown(text string) {
	p.withString("text", text).
		withString("parse_mode", parseModeMarkdown)
}
func (p *SendMessageParameters) HTML(text string) {
	p.withString("text", text).
		withString("parse_mode", parseModeHTML)
}
func (p *SendMessageParameters) DisableWebPagePreview() {
	p.withBool("disable_web_page_preview", true)
}

// The common parameters for all send media method
type _CommonSendMediaParameters struct {
	_CommonSendParameters
}
func (p *_CommonSendMediaParameters) Caption(caption string) {
	p.withString("caption", caption)
}
func (p *_CommonSendMediaParameters) MarkdownCaption(caption string) {
	p.withString("caption", caption).
		withString("parse_mode", parseModeMarkdown)
}
func (p *_CommonSendMediaParameters) HTMLCaption(caption string) {
	p.withString("caption", caption).
		withString("parse_mode", parseModeHTML)
}

// The parameters for `sendPhoto`
type SendPhotoParameters struct {
	_CommonSendMediaParameters
}
func (p *SendPhotoParameters) Photo(filename string, filedata io.Reader) {
	p.withFile("photo", filename, filedata)
}
func (p *SendPhotoParameters) PhotoUrl(url string) {
	p.withString("photo", url)
}
func (p *SendPhotoParameters) PhotoFileId(fileId string) {
	p.withString("photo", fileId)
}

// The parameters for `sendAudio`
type SendAudioParameters struct {
	_CommonSendMediaParameters
}
func (p *SendAudioParameters) Audio(filename string, filedata io.Reader) {
	p.withFile("audio", filename, filedata)
}
func (p *SendAudioParameters) AudioFileId(fileId string) {
	p.withString("audio", fileId)
}
func (p *SendAudioParameters) AudioUrl(url string) {
	p.withString("audio", url)
}
func (p *SendAudioParameters) Duration(duration int) {
	p.withInt("duration", duration)
}
func (p *SendAudioParameters) Performer(performer string) {
	p.withString("performer", performer)
}
func (p *SendAudioParameters) Title(title string) {
	p.withString("title", title)
}
func (p *SendAudioParameters) Thumb(filename string, filedata io.Reader) {
	p.withFile("thumb", filename, filedata)
}

// The parameters for `sendDocument`
type SendDocumentParameters struct {
	_CommonSendMediaParameters
}
func (p *SendDocumentParameters) Document(filename string, filedata io.Reader) {
	p.withFile("document", filename, filedata)
}
func (p *SendDocumentParameters) Thumb(filename string, filedata io.Reader) {
	p.withFile("thumb", filename, filedata)
}

// The parameters for `sendVideo`
type SendVideoParameters struct {
	_CommonSendMediaParameters
}
func (p *SendVideoParameters) Video(filename string, filedata io.Reader) {
	p.withFile("video", filename, filedata)
}
func (p *SendVideoParameters) VideoUrl(videoUrl string) {
	p.withString("video", videoUrl)
}
func (p *SendVideoParameters) Duration(duration int) {
	p.withInt("duration", duration)
}
func (p *SendVideoParameters) Size(width, height int) {
	p.withInt("width", width).withInt("height", height)
}
func (p *SendVideoParameters) Thumb(filename string, filedata io.Reader) {
	p.withFile("thumb", filename, filedata)
}
func (p *SendVideoParameters) SupportsStreaming() {
	p.withBool("supports_streaming", true)
}

// The parameters for `sendAnimation`
type SendAnimationParameters struct {
	_CommonSendMediaParameters
}
func (p *SendAnimationParameters) Animation(filename string, filedata io.Reader) {
	p.withFile("animation", filename, filedata)
}
func (p *SendAnimationParameters) Duration(duration int) {
	p.withInt("duration", duration)
}
func (p *SendAnimationParameters) Size(width, height int) {
	p.withInt("width", width).withInt("height", height)
}
func (p *SendAnimationParameters) Thumb(filename string, filedata io.Reader) {
	p.withFile("thumb", filename, filedata)
}

// The parameters for `sendVoice`
type SendVoiceParameters struct {
	_CommonSendMediaParameters
}
func (p *SendVoiceParameters) Voice(filename string, filedata io.Reader) {
	p.withFile("voice", filename, filedata)
}
func (p *SendVoiceParameters) Duration(duration int) {
	p.withInt("duration", duration)
}

// The parameters for `sendVideoNote`
type SendVideoNoteParameters struct {
	_CommonSendMediaParameters
}
func (p *SendVideoNoteParameters) VideoNote(filename string, filedata io.Reader) {
	p.withFile("video_note", filename, filedata)
}
func (p *SendVideoNoteParameters) Duration(duration int) {
	p.withInt("duration", duration)
}
func (p *SendVideoNoteParameters) SideLength(length int) {
	p.withInt("length", length)
}
func (p *SendVideoNoteParameters) Thumb(filename string, filedata io.Reader) {
	p.withFile("thumb", filename, filedata)
}

// The parameters for `sendLocation`
type SendLocationParameters struct {
	_CommonSendParameters
}
func (p *SendLocationParameters) Location(latitude, longitude float64) {
	p.withFloat("latitude", latitude).withFloat("longitude", longitude)
}
func (p *SendLocationParameters) LivePeriod(period int) {
	p.withInt("live_period", period)
}

// The parameters for `sendVenue`
type SendVenueParameters struct {
	_CommonSendParameters
}
func (p *SendVenueParameters) Location(latitude, longitude float64) {
	p.withFloat("latitude", latitude).withFloat("longitude", longitude)
}
func (p *SendVenueParameters) Title(title string) {
	p.withString("title", title)
}
func (p *SendVenueParameters) Address(address string) {
	p.withString("address", address)
}
func (p *SendVenueParameters) Foursquare(fsId, fsType string) {
	p.withString("foursquare_id", fsId).withString("foursquare_type", fsType)
}

// The parameters for `sendContact`
type SendContactParameters struct {
	_CommonSendParameters
}
func (p *SendContactParameters) PhoneNumber(phoneNumber string) {
	p.withString("phone_number", phoneNumber)
}
func (p *SendContactParameters) FirstName(firstName string) {
	p.withString("first_name", firstName)
}
func (p *SendContactParameters) LastName(lastName string) {
	p.withString("last_name", lastName)
}
func (p *SendContactParameters) VCard(vcard string) {
	p.withString("vcard", vcard)
}

type SendPollParameters struct {
	_CommonSendParameters
}
func (p *SendPollParameters) Question(questing string) {
	p.withString("question", questing)
}
func (p *SendPollParameters) Options(options ...string) {
	p.withJson("options", options)
}

// The parameters for `sendChatAction`
type SendChatActionParameters struct {
	ChatParameters
}
func (p *SendChatActionParameters) Typing() {
	p.withString("action", chatActionTyping)
}
func (p *SendChatActionParameters) UploadPhoto() {
	p.withString("action", chatActionUploadPhoto)
}
func (p *SendChatActionParameters) RecordVideo() {
	p.withString("action", chatActionRecordVideo)
}
func (p *SendChatActionParameters) UploadVideo() {
	p.withString("action", chatActionUploadVideo)
}
func (p *SendChatActionParameters) RecordAudio() {
	p.withString("action", chatActionRecordAudio)
}
func (p *SendChatActionParameters) UploadAudio() {
	p.withString("action", chatActionUploadAudio)
}
func (p *SendChatActionParameters) UploadDocument() {
	p.withString("action", chatActionUploadDocument)
}
func (p *SendChatActionParameters) FindLocation() {
	p.withString("action", chatActionFindLocation)
}
func (p *SendChatActionParameters) RecordVideoNote() {
	p.withString("action", chatActionRecordVideoNote)
}
func (p *SendChatActionParameters) UploadVideoNote() {
	p.withString("action", chatActionUploadVideoNote)
}

// The parameters for `deleteMessage`
type ChatMessageParameters struct {
	ChatParameters
}
func (p *ChatMessageParameters) MessageId(messageId int) {
	p.withInt("message_id", messageId)
}


type EditMessageReplyMarkupParameters struct {
	ChatMessageParameters
}
func (p *EditMessageReplyMarkupParameters) InlineMessageId(inlineMessageId int) {
	p.withInt("inline_message_id", inlineMessageId)
}
func (p *EditMessageMediaParameters) InlineKeyboard() {
	// TODO construct inline keyboard
}


type EditMessageTextParameters struct {
	EditMessageReplyMarkupParameters
}
func (p *EditMessageTextParameters) Text(text string) {
	p.withString("text", text)
}
func (p *EditMessageTextParameters) Markdown(text string) {
	p.withString("text", text).
		withString("parse_mode", parseModeMarkdown)
}
func (p *EditMessageTextParameters) HTML(text string) {
	p.withString("text", text).
		withString("parse_mode", parseModeHTML)
}
func (p *EditMessageTextParameters) DisableWebPagePreview() {
	p.withBool("disable_web_page_preview", true)
}


type EditMessageCaptionParameters struct {
	EditMessageReplyMarkupParameters
}
func (p *EditMessageCaptionParameters) Caption(caption string) {
	p.withString("caption", caption)
}
func (p *EditMessageCaptionParameters) MarkdownCaption(caption string) {
	p.withString("caption", caption).
		withString("parse_mode", parseModeMarkdown)
}
func (p *EditMessageCaptionParameters) HTMLCaption(caption string) {
	p.withString("caption", caption).
		withString("parse_mode", parseModeHTML)
}


type EditMessageLiveLocationParameters struct {
	EditMessageReplyMarkupParameters
}
func (p *EditMessageLiveLocationParameters) Location(latitude, longitude float64) {
	p.withFloat("latitude", latitude).withFloat("longitude", longitude)
}

type StopPollParameters struct {
	ChatMessageParameters
}
func (p *StopPollParameters) InlineKeyboard() {
	// TODO construct inline keyboard
}
