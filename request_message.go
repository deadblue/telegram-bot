package telegroid

import (
	"fmt"
	"io"
)

const (
	parseModeMarkdown = "Markdown"
	parseModeHTML     = "HTML"
)


// The request entity for Forward request
type ForwardMessageRequest struct {
	ChatRequest
}
func (r *ForwardMessageRequest) FromChatId(chatId int) {
	r._WithInt("from_chat_id", chatId)
}
func (r *ForwardMessageRequest) FromChannel(channelName string) {
	r._WithString("from_chat_id", fmt.Sprintf("@%s", channelName))
}
func (r *ForwardMessageRequest) MessageId(messageId int) {
	r._WithInt("message_id", messageId)
}
func (r *ForwardMessageRequest) DisableNotification() {
	r._WithBool("disable_notification", true)
}


//
type _BasicSendRequest struct {
	ChatRequest
}
func (r *_BasicSendRequest) ReplyToMessageId(messageId int) {
	r._WithInt("reply_to_message_id", messageId)
}
func (r *_BasicSendRequest) DisableNotification() {
	r._WithBool("disable_notification", true)
}
func (r *_BasicSendRequest) ForceReply(selective bool) {
	markup := map[string]bool{
		"force_reply": true, "selective": selective,
	}
	r._WithJson("reply_markup", markup)
}
func (r *_BasicSendRequest) RemoveKeyboard(selective bool) {
	markup := map[string]bool{
		"remove_keyboard": true, "selective": selective,
	}
	r._WithJson("reply_markup", markup)
}
// TODO: ReplyKeyboard & InlineKeyboard


type SendMessageRequest struct {
	_BasicSendRequest
}
func (r *SendMessageRequest) Text(text string) {
	r._WithString("text", text)
}
func (r *SendMessageRequest) Markdown(text string) {
	r._WithString("text", text).
		_WithString("parse_mode", parseModeMarkdown)
}
func (r *SendMessageRequest) HTML(text string) {
	r._WithString("text", text).
		_WithString("parse_mode", parseModeHTML)
}
func (r *SendMessageRequest) DisableWebPagePreview() {
	r._WithBool("disable_web_page_preview", true)
}


type _SendMediaRequest struct {
	_BasicSendRequest
}
func (r *_SendMediaRequest) Caption(caption string) {
	r._WithString("caption", caption)
}
func (r *_SendMediaRequest) MarkdownCaption(caption string) {
	r._WithString("caption", caption).
		_WithString("parse_mode", parseModeMarkdown)
}
func (r *_SendMediaRequest) HTMLCaption(caption string) {
	r._WithString("caption", caption).
		_WithString("parse_mode", parseModeHTML)
}


type SendPhotoRequest struct {
	_SendMediaRequest
}
func (r *SendPhotoRequest) Photo(filename string, filedata io.Reader) {
	r._WithFile("photo", filename, filedata)
}
func (r *SendPhotoRequest) PhotoUrl(url string) {
	r._WithString("photo", url)
}
func (r *SendPhotoRequest) PhotoFileId(fileId string) {
	r._WithString("photo", fileId)
}


type SendAudioRequest struct {
	_SendMediaRequest
}
func (r *SendAudioRequest) Audio(filename string, filedata io.Reader) {
	r._WithFile("audio", filename, filedata)
}
func (r *SendAudioRequest) AudioFileId(fileId string) {
	r._WithString("audio", fileId)
}
func (r *SendAudioRequest) AudioUrl(url string) {
	r._WithString("audio", url)
}
func (r *SendAudioRequest) Duration(duration int) {
	r._WithInt("duration", duration)
}
func (r *SendAudioRequest) Performer(performer string) {
	r._WithString("performer", performer)
}
func (r *SendAudioRequest) Title(title string) {
	r._WithString("title", title)
}
func (r *SendAudioRequest) Thumb(filename string, filedata io.Reader) {
	r._WithFile("thumb", filename, filedata)
}


type SendDocumentRequest struct {
	_SendMediaRequest
}
func (r *SendDocumentRequest) Document(filename string, filedata io.Reader) {
	r._WithFile("document", filename, filedata)
}
func (r *SendDocumentRequest) Thumb(filename string, filedata io.Reader) {
	r._WithFile("thumb", filename, filedata)
}


type SendVideoRequest struct {
	_SendMediaRequest
}
func (r *SendVideoRequest) Video(filename string, filedata io.Reader)  {
	r._WithFile("video", filename, filedata)
}
func (r *SendVideoRequest) Duration(duration int) {
	r._WithInt("duration", duration)
}
func (r *SendVideoRequest) Size(width, height int) {
	r._WithInt("width", width)._WithInt("height", height)
}
func (r *SendVideoRequest) Thumb(filename string, filedata io.Reader) {
	r._WithFile("thumb", filename, filedata)
}
func (r *SendVideoRequest) SupportsStreaming() {
	r._WithBool("supports_streaming", true)
}


type SendAnimationRequest struct {
	_SendMediaRequest
}
func (r *SendAnimationRequest) Animation(filename string, filedata io.Reader) {
	r._WithFile("animation", filename, filedata)
}
func (r *SendAnimationRequest) Duration(duration int) {
	r._WithInt("duration", duration)
}
func (r *SendAnimationRequest) Size(width, height int) {
	r._WithInt("width", width)._WithInt("height", height)
}
func (r *SendAnimationRequest) Thumb(filename string, filedata io.Reader) {
	r._WithFile("thumb", filename, filedata)
}


type SendVoiceRequest struct {
	_SendMediaRequest
}
func (r *SendVoiceRequest) Voice(filename string, filedata io.Reader) {
	r._WithFile("voice", filename, filedata)
}
func (r *SendVoiceRequest) Duration(duration int) {
	r._WithInt("duration", duration)
}


type SendVideoNoteRequest struct {
	_SendMediaRequest
}
func (r *SendVideoNoteRequest) VideoNote(filename string, filedata io.Reader) {
	r._WithFile("video_note", filename, filedata)
}
func (r *SendVideoNoteRequest) Duration(duration int) {
	r._WithInt("duration", duration)
}
func (r *SendVideoNoteRequest) SideLength(length int) {
	r._WithInt("length", length)
}
func (r *SendVideoNoteRequest) Thumb(filename string, filedata io.Reader) {
	r._WithFile("thumb", filename, filedata)
}


type SendLocationRequest struct {
	_BasicSendRequest
}
func (r *SendLocationRequest) Location(latitude, longitude float64) {
	r._WithFloat("latitude", latitude)._WithFloat("longitude", longitude)
}
func (r *SendLocationRequest) LivePeriod(period int) {
	r._WithInt("live_period", period)
}





