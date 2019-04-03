package telegroid

import (
	"fmt"
	"io"
)

const (
	// for parse mode
	parseModeMarkdown = "Markdown"
	parseModeHTML     = "HTML"

	// chat action
	chatActionTyping          = "typing"
	chatActionUploadPhoto     = "upload_photo"
	chatActionRecordVideo     = "record_video"
	chatActionUploadVideo     = "upload_video"
	chatActionRecordAudio     = "record_audio"
	chatActionUploadAudio     = "upload_audio"
	chatActionUploadDocument  = "upload_document"
	chatActionFindLocation    = "find_location"
	chatActionRecordVideoNote = "record_video_note"
	chatActionUploadVideoNote = "upload_video_note"
)

type ChatParameters struct {
	_BasicParameters
}

func (p *ChatParameters) ChatId(chatId int) {
	p.withInt("chat_id", chatId)
}
func (p *ChatParameters) Channel(channelName string) {
	p.withString("chat_id", fmt.Sprintf("@%s", channelName))
}

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

type _BasicSendParameters struct {
	ChatParameters
}

func (p *_BasicSendParameters) ReplyToMessageId(messageId int) {
	p.withInt("reply_to_message_id", messageId)
}
func (p *_BasicSendParameters) DisableNotification() {
	p.withBool("disable_notification", true)
}
func (p *_BasicSendParameters) ForceReply(selective bool) {
	markup := map[string]bool{
		"force_reply": true, "selective": selective,
	}
	p.withJson("reply_markup", markup)
}
func (p *_BasicSendParameters) RemoveKeyboard(selective bool) {
	markup := map[string]bool{
		"remove_keyboard": true, "selective": selective,
	}
	p.withJson("reply_markup", markup)
}
func (p *_BasicSendParameters) ReplyKeyboard() {
	// TODO
}
func (p *_BasicSendParameters) InlineKeyboard() {
	// TODO
}

type SendMessageParameters struct {
	_BasicSendParameters
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

type _SendMediaParameters struct {
	_BasicSendParameters
}

func (p *_SendMediaParameters) Caption(caption string) {
	p.withString("caption", caption)
}
func (p *_SendMediaParameters) MarkdownCaption(caption string) {
	p.withString("caption", caption).
		withString("parse_mode", parseModeMarkdown)
}
func (p *_SendMediaParameters) HTMLCaption(caption string) {
	p.withString("caption", caption).
		withString("parse_mode", parseModeHTML)
}

type SendPhotoParameters struct {
	_SendMediaParameters
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

type SendAudioParameters struct {
	_SendMediaParameters
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

type SendDocumentParameters struct {
	_SendMediaParameters
}

func (p *SendDocumentParameters) Document(filename string, filedata io.Reader) {
	p.withFile("document", filename, filedata)
}
func (p *SendDocumentParameters) Thumb(filename string, filedata io.Reader) {
	p.withFile("thumb", filename, filedata)
}

type SendVideoParameters struct {
	_SendMediaParameters
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

type SendAnimationParameters struct {
	_SendMediaParameters
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

type SendVoiceParameters struct {
	_SendMediaParameters
}

func (p *SendVoiceParameters) Voice(filename string, filedata io.Reader) {
	p.withFile("voice", filename, filedata)
}
func (p *SendVoiceParameters) Duration(duration int) {
	p.withInt("duration", duration)
}

type SendVideoNoteParameters struct {
	_SendMediaParameters
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

type SendLocationParameters struct {
	_BasicSendParameters
}

func (p *SendLocationParameters) Location(latitude, longitude float64) {
	p.withFloat("latitude", latitude).withFloat("longitude", longitude)
}
func (p *SendLocationParameters) LivePeriod(period int) {
	p.withInt("live_period", period)
}


type SendVenueParameters struct {
	_BasicSendParameters
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


type SendContactParameters struct {
	_BasicSendParameters
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

type DeleteMessageParameters struct {
	ChatParameters
}

func (p *DeleteMessageParameters) MessageId(messageId int) {
	p.withInt("message_id", messageId)
}

type EditMessageReplyMarkupParameters struct {
	DeleteMessageParameters
}

func (p *EditMessageReplyMarkupParameters) InlineMessageId(inlineMessageId int) {
	p.withInt("inline_message_id", inlineMessageId)
}
func (p *EditMessageMediaParameters) InlineKeyboard() {
	// TODO
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

type EditMessageMediaParameters struct {
	EditMessageReplyMarkupParameters
}

func (p *EditMessageMediaParameters) Media() {
	// TODO
}

type EditMessageLiveLocationParameters struct {
	EditMessageReplyMarkupParameters
}
func (p *EditMessageLiveLocationParameters) Location(latitude, longitude float64) {
	p.withFloat("latitude", latitude).withFloat("longitude", longitude)
}

