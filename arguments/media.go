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

// The base struct for all media builders
type basicMediaBuilder struct {
	data     _MapValue
	receiver func(_MapValue)
}

func (b *basicMediaBuilder) set(name string, value interface{}) {
	if b.data == nil {
		b.data = make(_MapValue)
	}
	b.data[name] = value
}
func (b *basicMediaBuilder) setType(value string) {
	b.set("type", value)
}
func (b *basicMediaBuilder) setCaption(value string) {
	b.set("caption", value)
}
func (b *basicMediaBuilder) setMarkdown() {
	b.set("parse_mode", parseModeMarkdown)
}
func (b *basicMediaBuilder) setHTML() {
	b.set("parse_mode", parseModeHTML)
}
func (b *basicMediaBuilder) setMedia(value interface{}) {
	b.set("media", value)
}
func (b *basicMediaBuilder) setThumb(value interface{}) {
	b.set("thumb", value)
}
func (b *basicMediaBuilder) setDuration(value int) {
	b.set("duration", value)
}
func (b *basicMediaBuilder) Finish() {
	b.receiver(b.data)
}

// The implementation for "MediaPhotoBuilder"
type implMediaPhotoBuilder struct {
	basicMediaBuilder
}

func (b *implMediaPhotoBuilder) Init(receiver func(_MapValue)) MediaPhotoBuilder {
	b.receiver = receiver
	b.setType(mediaPhoto)
	return b
}
func (b *implMediaPhotoBuilder) Caption(caption string) MediaPhotoBuilder {
	b.setCaption(caption)
	return b
}
func (b *implMediaPhotoBuilder) Markdown() MediaPhotoBuilder {
	b.setMarkdown()
	return b
}
func (b *implMediaPhotoBuilder) HTML() MediaPhotoBuilder {
	b.setHTML()
	return b
}
func (b *implMediaPhotoBuilder) MediaUrl(url string) MediaPhotoBuilder {
	b.setMedia(url)
	return b
}
func (b *implMediaPhotoBuilder) MediaFileId(fileId string) MediaPhotoBuilder {
	b.setMedia(fileId)
	return b
}
func (b *implMediaPhotoBuilder) MediaFile(file InputFile) MediaPhotoBuilder {
	b.setMedia(file)
	return b
}

// The implementation for "MediaVideoBuilder"
type implMediaVideoBuilder struct {
	basicMediaBuilder
}

func (b *implMediaVideoBuilder) Init(receiver func(_MapValue)) MediaVideoBuilder {
	b.receiver = receiver
	b.setType(mediaVideo)
	return b
}
func (b *implMediaVideoBuilder) Caption(caption string) MediaVideoBuilder {
	b.setCaption(caption)
	return b
}
func (b *implMediaVideoBuilder) Markdown() MediaVideoBuilder {
	b.setMarkdown()
	return b
}
func (b *implMediaVideoBuilder) HTML() MediaVideoBuilder {
	b.setHTML()
	return b
}
func (b *implMediaVideoBuilder) MediaUrl(url string) MediaVideoBuilder {
	b.setMedia(url)
	return b
}
func (b *implMediaVideoBuilder) MediaFileId(fileId string) MediaVideoBuilder {
	b.setMedia(fileId)
	return b
}
func (b *implMediaVideoBuilder) MediaFile(file InputFile) MediaVideoBuilder {
	b.setMedia(file)
	return b
}
func (b *implMediaVideoBuilder) ThumbFileId(fileId string) MediaVideoBuilder {
	b.setThumb(fileId)
	return b
}
func (b *implMediaVideoBuilder) ThumbFile(file InputFile) MediaVideoBuilder {
	b.setThumb(file)
	return b
}
func (b *implMediaVideoBuilder) Duration(duration int) MediaVideoBuilder {
	b.setDuration(duration)
	return b
}
func (b *implMediaVideoBuilder) Size(width, height int) MediaVideoBuilder {
	b.set("width", width)
	b.set("height", height)
	return b
}
func (b *implMediaVideoBuilder) SupportsStreaming() MediaVideoBuilder {
	b.set("supports_streaming", true)
	return b
}

// The implementation for "MediaAnimationBuilder"
type implMediaAnimationBuilder struct {
	basicMediaBuilder
}

func (b *implMediaAnimationBuilder) Init(receiver func(_MapValue)) MediaAnimationBuilder {
	b.receiver = receiver
	b.setType(mediaAnimation)
	return b
}
func (b *implMediaAnimationBuilder) Caption(caption string) MediaAnimationBuilder {
	b.setCaption(caption)
	return b
}
func (b *implMediaAnimationBuilder) Markdown() MediaAnimationBuilder {
	b.setMarkdown()
	return b
}
func (b *implMediaAnimationBuilder) HTML() MediaAnimationBuilder {
	b.setHTML()
	return b
}
func (b *implMediaAnimationBuilder) MediaUrl(url string) MediaAnimationBuilder {
	b.setMedia(url)
	return b
}
func (b *implMediaAnimationBuilder) MediaFileId(fileId string) MediaAnimationBuilder {
	b.setMedia(fileId)
	return b
}
func (b *implMediaAnimationBuilder) MediaFile(file InputFile) MediaAnimationBuilder {
	b.setMedia(file)
	return b
}
func (b *implMediaAnimationBuilder) ThumbFileId(fileId string) MediaAnimationBuilder {
	b.setThumb(fileId)
	return b
}
func (b *implMediaAnimationBuilder) ThumbFile(file InputFile) MediaAnimationBuilder {
	b.setThumb(file)
	return b
}
func (b *implMediaAnimationBuilder) Duration(duration int) MediaAnimationBuilder {
	b.setDuration(duration)
	return b
}
func (b *implMediaAnimationBuilder) Size(width, height int) MediaAnimationBuilder {
	b.set("width", width)
	b.set("height", height)
	return b
}

// The implementation for "MediaAudioBuilder"
type implMediaAudioBuilder struct {
	basicMediaBuilder
}

func (b *implMediaAudioBuilder) Init(receiver func(_MapValue)) MediaAudioBuilder {
	b.receiver = receiver
	b.setType(mediaAudio)
	return b
}
func (b *implMediaAudioBuilder) Caption(caption string) MediaAudioBuilder {
	b.setCaption(caption)
	return b
}
func (b *implMediaAudioBuilder) Markdown() MediaAudioBuilder {
	b.setMarkdown()
	return b
}
func (b *implMediaAudioBuilder) HTML() MediaAudioBuilder {
	b.setHTML()
	return b
}
func (b *implMediaAudioBuilder) MediaUrl(url string) MediaAudioBuilder {
	b.setMedia(url)
	return b
}
func (b *implMediaAudioBuilder) MediaFileId(fileId string) MediaAudioBuilder {
	b.setMedia(fileId)
	return b
}
func (b *implMediaAudioBuilder) MediaFile(file InputFile) MediaAudioBuilder {
	b.setMedia(file)
	return b
}
func (b *implMediaAudioBuilder) ThumbFileId(fileId string) MediaAudioBuilder {
	b.setThumb(fileId)
	return b
}
func (b *implMediaAudioBuilder) ThumbFile(file InputFile) MediaAudioBuilder {
	b.setThumb(file)
	return b
}
func (b *implMediaAudioBuilder) Duration(duration int) MediaAudioBuilder {
	b.setDuration(duration)
	return b
}
func (b *implMediaAudioBuilder) Performer(performer string) MediaAudioBuilder {
	b.set("performer", performer)
	return b
}
func (b *implMediaAudioBuilder) Title(title string) MediaAudioBuilder {
	b.set("title", title)
	return b
}

// The implementation for "MediaDocumentBuilder"
type implMediaDocumentBuilder struct {
	basicMediaBuilder
}

func (b *implMediaDocumentBuilder) Init(receiver func(_MapValue)) MediaDocumentBuilder {
	b.receiver = receiver
	b.setType(mediaDocument)
	return b
}
func (b *implMediaDocumentBuilder) Caption(caption string) MediaDocumentBuilder {
	b.setCaption(caption)
	return b
}
func (b *implMediaDocumentBuilder) Markdown() MediaDocumentBuilder {
	b.setMarkdown()
	return b
}
func (b *implMediaDocumentBuilder) HTML() MediaDocumentBuilder {
	b.setHTML()
	return b
}
func (b *implMediaDocumentBuilder) MediaUrl(url string) MediaDocumentBuilder {
	b.setMedia(url)
	return b
}
func (b *implMediaDocumentBuilder) MediaFileId(fileId string) MediaDocumentBuilder {
	b.setMedia(fileId)
	return b
}
func (b *implMediaDocumentBuilder) MediaFile(file InputFile) MediaDocumentBuilder {
	b.setMedia(file)
	return b
}
func (b *implMediaDocumentBuilder) ThumbFileId(fileId string) MediaDocumentBuilder {
	b.setThumb(fileId)
	return b
}
func (b *implMediaDocumentBuilder) ThumbFile(file InputFile) MediaDocumentBuilder {
	b.setThumb(file)
	return b
}
