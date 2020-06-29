package parameters

import (
	"github.com/deadblue/telegram-bot/types"
)

// Interface for all InputMedia
// Reference: https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	Type() types.InputMediaType
}

// Information for InputMediaPhoto
// Reference: https://core.telegram.org/bots/api#inputmediaphoto
type InputPhoto struct {
	media   InputFile
	caption FormattedText
}

func (i *InputPhoto) Type() types.InputMediaType {
	return types.MediaPhoto
}
func (i *InputPhoto) Media(file InputFile) *InputPhoto {
	i.media = file
	return i
}
func (i *InputPhoto) Caption(text FormattedText) *InputPhoto {
	i.caption = text
	return i
}

// Information for InputMediaDocument
// Reference: https://core.telegram.org/bots/api#inputmediadocument
type InputDocument struct {
	media   InputFile
	caption FormattedText
	thumb   InputFile
}

func (i *InputDocument) Type() types.InputMediaType {
	return types.MediaDocument
}
func (i *InputDocument) Media(file InputFile) *InputDocument {
	i.media = file
	return i
}
func (i *InputDocument) Caption(text FormattedText) *InputDocument {
	i.caption = text
	return i
}
func (i *InputDocument) Thumb(file InputFile) *InputDocument {
	i.thumb = file
	return i
}

// Information for InputMediaAnimation
// Reference: https://core.telegram.org/bots/api#inputmediaanimation
type InputAnimation struct {
	media    InputFile
	caption  FormattedText
	thumb    InputFile
	width    int
	height   int
	duration int
}

func (i *InputAnimation) Type() types.InputMediaType {
	return types.MediaAnimation
}
func (i *InputAnimation) Media(file InputFile) *InputAnimation {
	i.media = file
	return i
}
func (i *InputAnimation) Caption(text FormattedText) *InputAnimation {
	i.caption = text
	return i
}
func (i *InputAnimation) Thumb(file InputFile) *InputAnimation {
	i.thumb = file
	return i
}
func (i *InputAnimation) Dimension(width, height int) *InputAnimation {
	i.width, i.height = width, height
	return i
}
func (i *InputAnimation) Duration(duration int) *InputAnimation {
	i.duration = duration
	return i
}

// Information for InputMediaVideo
// Reference: https://core.telegram.org/bots/api#inputmediavideo
type InputVideo struct {
	media             InputFile
	caption           FormattedText
	thumb             InputFile
	width             int
	height            int
	duration          int
	supportsStreaming bool
}

func (i *InputVideo) Type() types.InputMediaType {
	return types.MediaVideo
}
func (i *InputVideo) Media(file InputFile) *InputVideo {
	i.media = file
	return i
}
func (i *InputVideo) Caption(text FormattedText) *InputVideo {
	i.caption = text
	return i
}
func (i *InputVideo) Thumb(file InputFile) *InputVideo {
	i.thumb = file
	return i
}
func (i *InputVideo) Dimension(width, height int) *InputVideo {
	i.width, i.height = width, height
	return i
}
func (i *InputVideo) Duration(duration int) *InputVideo {
	i.duration = duration
	return i
}
func (i *InputVideo) SupportsStreaming() *InputVideo {
	i.supportsStreaming = true
	return i
}

// Information for InputMediaAudio
// Reference: https://core.telegram.org/bots/api#inputmediaaudio
type InputAudio struct {
	media     InputFile
	caption   FormattedText
	thumb     InputFile
	duration  int
	title     string
	performer string
}

func (i *InputAudio) Type() types.InputMediaType {
	return types.MediaAudio
}
func (i *InputAudio) Media(file InputFile) *InputAudio {
	i.media = file
	return i
}
func (i *InputAudio) Caption(text FormattedText) *InputAudio {
	i.caption = text
	return i
}
func (i *InputAudio) Thumb(file InputFile) *InputAudio {
	i.thumb = file
	return i
}
func (i *InputAudio) Duration(duration int) *InputAudio {
	i.duration = duration
	return i
}
func (i *InputAudio) Title(title string) *InputAudio {
	i.title = title
	return i
}
func (i *InputAudio) Performer(performer string) *InputAudio {
	i.performer = performer
	return i
}
