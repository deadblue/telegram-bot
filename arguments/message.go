package arguments

import (
	"fmt"
	"io"
)

type ForwardMessageArgs struct {
	ChatArgs
}
func (a *ForwardMessageArgs) FromChatId(chatId int) {
	a.withInt("from_chat_id", chatId)
}
func (a *ForwardMessageArgs) FromChannel(channelName string) {
	a.withString("from_chat_id", fmt.Sprintf("@%s", channelName))
}
func (a *ForwardMessageArgs) MessageId(messageId int) {
	a.withInt("message_id", messageId)
}
func (a *ForwardMessageArgs) DisableNotification() {
	a.withBool("disable_notification", true)
}

type _CommonSendArgs struct {
	ChatArgs
}
func (a *_CommonSendArgs) ReplyToMessageId(messageId int) {
	a.withInt("reply_to_message_id", messageId)
}
func (a *_CommonSendArgs) DisableNotification() {
	a.withBool("disable_notification", true)
}
func (a *_CommonSendArgs) ForceReply(selective bool) {
	markup := map[string]bool{
		"force_reply": true, "selective": selective,
	}
	a.withJson("reply_markup", markup)
}
func (a *_CommonSendArgs) RemoveKeyboard(selective bool) {
	markup := map[string]bool{
		"remove_keyboard": true, "selective": selective,
	}
	a.withJson("reply_markup", markup)
}
func (a *_CommonSendArgs) ReplyKeyboard() *ReplyKeyboardBuilder {
	return &ReplyKeyboardBuilder{
		holder: a._BasicArgs,
		name:   "reply_markup",
	}
}
func (a *_CommonSendArgs) InlineKeyboard() *InlineKeyboardBuilder {
	return &InlineKeyboardBuilder{
		holder: a._BasicArgs,
		name:   "reply_markup",
	}
}

type SendMessageArgs struct {
	_CommonSendArgs
}
func (a *SendMessageArgs) Text(text string) {
	a.withString("text", text)
}
func (a *SendMessageArgs) Markdown(text string) {
	a.withString("text", text).
		withString("parse_mode", parseModeMarkdown)
}
func (a *SendMessageArgs) HTML(text string) {
	a.withString("text", text).
		withString("parse_mode", parseModeHTML)
}
func (a *SendMessageArgs) DisableWebPagePreview() {
	a.withBool("disable_web_page_preview", true)
}

type _CommonSendMediaArgs struct {
	_CommonSendArgs
}
func (a *_CommonSendMediaArgs) Caption(caption string) {
	a.withString("caption", caption)
}
func (a *_CommonSendMediaArgs) MarkdownCaption(caption string) {
	a.withString("caption", caption).
		withString("parse_mode", parseModeMarkdown)
}
func (a *_CommonSendMediaArgs) HTMLCaption(caption string) {
	a.withString("caption", caption).
		withString("parse_mode", parseModeHTML)
}

type SendPhotoArgs struct {
	_CommonSendMediaArgs
}
func (a *SendPhotoArgs) Photo(filename string, filedata io.Reader) {
	a.withFile("photo", filename, filedata)
}
func (a *SendPhotoArgs) PhotoUrl(url string) {
	a.withString("photo", url)
}
func (a *SendPhotoArgs) PhotoFileId(fileId string) {
	a.withString("photo", fileId)
}

type SendAudioArgs struct {
	_CommonSendMediaArgs
}
func (a *SendAudioArgs) Audio(filename string, filedata io.Reader) {
	a.withFile("audio", filename, filedata)
}
func (a *SendAudioArgs) AudioFileId(fileId string) {
	a.withString("audio", fileId)
}
func (a *SendAudioArgs) AudioUrl(url string) {
	a.withString("audio", url)
}
func (a *SendAudioArgs) Duration(duration int) {
	a.withInt("duration", duration)
}
func (a *SendAudioArgs) Performer(performer string) {
	a.withString("performer", performer)
}
func (a *SendAudioArgs) Title(title string) {
	a.withString("title", title)
}
func (a *SendAudioArgs) Thumb(filename string, filedata io.Reader) {
	a.withFile("thumb", filename, filedata)
}

type SendDocumentArgs struct {
	_CommonSendMediaArgs
}
func (a *SendDocumentArgs) Document(filename string, filedata io.Reader) {
	a.withFile("document", filename, filedata)
}
func (a *SendDocumentArgs) Thumb(filename string, filedata io.Reader) {
	a.withFile("thumb", filename, filedata)
}

type SendVideoArgs struct {
	_CommonSendMediaArgs
}
func (a *SendVideoArgs) Video(filename string, filedata io.Reader) {
	a.withFile("video", filename, filedata)
}
func (a *SendVideoArgs) VideoUrl(videoUrl string) {
	a.withString("video", videoUrl)
}
func (a *SendVideoArgs) Duration(duration int) {
	a.withInt("duration", duration)
}
func (a *SendVideoArgs) Size(width, height int) {
	a.withInt("width", width).withInt("height", height)
}
func (a *SendVideoArgs) Thumb(filename string, filedata io.Reader) {
	a.withFile("thumb", filename, filedata)
}
func (a *SendVideoArgs) SupportsStreaming() {
	a.withBool("supports_streaming", true)
}

type SendAnimationArgs struct {
	_CommonSendMediaArgs
}
func (a *SendAnimationArgs) Animation(filename string, filedata io.Reader) {
	a.withFile("animation", filename, filedata)
}
func (a *SendAnimationArgs) Duration(duration int) {
	a.withInt("duration", duration)
}
func (a *SendAnimationArgs) Size(width, height int) {
	a.withInt("width", width).withInt("height", height)
}
func (a *SendAnimationArgs) Thumb(filename string, filedata io.Reader) {
	a.withFile("thumb", filename, filedata)
}

type SendVoiceArgs struct {
	_CommonSendMediaArgs
}
func (a *SendVoiceArgs) Voice(filename string, filedata io.Reader) {
	a.withFile("voice", filename, filedata)
}
func (a *SendVoiceArgs) Duration(duration int) {
	a.withInt("duration", duration)
}

type SendVideoNoteArgs struct {
	_CommonSendMediaArgs
}
func (a *SendVideoNoteArgs) VideoNote(filename string, filedata io.Reader) {
	a.withFile("video_note", filename, filedata)
}
func (a *SendVideoNoteArgs) Duration(duration int) {
	a.withInt("duration", duration)
}
func (a *SendVideoNoteArgs) SideLength(length int) {
	a.withInt("length", length)
}
func (a *SendVideoNoteArgs) Thumb(filename string, filedata io.Reader) {
	a.withFile("thumb", filename, filedata)
}

type SendLocationArgs struct {
	_CommonSendArgs
}
func (a *SendLocationArgs) Location(latitude, longitude float64) {
	a.withFloat("latitude", latitude).withFloat("longitude", longitude)
}
func (a *SendLocationArgs) LivePeriod(period int) {
	a.withInt("live_period", period)
}

type SendVenueArgs struct {
	_CommonSendArgs
}
func (a *SendVenueArgs) Location(latitude, longitude float64) {
	a.withFloat("latitude", latitude).withFloat("longitude", longitude)
}
func (a *SendVenueArgs) Title(title string) {
	a.withString("title", title)
}
func (a *SendVenueArgs) Address(address string) {
	a.withString("address", address)
}
func (a *SendVenueArgs) Foursquare(fsId, fsType string) {
	a.withString("foursquare_id", fsId).withString("foursquare_type", fsType)
}

type SendContactArgs struct {
	_CommonSendArgs
}
func (a *SendContactArgs) PhoneNumber(phoneNumber string) {
	a.withString("phone_number", phoneNumber)
}
func (a *SendContactArgs) FirstName(firstName string) {
	a.withString("first_name", firstName)
}
func (a *SendContactArgs) LastName(lastName string) {
	a.withString("last_name", lastName)
}
func (a *SendContactArgs) VCard(vcard string) {
	a.withString("vcard", vcard)
}

type SendPollArgs struct {
	_CommonSendArgs
}
func (a *SendPollArgs) Question(questing string) {
	a.withString("question", questing)
}
func (a *SendPollArgs) Options(options ...string) {
	a.withJson("options", options)
}

type SendChatActionArgs struct {
	ChatArgs
}
func (a *SendChatActionArgs) Typing() {
	a.withString("action", chatActionTyping)
}
func (a *SendChatActionArgs) UploadPhoto() {
	a.withString("action", chatActionUploadPhoto)
}
func (a *SendChatActionArgs) RecordVideo() {
	a.withString("action", chatActionRecordVideo)
}
func (a *SendChatActionArgs) UploadVideo() {
	a.withString("action", chatActionUploadVideo)
}
func (a *SendChatActionArgs) RecordAudio() {
	a.withString("action", chatActionRecordAudio)
}
func (a *SendChatActionArgs) UploadAudio() {
	a.withString("action", chatActionUploadAudio)
}
func (a *SendChatActionArgs) UploadDocument() {
	a.withString("action", chatActionUploadDocument)
}
func (a *SendChatActionArgs) FindLocation() {
	a.withString("action", chatActionFindLocation)
}
func (a *SendChatActionArgs) RecordVideoNote() {
	a.withString("action", chatActionRecordVideoNote)
}
func (a *SendChatActionArgs) UploadVideoNote() {
	a.withString("action", chatActionUploadVideoNote)
}

// The parameters for `deleteMessage`
type ChatMessageArgs struct {
	ChatArgs
}
func (a *ChatMessageArgs) MessageId(messageId int) {
	a.withInt("message_id", messageId)
}


type EditMessageReplyMarkupArgs struct {
	ChatMessageArgs
}
func (a *EditMessageReplyMarkupArgs) InlineMessageId(inlineMessageId int) {
	a.withInt("inline_message_id", inlineMessageId)
}
func (a *EditMessageReplyMarkupArgs) InlineKeyboard() {
	// TODO construct inline keyboard
}


type EditMessageTextArgs struct {
	EditMessageReplyMarkupArgs
}
func (a *EditMessageTextArgs) Text(text string) {
	a.withString("text", text)
}
func (a *EditMessageTextArgs) Markdown(text string) {
	a.withString("text", text).
		withString("parse_mode", parseModeMarkdown)
}
func (a *EditMessageTextArgs) HTML(text string) {
	a.withString("text", text).
		withString("parse_mode", parseModeHTML)
}
func (a *EditMessageTextArgs) DisableWebPagePreview() {
	a.withBool("disable_web_page_preview", true)
}


type EditMessageCaptionArgs struct {
	EditMessageReplyMarkupArgs
}
func (a *EditMessageCaptionArgs) Caption(caption string) {
	a.withString("caption", caption)
}
func (a *EditMessageCaptionArgs) MarkdownCaption(caption string) {
	a.withString("caption", caption).
		withString("parse_mode", parseModeMarkdown)
}
func (a *EditMessageCaptionArgs) HTMLCaption(caption string) {
	a.withString("caption", caption).
		withString("parse_mode", parseModeHTML)
}

type EditMessageLiveLocationArgs struct {
	EditMessageReplyMarkupArgs
}
func (a *EditMessageLiveLocationArgs) Location(latitude, longitude float64) {
	a.withFloat("latitude", latitude).withFloat("longitude", longitude)
}

type StopPollArgs struct {
	ChatMessageArgs
}
func (a *StopPollArgs) InlineKeyboard() {
	// TODO construct inline keyboard
}
