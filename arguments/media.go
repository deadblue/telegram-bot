package arguments

import (
	"fmt"
	"io"
)

// TODO: implement this...

type MediaPhotoBuilder interface {

	// Set caption
	Caption(caption string) MediaPhotoBuilder

	// Set caption parsing mode to Markdown
	Markdown() MediaPhotoBuilder

	// Set caption parsing mode to HTML
	HTML() MediaPhotoBuilder

	// Set media with URL
	MediaUrl(url string) MediaPhotoBuilder

	// Set media with FileId
	MediaFileId(fileId string) MediaPhotoBuilder

	// Set media with file
	MediaFile(file InputFile) MediaPhotoBuilder

	// Apply the settings
	Finish()

}

type MediaVideoBuilder interface {

	// Set caption
	Caption(caption string) MediaVideoBuilder

	// Set caption parsing mode to Markdown
	Markdown() MediaVideoBuilder

	// Set caption parsing mode to HTML
	HTML() MediaVideoBuilder

	// Set video with URL
	MediaUrl(url string) MediaVideoBuilder

	// Set video with existing fileId
	MediaFileId(fileId string) MediaVideoBuilder

	// Set video with file
	MediaFile(file InputFile) MediaVideoBuilder

	// Set thumb with existing fileId
	ThumbFileId(fileId string) MediaVideoBuilder

	// Set thumb with file
	ThumbFile(file InputFile) MediaVideoBuilder

	// Set video size
	Size(width, height int) MediaVideoBuilder

	// Set video duration in seconds
	Duration(duration int) MediaVideoBuilder

	// Set video supports streaming
	SupportsStreaming() MediaVideoBuilder

	// Apply the settings
	Finish()

}

type MediaAnimationBuilder interface {
	Caption(caption string) MediaAnimationBuilder
	Markdown() MediaAnimationBuilder
	HTML() MediaAnimationBuilder
	MediaUrl(url string) MediaAnimationBuilder
	MediaFileId(fileId string) MediaAnimationBuilder
	MediaFile(file InputFile) MediaAnimationBuilder
	ThumbFileId(fileId string) MediaAnimationBuilder
	ThumbFile(file InputFile) MediaAnimationBuilder
	Size(width, height int) MediaAnimationBuilder
	Duration(duration int) MediaAnimationBuilder
	Finish()
}

type MediaAudioBuilder interface {
	Caption(caption string) MediaAudioBuilder
	Markdown() MediaAudioBuilder
	HTML() MediaAudioBuilder
	MediaUrl(url string) MediaAudioBuilder
	MediaFileId(fileId string) MediaAudioBuilder
	MediaFile(file InputFile) MediaAudioBuilder
	ThumbFileId(fileId string) MediaAudioBuilder
	ThumbFile(file InputFile) MediaAudioBuilder
	Duration(duration int) MediaAudioBuilder
	Performer(performer string) MediaAudioBuilder
	Title(title string) MediaAudioBuilder
	Finish()
}

type MediaDocumentBuilder interface {
	Caption(caption string) MediaDocumentBuilder
	Markdown() MediaDocumentBuilder
	HTML() MediaDocumentBuilder
	MediaUrl(url string) MediaDocumentBuilder
	MediaFileId(fileId string) MediaDocumentBuilder
	MediaFile(file InputFile) MediaDocumentBuilder
	ThumbFileId(fileId string) MediaDocumentBuilder
	ThumbFile(file InputFile) MediaDocumentBuilder
	Finish()
}


type implMediaPhotoBuilder struct {
	receiver func(map[string]interface{})
	data map[string]interface{}
}
func (b *implMediaPhotoBuilder) getData() map[string]interface{} {
	if b.data ==  nil {
		b.data = make(map[string]interface{})
	}
	return b.data
}
func (b *implMediaPhotoBuilder) Caption(caption string) MediaPhotoBuilder {
	b.getData()["caption"] = caption
	return b
}
func (b *implMediaPhotoBuilder) Markdown() MediaPhotoBuilder {
	b.getData()["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaPhotoBuilder) HTML() MediaPhotoBuilder {
	b.getData()["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaPhotoBuilder) MediaUrl(url string) MediaPhotoBuilder {
	b.getData()["media"] = url
	return b
}
func (b *implMediaPhotoBuilder) MediaFileId(fileId string) MediaPhotoBuilder {
	b.getData()["media"] = fileId
	return b
}
func (b *implMediaPhotoBuilder) MediaFile(file InputFile) MediaPhotoBuilder {
	b.getData()["media"] = file
	return b
}
func (b *implMediaPhotoBuilder) Finish() {
	data := b.getData()
	data["type"] = mediaPhoto
	b.receiver(data)
}


type implMediaVideoBuilder struct {
	receiver func(map[string]interface{})
	data map[string]interface{}
}
func (b *implMediaVideoBuilder) getData() map[string]interface{} {
	if b.data == nil {
		b.data = make(map[string]interface{})
	}
	return b.data
}
func (b *implMediaVideoBuilder) Caption(caption string) MediaVideoBuilder {
	b.getData()["caption"] = caption
	return b
}
func (b *implMediaVideoBuilder) Markdown() MediaVideoBuilder {
	b.getData()["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaVideoBuilder) HTML() MediaVideoBuilder {
	b.getData()["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaVideoBuilder) MediaUrl(url string) MediaVideoBuilder {
	b.getData()["media"] = url
	return b
}
func (b *implMediaVideoBuilder) MediaFileId(fileId string) MediaVideoBuilder {
	b.getData()["media"] = fileId
	return b
}
func (b *implMediaVideoBuilder) MediaFile(file InputFile) MediaVideoBuilder {
	b.getData()["media"] = file
	return b
}
func (b *implMediaVideoBuilder) ThumbFileId(fileId string) MediaVideoBuilder {
	b.getData()["thumb"] = fileId
	return b
}
func (b *implMediaVideoBuilder) ThumbFile(file InputFile) MediaVideoBuilder {
	b.getData()["thumb"] = file
	return b
}
func (b *implMediaVideoBuilder) Size(width, height int) MediaVideoBuilder {
	data := b.getData()
	data["width"] = width
	data["height"] = height
	return b
}
func (b *implMediaVideoBuilder) Duration(duration int) MediaVideoBuilder {
	b.getData()["duration"] = duration
	return b
}
func (b *implMediaVideoBuilder) SupportsStreaming() MediaVideoBuilder {
	b.getData()["supports_streaming"] = true
	return b
}
func (b *implMediaVideoBuilder) Finish() {
	data := b.getData()
	data["type"] = mediaVideo
	b.receiver(data)
}


type SendMediaGroupArgs struct {
	ChatArgs

	mediaBuf []map[string]interface{}
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
func (a *SendMediaGroupArgs) MediaPhoto() MediaPhotoBuilder {
	return &implMediaPhotoBuilder{ receiver:a.receiveMedia }
}
func (a *SendMediaGroupArgs) MediaVideo() MediaVideoBuilder {
	return &implMediaVideoBuilder{ receiver:a.receiveMedia }
}
func (a *SendMediaGroupArgs) ReplyToMessageId(messageId int) {
	a.getForm().WithInt("reply_to_message_id", messageId)
}
func (a *SendMediaGroupArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}
func (a *SendMediaGroupArgs) Finish() (contentType string, data io.Reader) {
	form := a.getForm()
	for i := 0; i < a.mediaCount; i ++ {
		item := a.mediaBuf[i]
		// Uploading media file
		mf, ok := item["media"].(InputFile)
		if ok {
			attachName := fmt.Sprintf("attach_media_%d", i)
			item["media"] = fmt.Sprintf("attach://%s", attachName)
			form.WithFile(attachName, mf)
		}
		// Uploading thumb file
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


type EditMessageMediaArgs struct {
	EditMessageReplyMarkupArgs
}
func (a *EditMessageMediaArgs) receiveMedia(media map[string]interface{}) {
	a.getForm().WithJson("media", media)
}
func (a *EditMessageMediaArgs) MediaPhoto() MediaPhotoBuilder {
	return &implMediaPhotoBuilder{ receiver: a.receiveMedia }
}
func (a *EditMessageMediaArgs) MediaVideo() MediaVideoBuilder {
	return &implMediaVideoBuilder{ receiver:a.receiveMedia }
}
func (a *EditMessageMediaArgs) MediaAnimation() MediaAnimationBuilder {
	return nil
}
func (a *EditMessageMediaArgs) MediaAudio() MediaAudioBuilder {
	return nil
}
func (a *EditMessageMediaArgs) MediaDocument() MediaDocumentBuilder {
	return nil
}
