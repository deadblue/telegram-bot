package types

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
