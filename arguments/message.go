package arguments

import (
	"fmt"
	"io"
)

type ForwardMessageArgs struct {
	_ChatArgs
}

func (a *ForwardMessageArgs) FromChatId(chatId int) {
	a.getForm().WithInt("from_chat_id", chatId)
}
func (a *ForwardMessageArgs) FromChannel(channelName string) {
	a.getForm().WithString("from_chat_id", fmt.Sprintf("@%s", channelName))
}
func (a *ForwardMessageArgs) MessageId(messageId int) {
	a.getForm().WithInt("message_id", messageId)
}
func (a *ForwardMessageArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}

type _CommonSendArgs struct {
	_ChatArgs
}

func (a *_CommonSendArgs) ReplyToMessageId(messageId int) {
	a.getForm().WithInt("reply_to_message_id", messageId)
}
func (a *_CommonSendArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}
func (a *_CommonSendArgs) ForceReply(selective bool) {
	markup := map[string]bool{
		"force_reply": true, "selective": selective,
	}
	a.getForm().WithJson("reply_markup", markup)
}
func (a *_CommonSendArgs) RemoveKeyboard(selective bool) {
	markup := map[string]bool{
		"remove_keyboard": true, "selective": selective,
	}
	a.getForm().WithJson("reply_markup", markup)
}
func (a *_CommonSendArgs) ReplyKeyboard() ReplyKeyboardBuilder {
	return &implReplyKeyboardBuilder{
		form: a.getForm(),
		name: "reply_markup",
	}
}
func (a *_CommonSendArgs) InlineKeyboard() InlineKeyboardBuilder {
	return &implInlineKeyboardBuilder{
		form: a.getForm(),
		name: "reply_markup",
	}
}

type SendMessageArgs struct {
	_CommonSendArgs
}

func (a *SendMessageArgs) Text(text string) ParseModeSelector {
	form := a.getForm()
	form.WithString("text", text)
	return &implParseModeSelector{form: form}
}
func (a *SendMessageArgs) DisableWebPagePreview() {
	a.getForm().WithBool("disable_web_page_preview", true)
}

type _CommonSendMediaArgs struct {
	_CommonSendArgs
}

func (a *_CommonSendMediaArgs) Caption(caption string) ParseModeSelector {
	form := a.getForm()
	form.WithString("caption", caption)
	return &implParseModeSelector{form: form}
}

type SendPhotoArgs struct {
	_CommonSendMediaArgs
}

func (a *SendPhotoArgs) Photo(file InputFile) {
	a.getForm().WithFile("photo", file)
}
func (a *SendPhotoArgs) PhotoUrl(url string) {
	a.getForm().WithString("photo", url)
}
func (a *SendPhotoArgs) PhotoFileId(fileId string) {
	a.getForm().WithString("photo", fileId)
}

type SendAudioArgs struct {
	_CommonSendMediaArgs
}

func (a *SendAudioArgs) Audio(file InputFile) {
	a.getForm().WithFile("audio", file)
}
func (a *SendAudioArgs) AudioFileId(fileId string) {
	a.getForm().WithString("audio", fileId)
}
func (a *SendAudioArgs) AudioUrl(url string) {
	a.getForm().WithString("audio", url)
}
func (a *SendAudioArgs) Duration(duration int) {
	a.getForm().WithInt("duration", duration)
}
func (a *SendAudioArgs) Performer(performer string) {
	a.getForm().WithString("performer", performer)
}
func (a *SendAudioArgs) Title(title string) {
	a.getForm().WithString("title", title)
}
func (a *SendAudioArgs) Thumb(file InputFile) {
	a.getForm().WithFile("thumb", file)
}

type SendDocumentArgs struct {
	_CommonSendMediaArgs
}

func (a *SendDocumentArgs) Document(file InputFile) {
	a.getForm().WithFile("document", file)
}
func (a *SendDocumentArgs) Thumb(file InputFile) {
	a.getForm().WithFile("thumb", file)
}

type SendVideoArgs struct {
	_CommonSendMediaArgs
}

func (a *SendVideoArgs) Video(file InputFile) {
	a.getForm().WithFile("video", file)
}
func (a *SendVideoArgs) VideoUrl(videoUrl string) {
	a.getForm().WithString("video", videoUrl)
}
func (a *SendVideoArgs) Duration(duration int) {
	a.getForm().WithInt("duration", duration)
}
func (a *SendVideoArgs) Size(width, height int) {
	a.getForm().
		WithInt("width", width).
		WithInt("height", height)
}
func (a *SendVideoArgs) Thumb(file InputFile) {
	a.getForm().WithFile("thumb", file)
}
func (a *SendVideoArgs) SupportsStreaming() {
	a.getForm().WithBool("supports_streaming", true)
}

type SendAnimationArgs struct {
	_CommonSendMediaArgs
}

func (a *SendAnimationArgs) Animation(file InputFile) {
	a.getForm().WithFile("animation", file)
}
func (a *SendAnimationArgs) Duration(duration int) {
	a.getForm().WithInt("duration", duration)
}
func (a *SendAnimationArgs) Size(width, height int) {
	a.getForm().
		WithInt("width", width).
		WithInt("height", height)
}
func (a *SendAnimationArgs) Thumb(file InputFile) {
	a.getForm().WithFile("thumb", file)
}

type SendVoiceArgs struct {
	_CommonSendMediaArgs
}

func (a *SendVoiceArgs) Voice(file InputFile) {
	a.getForm().WithFile("voice", file)
}
func (a *SendVoiceArgs) Duration(duration int) {
	a.getForm().WithInt("duration", duration)
}

type SendVideoNoteArgs struct {
	_CommonSendMediaArgs
}

func (a *SendVideoNoteArgs) VideoNote(file InputFile) {
	a.getForm().WithFile("video_note", file)
}
func (a *SendVideoNoteArgs) Duration(duration int) {
	a.getForm().WithInt("duration", duration)
}
func (a *SendVideoNoteArgs) SideLength(length int) {
	a.getForm().WithInt("length", length)
}
func (a *SendVideoNoteArgs) Thumb(file InputFile) {
	a.getForm().WithFile("thumb", file)
}

type SendMediaGroupArgs struct {
	_ChatArgs

	mediaBuf   []map[string]interface{}
	mediaCount int
}

func (a *SendMediaGroupArgs) receiveMedia(media map[string]interface{}) {
	if a.mediaBuf == nil {
		a.mediaBuf = make([]map[string]interface{}, 10)
		a.mediaCount = 0
	}
	// Do not accpet more than 10 media
	if a.mediaCount >= 10 {
		return
	}
	// store media item
	a.mediaBuf[a.mediaCount] = media
	a.mediaCount += 1
}

// Add a photo media to the group.
// According to the API limit, you can add
// at most 10 photo/video media to one group.
func (a *SendMediaGroupArgs) MediaPhoto() MediaPhotoBuilder {
	b := new(implMediaPhotoBuilder)
	b.Init(a.receiveMedia)
	return b
}

// Add a video media to the group.
// According to the API limit, you can add
// at most 10 photo/video media to one group.
func (a *SendMediaGroupArgs) MediaVideo() MediaVideoBuilder {
	b := new(implMediaVideoBuilder)
	b.Init(a.receiveMedia)
	return b
}
func (a *SendMediaGroupArgs) ReplyToMessageId(messageId int) {
	a.getForm().WithInt("reply_to_message_id", messageId)
}
func (a *SendMediaGroupArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}
func (a *SendMediaGroupArgs) Archive() (contentType string, data io.Reader) {
	form := a.getForm()
	for i := 0; i < a.mediaCount; i++ {
		item := a.mediaBuf[i]
		// Upload media file
		mf, ok := item["media"].(InputFile)
		if ok {
			attachName := fmt.Sprintf("attach_media_%d", i)
			item["media"] = fmt.Sprintf("attach://%s", attachName)
			form.WithFile(attachName, mf)
		}
		// Upload thumb file
		tf, ok := item["thumb"].(InputFile)
		if ok {
			attachName := fmt.Sprintf("attach_thumb_%d", i)
			item["thumb"] = fmt.Sprintf("attach://%s", attachName)
			form.WithFile(attachName, tf)
		}
	}
	form.WithJson("media", a.mediaBuf[:a.mediaCount])
	return form.Close()
}

type SendLocationArgs struct {
	_CommonSendArgs
}

func (a *SendLocationArgs) Location(latitude, longitude float64) {
	a.getForm().WithFloat("latitude", latitude).WithFloat("longitude", longitude)
}
func (a *SendLocationArgs) LivePeriod(period int) {
	a.getForm().WithInt("live_period", period)
}

type SendVenueArgs struct {
	_CommonSendArgs
}

func (a *SendVenueArgs) Location(latitude, longitude float64) {
	a.getForm().
		WithFloat("latitude", latitude).
		WithFloat("longitude", longitude)
}
func (a *SendVenueArgs) Title(title string) {
	a.getForm().WithString("title", title)
}
func (a *SendVenueArgs) Address(address string) {
	a.getForm().WithString("address", address)
}
func (a *SendVenueArgs) Foursquare(fsId, fsType string) {
	a.getForm().
		WithString("foursquare_id", fsId).
		WithString("foursquare_type", fsType)
}

type SendContactArgs struct {
	_CommonSendArgs
}

func (a *SendContactArgs) PhoneNumber(phoneNumber string) {
	a.getForm().WithString("phone_number", phoneNumber)
}
func (a *SendContactArgs) FirstName(firstName string) {
	a.getForm().WithString("first_name", firstName)
}
func (a *SendContactArgs) LastName(lastName string) {
	a.getForm().WithString("last_name", lastName)
}
func (a *SendContactArgs) VCard(vcard string) {
	a.getForm().WithString("vcard", vcard)
}

type SendPollArgs struct {
	_CommonSendArgs
}

func (a *SendPollArgs) Question(questing string) {
	a.getForm().WithString("question", questing)
}
func (a *SendPollArgs) Options(options ...string) {
	a.getForm().WithJson("options", options)
}

type SendChatActionArgs struct {
	_ChatArgs
}

func (a *SendChatActionArgs) Typing() {
	a.getForm().WithString("action", chatActionTyping)
}
func (a *SendChatActionArgs) UploadPhoto() {
	a.getForm().WithString("action", chatActionUploadPhoto)
}
func (a *SendChatActionArgs) RecordVideo() {
	a.getForm().WithString("action", chatActionRecordVideo)
}
func (a *SendChatActionArgs) UploadVideo() {
	a.getForm().WithString("action", chatActionUploadVideo)
}
func (a *SendChatActionArgs) RecordAudio() {
	a.getForm().WithString("action", chatActionRecordAudio)
}
func (a *SendChatActionArgs) UploadAudio() {
	a.getForm().WithString("action", chatActionUploadAudio)
}
func (a *SendChatActionArgs) UploadDocument() {
	a.getForm().WithString("action", chatActionUploadDocument)
}
func (a *SendChatActionArgs) FindLocation() {
	a.getForm().WithString("action", chatActionFindLocation)
}
func (a *SendChatActionArgs) RecordVideoNote() {
	a.getForm().WithString("action", chatActionRecordVideoNote)
}
func (a *SendChatActionArgs) UploadVideoNote() {
	a.getForm().WithString("action", chatActionUploadVideoNote)
}

type _ChatMessageArgs struct {
	_ChatArgs
}

func (a *_ChatMessageArgs) MessageId(messageId int) {
	a.getForm().WithInt("message_id", messageId)
}

type ChatMessageArgs _ChatMessageArgs

type _EditMessageReplyMarkupArgs struct {
	_ChatMessageArgs
}

func (a *_EditMessageReplyMarkupArgs) InlineMessageId(inlineMessageId int) {
	a.getForm().WithInt("inline_message_id", inlineMessageId)
}
func (a *_EditMessageReplyMarkupArgs) InlineKeyboard() InlineKeyboardBuilder {
	return &implInlineKeyboardBuilder{
		form: a.getForm(),
		name: "reply_markup",
	}
}

type EditMessageReplyMarkupArgs _EditMessageReplyMarkupArgs

type EditMessageTextArgs struct {
	_EditMessageReplyMarkupArgs
}

func (a *EditMessageTextArgs) Text(text string) ParseModeSelector {
	form := a.getForm()
	form.WithString("text", text)
	return &implParseModeSelector{form: form}
}
func (a *EditMessageTextArgs) DisableWebPagePreview() {
	a.getForm().WithBool("disable_web_page_preview", true)
}

type EditMessageCaptionArgs struct {
	_EditMessageReplyMarkupArgs
}

func (a *EditMessageCaptionArgs) Caption(caption string) ParseModeSelector {
	form := a.getForm()
	form.WithString("caption", caption)
	return &implParseModeSelector{form: form}
}

type EditMessageMediaArgs struct {
	_EditMessageReplyMarkupArgs
}

func (a *EditMessageMediaArgs) receiveMedia(media map[string]interface{}) {
	if media == nil {
		return
	}
	form := a.getForm()
	// Upload media file
	mf, ok := media["media"].(InputFile)
	if ok {
		attachName := "attach_media_file"
		media["media"] = fmt.Sprintf("attach://%s", attachName)
		form.WithFile(attachName, mf)
	}
	// Upload thumb file
	tf, ok := media["thumb"].(InputFile)
	if ok {
		attachName := "attach_thumb_file"
		media["thumb"] = fmt.Sprintf("attach://%s", attachName)
		form.WithFile(attachName, tf)
	}
	form.WithJson("media", media)
}
func (a *EditMessageMediaArgs) MediaPhoto() MediaPhotoBuilder {
	b := new(implMediaPhotoBuilder)
	b.Init(a.receiveMedia)
	return b
}
func (a *EditMessageMediaArgs) MediaVideo() MediaVideoBuilder {
	b := new(implMediaVideoBuilder)
	b.Init(a.receiveMedia)
	return b
}
func (a *EditMessageMediaArgs) MediaAnimation() MediaAnimationBuilder {
	b := new(implMediaAnimationBuilder)
	b.Init(a.receiveMedia)
	return b
}
func (a *EditMessageMediaArgs) MediaAudio() MediaAudioBuilder {
	b := new(implMediaAudioBuilder)
	b.Init(a.receiveMedia)
	return b
}
func (a *EditMessageMediaArgs) MediaDocument() MediaDocumentBuilder {
	b := new(implMediaDocumentBuilder)
	b.Init(a.receiveMedia)
	return b
}

type EditMessageLiveLocationArgs struct {
	_EditMessageReplyMarkupArgs
}

func (a *EditMessageLiveLocationArgs) Location(latitude, longitude float64) {
	a.getForm().
		WithFloat("latitude", latitude).
		WithFloat("longitude", longitude)
}

type StopPollArgs struct {
	_ChatMessageArgs
}

func (a *StopPollArgs) InlineKeyboard() InlineKeyboardBuilder {
	return &implInlineKeyboardBuilder{
		form: a.getForm(),
		name: "reply_markup",
	}
}
