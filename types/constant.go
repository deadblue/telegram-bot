package types

type ParseMode string

const (
	ModeHtml       ParseMode = "HTML"
	ModeMarkdown   ParseMode = "Markdown"
	ModeMarkdownV2 ParseMode = "MarkdownV2"
)

type InlineQueryResultType string

const (
	TypeArticle  InlineQueryResultType = "article"
	TypePhoto    InlineQueryResultType = "photo"
	TypeGif      InlineQueryResultType = "gif"
	TypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	TypeVideo    InlineQueryResultType = "video"
	TypeAudio    InlineQueryResultType = "audio"
	TypeVoice    InlineQueryResultType = "voice"
	TypeDocument InlineQueryResultType = "document"
	TypeLocation InlineQueryResultType = "location"
	TypeVenue    InlineQueryResultType = "venue"
	TypeContact  InlineQueryResultType = "contact"
	TypeGame     InlineQueryResultType = "game"
)
