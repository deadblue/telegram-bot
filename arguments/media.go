package arguments

type MediaPhotoBuilder interface {

	ArgumentBuilder

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

	// Set media with file data
	MediaFile(file InputFile) MediaPhotoBuilder

}

type MediaVideoBuilder interface {

	ArgumentBuilder

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

	// Set video with file data
	MediaFile(file InputFile) MediaVideoBuilder

	// Set thumb with existing fileId
	ThumbFileId(fileId string) MediaVideoBuilder

	// Set thumb with file data
	ThumbFile(file InputFile) MediaVideoBuilder

	// Set video size
	Size(width, height int) MediaVideoBuilder

	// Set video duration in seconds
	Duration(duration int) MediaVideoBuilder

	// Set video supports streaming
	SupportsStreaming() MediaVideoBuilder

}

type MediaAnimationBuilder interface {

	ArgumentBuilder

	// Set caption
	Caption(caption string) MediaAnimationBuilder

	// Set caption parsing mode to Markdown
	Markdown() MediaAnimationBuilder

	// Set caption parsing mode to HTML
	HTML() MediaAnimationBuilder

	// Set animation with URL
	MediaUrl(url string) MediaAnimationBuilder

	// Set animation with existing fileId
	MediaFileId(fileId string) MediaAnimationBuilder

	// Set animation with file data
	MediaFile(file InputFile) MediaAnimationBuilder

	// Set thumb with existing fileId
	ThumbFileId(fileId string) MediaAnimationBuilder

	// Set thumb with file data
	ThumbFile(file InputFile) MediaAnimationBuilder

	// Set animation size
	Size(width, height int) MediaAnimationBuilder

	// Set animation duration in seconds
	Duration(duration int) MediaAnimationBuilder

}

type MediaAudioBuilder interface {

	ArgumentBuilder

	// Set caption
	Caption(caption string) MediaAudioBuilder

	// Set caption parsing mode to Markdown
	Markdown() MediaAudioBuilder

	// Set caption parsing mode to HTML
	HTML() MediaAudioBuilder

	// Set audio with URL
	MediaUrl(url string) MediaAudioBuilder

	// Set audio with existing fileId
	MediaFileId(fileId string) MediaAudioBuilder

	// Set audio with file data
	MediaFile(file InputFile) MediaAudioBuilder

	// Set thumb with existing fileId
	ThumbFileId(fileId string) MediaAudioBuilder

	// Set thumb with file data
	ThumbFile(file InputFile) MediaAudioBuilder

	// Set audio duration in seconds
	Duration(duration int) MediaAudioBuilder

	// Set audio performer
	Performer(performer string) MediaAudioBuilder

	// Set audio title
	Title(title string) MediaAudioBuilder

}

type MediaDocumentBuilder interface {

	ArgumentBuilder

	// Set caption
	Caption(caption string) MediaDocumentBuilder

	// Set caption parsing mode to Markdown
	Markdown() MediaDocumentBuilder

	// Set caption parsing mode to HTML
	HTML() MediaDocumentBuilder

	// Set document with URL
	MediaUrl(url string) MediaDocumentBuilder

	// Set video with existing fileId
	MediaFileId(fileId string) MediaDocumentBuilder

	// Set video with file data
	MediaFile(file InputFile) MediaDocumentBuilder

	// Set thumb with existing fileId
	ThumbFileId(fileId string) MediaDocumentBuilder

	// Set thumb with file
	ThumbFile(file InputFile) MediaDocumentBuilder

}


type basicMediaBuilder struct {
	receiver func(map[string]interface{})
	data map[string]interface{}
}
func (b *basicMediaBuilder) Init(receiver func(map[string]interface{})) {
	b.receiver = receiver
	b.data = make(map[string]interface{})
}


type implMediaPhotoBuilder struct {
	basicMediaBuilder
}
func (b *implMediaPhotoBuilder) Caption(caption string) MediaPhotoBuilder {
	b.data["caption"] = caption
	return b
}
func (b *implMediaPhotoBuilder) Markdown() MediaPhotoBuilder {
	b.data["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaPhotoBuilder) HTML() MediaPhotoBuilder {
	b.data["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaPhotoBuilder) MediaUrl(url string) MediaPhotoBuilder {
	b.data["media"] = url
	return b
}
func (b *implMediaPhotoBuilder) MediaFileId(fileId string) MediaPhotoBuilder {
	b.data["media"] = fileId
	return b
}
func (b *implMediaPhotoBuilder) MediaFile(file InputFile) MediaPhotoBuilder {
	b.data["media"] = file
	return b
}
func (b *implMediaPhotoBuilder) Finish() {
	b.data["type"] = mediaPhoto
	b.receiver(b.data)
}


type implMediaVideoBuilder struct {
	basicMediaBuilder
}
func (b *implMediaVideoBuilder) Caption(caption string) MediaVideoBuilder {
	b.data["caption"] = caption
	return b
}
func (b *implMediaVideoBuilder) Markdown() MediaVideoBuilder {
	b.data["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaVideoBuilder) HTML() MediaVideoBuilder {
	b.data["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaVideoBuilder) MediaUrl(url string) MediaVideoBuilder {
	b.data["media"] = url
	return b
}
func (b *implMediaVideoBuilder) MediaFileId(fileId string) MediaVideoBuilder {
	b.data["media"] = fileId
	return b
}
func (b *implMediaVideoBuilder) MediaFile(file InputFile) MediaVideoBuilder {
	b.data["media"] = file
	return b
}
func (b *implMediaVideoBuilder) ThumbFileId(fileId string) MediaVideoBuilder {
	b.data["thumb"] = fileId
	return b
}
func (b *implMediaVideoBuilder) ThumbFile(file InputFile) MediaVideoBuilder {
	b.data["thumb"] = file
	return b
}
func (b *implMediaVideoBuilder) Size(width, height int) MediaVideoBuilder {
	b.data["width"] = width
	b.data["height"] = height
	return b
}
func (b *implMediaVideoBuilder) Duration(duration int) MediaVideoBuilder {
	b.data["duration"] = duration
	return b
}
func (b *implMediaVideoBuilder) SupportsStreaming() MediaVideoBuilder {
	b.data["supports_streaming"] = true
	return b
}
func (b *implMediaVideoBuilder) Finish() {
	b.data["type"] = mediaVideo
	b.receiver(b.data)
}


type implMediaAnimationBuilder struct {
	basicMediaBuilder
}
func (b *implMediaAnimationBuilder) Caption(caption string) MediaAnimationBuilder {
	b.data["caption"] = caption
	return b
}
func (b *implMediaAnimationBuilder) Markdown() MediaAnimationBuilder {
	b.data["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaAnimationBuilder) HTML() MediaAnimationBuilder {
	b.data["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaAnimationBuilder) MediaUrl(url string) MediaAnimationBuilder {
	b.data["media"] = url
	return b
}
func (b *implMediaAnimationBuilder) MediaFileId(fileId string) MediaAnimationBuilder {
	b.data["media"] = fileId
	return b
}
func (b *implMediaAnimationBuilder) MediaFile(file InputFile) MediaAnimationBuilder {
	b.data["media"] = file
	return b
}
func (b *implMediaAnimationBuilder) ThumbFileId(fileId string) MediaAnimationBuilder {
	b.data["thumb"] = fileId
	return b
}
func (b *implMediaAnimationBuilder) ThumbFile(file InputFile) MediaAnimationBuilder {
	b.data["thumb"] = file
	return b
}
func (b *implMediaAnimationBuilder) Size(width, height int) MediaAnimationBuilder {
	b.data["width"] = width
	b.data["height"] = height
	return b
}
func (b *implMediaAnimationBuilder) Duration(duration int) MediaAnimationBuilder {
	b.data["duration"] = duration
	return b
}
func (b *implMediaAnimationBuilder) Finish() {
	b.data["type"] = mediaAnimation
	b.receiver(b.data)
}


type implMediaAudioBuilder struct {
	basicMediaBuilder
}
func (b *implMediaAudioBuilder) Caption(caption string) MediaAudioBuilder {
	b.data["caption"] = caption
	return b
}
func (b *implMediaAudioBuilder) Markdown() MediaAudioBuilder {
	b.data["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaAudioBuilder) HTML() MediaAudioBuilder {
	b.data["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaAudioBuilder) MediaUrl(url string) MediaAudioBuilder {
	b.data["media"] = url
	return b
}
func (b *implMediaAudioBuilder) MediaFileId(fileId string) MediaAudioBuilder {
	b.data["media"] = fileId
	return b
}
func (b *implMediaAudioBuilder) MediaFile(file InputFile) MediaAudioBuilder {
	b.data["media"] = file
	return b
}
func (b *implMediaAudioBuilder) ThumbFileId(fileId string) MediaAudioBuilder {
	b.data["thumb"] = fileId
	return b
}
func (b *implMediaAudioBuilder) ThumbFile(file InputFile) MediaAudioBuilder {
	b.data["thumb"] = file
	return b
}
func (b *implMediaAudioBuilder) Duration(duration int) MediaAudioBuilder {
	b.data["duration"] = duration
	return b
}
func (b *implMediaAudioBuilder) Performer(performer string) MediaAudioBuilder {
	b.data["performer"] = performer
	return b
}
func (b *implMediaAudioBuilder) Title(title string) MediaAudioBuilder {
	b.data["title"] = title
	return b
}
func (b *implMediaAudioBuilder) Finish() {
	b.data["type"] = mediaAudio
	b.receiver(b.data)
}


type implMediaDocumentBuilder struct {
	basicMediaBuilder
}
func (b *implMediaDocumentBuilder) Caption(caption string) MediaDocumentBuilder {
	b.data["caption"] = caption
	return b
}
func (b *implMediaDocumentBuilder) Markdown() MediaDocumentBuilder {
	b.data["parse_mode"] = parseModeMarkdown
	return b
}
func (b *implMediaDocumentBuilder) HTML() MediaDocumentBuilder {
	b.data["parse_mode"] = parseModeHTML
	return b
}
func (b *implMediaDocumentBuilder) MediaUrl(url string) MediaDocumentBuilder {
	b.data["media"] = url
	return b
}
func (b *implMediaDocumentBuilder) MediaFileId(fileId string) MediaDocumentBuilder {
	b.data["media"] = fileId
	return b
}
func (b *implMediaDocumentBuilder) MediaFile(file InputFile) MediaDocumentBuilder {
	b.data["media"] = file
	return b
}
func (b *implMediaDocumentBuilder) ThumbFileId(fileId string) MediaDocumentBuilder {
	b.data["thumb"] = fileId
	return b
}
func (b *implMediaDocumentBuilder) ThumbFile(file InputFile) MediaDocumentBuilder {
	b.data["thumb"] = file
	return b
}
func (b *implMediaDocumentBuilder) Finish() {
	b.data["type"] = mediaDocument
	b.receiver(b.data)
}
