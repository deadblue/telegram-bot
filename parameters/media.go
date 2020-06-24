package parameters

import "github.com/deadblue/telegroid/types"

// Interface for all InputMedia
// Reference: https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	Type() types.InputMediaType
}

// Information for InputMediaPhoto
// Reference: https://core.telegram.org/bots/api#inputmediaphoto
type InputPhoto struct {
	caption FormattedText
	media   *InputFile
}

func (i *InputPhoto) Type() types.InputMediaType {
	return types.MediaPhoto
}
func (i *InputPhoto) Media(file *InputFile) {
	i.media = file
}
func (i *InputPhoto) Caption(text FormattedText) {
	i.caption = text
}

// Information for InputMediaDocument
// Reference: https://core.telegram.org/bots/api#inputmediadocument
type InputDocument struct {
	InputPhoto
	thumb *InputFile
}

func (i *InputDocument) Type() types.InputMediaType {
	return types.MediaDocument
}
func (i *InputDocument) Thumb(file *InputFile) {
	i.thumb = file
}

// Information for InputMediaAnimation
// Reference: https://core.telegram.org/bots/api#inputmediaanimation
type InputAnimation struct {
	InputDocument
	width    int
	height   int
	duration int
}

func (i *InputAnimation) Type() types.InputMediaType {
	return types.MediaAnimation
}
func (i *InputAnimation) Dimension(width, height int) {
	i.width, i.height = width, height
}
func (i *InputAnimation) Duration(duration int) {
	i.duration = duration
}

// Information for InputMediaVideo
// Reference: https://core.telegram.org/bots/api#inputmediavideo
type InputVideo struct {
	InputAnimation
	supportsStreaming bool
}

func (i *InputVideo) Type() types.InputMediaType {
	return types.MediaVideo
}
func (i *InputVideo) SupportsStreaming() {
	i.supportsStreaming = true
}

// Information for InputMediaAudio
// Reference: https://core.telegram.org/bots/api#inputmediaaudio
type InputAudio struct {
	InputDocument
	duration  int
	title     string
	performer string
}

func (i *InputAudio) Type() types.InputMediaType {
	return types.MediaAudio
}
func (i *InputAudio) Duration(duration int) {
	i.duration = duration
}
func (i *InputAudio) Title(title string) {
	i.title = title
}
func (i *InputAudio) Performer(performer string) {
	i.performer = performer
}
