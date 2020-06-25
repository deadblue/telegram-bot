package types

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type Video struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     int        `json:"file_size"`
}

type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

type InputMediaAnimation struct {
	Type      InputMediaType `json:"type"`
	Media     string         `json:"media"`
	Thumb     string         `json:"thumb,omitempty"`
	Caption   string         `json:"caption,omitempty"`
	ParseMode ParseMode      `json:"parse_mode,omitempty"`
	Width     int            `json:"width,omitempty"`
	Height    int            `json:"height,omitempty"`
	Duration  int            `json:"duration,omitempty"`
}

type InputMediaAudio struct {
	Type      InputMediaType `json:"type"`
	Media     string         `json:"media"`
	Thumb     string         `json:"thumb,omitempty"`
	Caption   string         `json:"caption,omitempty"`
	ParseMode ParseMode      `json:"parse_mode,omitempty"`
	Duration  int            `json:"duration,omitempty"`
	Performer string         `json:"performer,omitempty"`
	Title     string         `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type      InputMediaType `json:"type"`
	Media     string         `json:"media"`
	Thumb     string         `json:"thumb,omitempty"`
	Caption   string         `json:"caption,omitempty"`
	ParseMode ParseMode      `json:"parse_mode,omitempty"`
}

type InputMediaPhoto struct {
	Type      InputMediaType `json:"type"`
	Media     string         `json:"media"`
	Caption   string         `json:"caption,omitempty"`
	ParseMode ParseMode      `json:"parse_mode,omitempty"`
}

type InputMediaVideo struct {
	Type              InputMediaType `json:"type"`
	Media             string         `json:"media"`
	Thumb             string         `json:"thumb,omitempty"`
	Caption           string         `json:"caption,omitempty"`
	ParseMode         ParseMode      `json:"parse_mode,omitempty"`
	Width             int            `json:"width,omitempty"`
	Height            int            `json:"height,omitempty"`
	Duration          int            `json:"duration,omitempty"`
	SupportsStreaming bool           `json:"supports_streaming,omitempty"`
}
