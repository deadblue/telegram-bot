package types

type InlineQuery struct {
	Id       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageId string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

type InlineQueryResultBase struct {
	Type                InlineQueryResultType `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Description         string                `json:"description"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
}

type InlineQueryResultArticle struct {
	InlineQueryResultBase
	Url         string `json:"url"`
	HideUrl     bool   `json:"hide_url"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InlineQueryResultPhoto struct {
	InlineQueryResultBase
	PhotoUrl    string `json:"photo_url"`
	PhotoWidth  int    `json:"photo_width"`
	PhotoHeight int    `json:"photo_height"`
	ThumbUrl    string `json:"thumb_url"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultGif struct {
	InlineQueryResultBase
	GifUrl      string `json:"gif_url"`
	GifWidth    int    `json:"gif_width"`
	GifHeight   int    `json:"gif_height"`
	GifDuration int    `json:"gif_duration"`
	ThumbUrl    string `json:"thumb_url"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultMpeg4Gif struct {
	InlineQueryResultBase
	Mpeg4Url      string `json:"mpeg4_url"`
	Mpeg4Width    int    `json:"mpeg4_width"`
	Mpeg4Height   int    `json:"mpeg4_height"`
	Mpeg4Duration int    `json:"mpeg4_duration"`
	ThumbUrl      string `json:"thumb_url"`
	Caption       string `json:"caption"`
	ParseMode     string `json:"parse_mode"`
}

type InlineQueryResultVideo struct {
	InlineQueryResultBase
	VideoUrl      string `json:"video_url"`
	VideoWidth    int    `json:"video_width"`
	VideoHeight   int    `json:"video_height"`
	VideoDuration int    `json:"video_duration"`
	MimeType      string `json:"mime_type"`
	ThumbUrl      string `json:"thumb_url"`
	Caption       string `json:"caption"`
	ParseMode     string `json:"parse_mode"`
}

type InlineQueryResultAudio struct {
	InlineQueryResultBase
	AudioUrl      string `json:"audio_url"`
	AudioDuration int    `json:"audio_duration"`
	Performer     string `json:"performer"`
	Caption       string `json:"caption"`
	ParseMode     string `json:"parse_mode"`
}

type InlineQueryResultVoice struct {
	InlineQueryResultBase
	VoiceUrl      string `json:"voice_url"`
	VoiceDuration string `json:"voice_duration"`
	Caption       string `json:"caption"`
	ParseMode     string `json:"parse_mode"`
}

type InlineQueryResultDocument struct {
	InlineQueryResultBase
	DocumentUrl string `json:"document_url"`
	MimeType    string `json:"mime_type"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InlineQueryResultLocation struct {
	InlineQueryResultBase
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	LivePeriod  int     `json:"live_period"`
	ThumbUrl    string  `json:"thumb_url"`
	ThumbWidth  int     `json:"thumb_width"`
	ThumbHeight int     `json:"thumb_height"`
}

type InlineQueryResultVenue struct {
	InlineQueryResultBase
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Address        string  `json:"address"`
	FoursquareId   string  `json:"foursquare_id"`
	FoursquareType string  `json:"foursquare_type"`
	ThumbUrl       string  `json:"thumb_url"`
	ThumbWidth     int     `json:"thumb_width"`
	ThumbHeight    int     `json:"thumb_height"`
}

type InlineQueryResultContact struct {
	InlineQueryResultBase
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InlineQueryResultCachedPhoto struct {
	InlineQueryResultBase
	PhotoFileId string `json:"photo_file_id"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultCachedGif struct {
	InlineQueryResultBase
	GifFileId string `json:"gif_file_id"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	InlineQueryResultBase
	Mpeg4FileId string `json:"mpeg4_file_id"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultCachedSticker struct {
	InlineQueryResultBase
	StickerFileId string `json:"sticker_file_id"`
}

type InlineQueryResultCachedDocument struct {
	InlineQueryResultBase
	DocumentFileId string `json:"document_file_id"`
	Caption        string `json:"caption"`
	ParseMode      string `json:"parse_mode"`
}

type InlineQueryResultCachedVideo struct {
	InlineQueryResultBase
	VideoFileId string `json:"video_file_id"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultCachedVoice struct {
	InlineQueryResultBase
	VoiceFileId string `json:"voice_file_id"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultCachedAudio struct {
	InlineQueryResultBase
	AudioFileId string `json:"audio_file_id"`
	Caption     string `json:"caption"`
	ParseMode   string `json:"parse_mode"`
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	Id            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup"`
}

type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LivePeriod int     `json:"live_period"`
}

type InputVenueMessageContent struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Title          string  `json:"title"`
	Address        string  `json:"address"`
	FoursquareId   string  `json:"foursquare_id"`
	FoursquareType string  `json:"foursquare_type"`
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
}

type InputMessageContent interface {
	IsInputMessageContent() bool
}

func (c *InputTextMessageContent) IsInputMessageContent() bool {
	return true
}

func (c *InputLocationMessageContent) IsInputMessageContent() bool {
	return true
}

func (c *InputVenueMessageContent) IsInputMessageContent() bool {
	return true
}

func (c *InputContactMessageContent) IsInputMessageContent() bool {
	return true
}
