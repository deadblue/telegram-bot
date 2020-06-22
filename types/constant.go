package types

type ParseMode string

const (
	ModePlain      ParseMode = ""
	ModeHtml       ParseMode = "HTML"
	ModeMarkdown   ParseMode = "Markdown"
	ModeMarkdownV2 ParseMode = "MarkdownV2"
)

type InputMediaType string

const (
	MediaAnimation InputMediaType = "animation"
	MediaAudio     InputMediaType = "audio"
	MediaDocument  InputMediaType = "document"
	MediaPhoto     InputMediaType = "photo"
	MediaVideo     InputMediaType = "video"
)

type InlineQueryResultType string

const (
	ResultArticle  InlineQueryResultType = "article"
	ResultPhoto    InlineQueryResultType = "photo"
	ResultGif      InlineQueryResultType = "gif"
	ResultMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	ResultVideo    InlineQueryResultType = "video"
	ResultAudio    InlineQueryResultType = "audio"
	ResultVoice    InlineQueryResultType = "voice"
	ResultDocument InlineQueryResultType = "document"
	ResultLocation InlineQueryResultType = "location"
	ResultVenue    InlineQueryResultType = "venue"
	ResultContact  InlineQueryResultType = "contact"
	ResultGame     InlineQueryResultType = "game"
)
