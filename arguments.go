package telegroid

type ParseMode string
type ChatAction string

const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML     ParseMode = "HTML"

	ChatActionTyping          ChatAction = "typing"
	ChatActionUploadPhoto     ChatAction = "upload_photo"
	ChatActionRecordVideo     ChatAction = "record_video"
	ChatActionUploadVideo     ChatAction = "upload_video"
	ChatActionRecordAudio     ChatAction = "record_audio"
	ChatActionUploadAudio     ChatAction = "upload_audio"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

type SetWebhookArguments struct {
	Url            string    `parameter:"url"`
	Certificate    InputFile `parameter:"certificate"`
	MaxConnections int       `parameter:"max_connections"`
	AllowedUpdates []string  `parameter:"allowed_updates"`
}

type GetUpdatesArguments struct {
	Offset         int      `parameter:"offset"`
	Limit          int      `parameter:"limit"`
	Timeout        int      `parameter:"timeout"`
	AllowedUpdates []string `parameter:"allowed_updates"`
}

type ForwardMessageArguments struct {
	ChatId              string `parameter:"chat_id"`
	FromChatId          string `parameter:"from_chat_id"`
	DisableNotification bool   `parameter:"disable_notification"`
	MessageId           int    `parameter:"message_id"`
}

type AbstractSendArguments struct {
	ChatId            string      `parameter:"chat_id"`
	DisableNotificate bool        `parameter:"disable_notificate"`
	ReplyToMessageId  int         `parameter:"reply_to_message_id"`
	ReplyMarkup       interface{} `parameter:"reply_markup"`
}

type SendMessageArguments struct {
	AbstractSendArguments
	Text                  string    `parameter:"text"`
	ParseMode             ParseMode `parameter:"parse_mode"`
	DisableWebPagePreview bool      `parameter:"disable_web_page_preview"`
}

type SendPhotoArguments struct {
	AbstractSendArguments
	Photo     InputFile `parameter:"photo"`
	Caption   string    `parameter:"caption"`
	ParseMode ParseMode `parameter:"parse_mode"`
}

type SendAudioArguments struct {
	AbstractSendArguments
	Audio     InputFile `parameter:"audio"`
	Caption   string    `parameter:"caption"`
	ParseMode ParseMode `parameter:"parse_mode"`
	Duration  int       `parameter:"duration"`
	Performer string    `parameter:"performer"`
	Title     string    `parameter:"title"`
}

type SendDocumentArguments struct {
	AbstractSendArguments
	Document  InputFile `parameter:"document"`
	Caption   string    `parameter:"caption"`
	ParseMode ParseMode `parameter:"parse_mode"`
}

type SendVideoArguments struct {
	AbstractSendArguments
	Video             string    `parameter:"video"`
	Duration          int       `parameter:"duration"`
	Width             int       `parameter:"width"`
	Height            int       `parameter:"height"`
	Caption           string    `parameter:"caption"`
	ParseMode         ParseMode `parameter:"parse_mode"`
	SupportsStreaming bool      `parameter:"supports_streaming"`
}

type SendVoiceArguments struct {
	AbstractSendArguments
	Voice     InputFile `parameter:"voice"`
	Caption   string    `parameter:"caption"`
	ParseMode ParseMode `parameter:"parse_mode"`
	Duration  int       `parameter:"duration"`
}

type SendVideoNoteArguments struct {
	AbstractSendArguments
	VideoNote InputFile `parameter:"video_note"`
	Duration  int       `parameter:"duration"`
	Length    int       `parameter:"length"`
}

type SendMediaGroupArguments struct {
	ChatId              string        `parameter:"chat_id"`
	Media               []interface{} `parameter:"media"`
	DisableNotification bool          `parameter:"disable_notification"`
	ReplyToMessageId    int           `parameter:"reply_to_message_id"`
}

type SendLocationArguments struct {
	AbstractSendArguments
	Latitude   float64 `parameter:"latitude"`
	Longitude  float64 `parameter:"longitude"`
	LivePeriod int     `parameter:"live_period"`
}

type EditMessageLiveLocationArguments struct {
	ChatId          string      `parameter:"chat_id"`
	MessageId       int         `parameter:"message_id"`
	InlineMessageId string      `parameter:"inline_message_id"`
	Latitude        float64     `parameter:"latitude"`
	Longitude       float64     `parameter:"longitude"`
	ReplyMarkup     interface{} `parameter:"reply_markup"`
}

type StopMessageLiveLocationArguments struct {
	ChatId          string      `parameter:"chat_id"`
	MessageId       int         `parameter:"message_id"`
	InlineMessageId string      `parameter:"inline_message_id"`
	ReplyMarkup     interface{} `parameter:"reply_markup"`
}

type SendVenueArguments struct {
	AbstractSendArguments
	Latitude     float64 `parameter:"latitude"`
	Longitude    float64 `parameter:"longitude"`
	Title        string  `parameter:"title"`
	Address      string  `parameter:"address"`
	FoursquareId string  `parameter:"foursquare_id"`
}

type SendContentArguments struct {
	AbstractSendArguments
	PhoneNumber string `parameter:"phone_number"`
	FirstName   string `parameter:"first_name"`
	LastName    string `parameter:"last_name"`
}

type SendChatActionArguments struct {
	ChatId string     `parameter:"chat_id"`
	Action ChatAction `parameter:"action"`
}

type GetUserProfilePhotosArguments struct {
	UserId int `parameter:"user_id"`
	Offset int `parameter:"offset"`
	Limit  int `parameter:"limit"`
}

type GetFileArguments struct {
	FileId string `parameter:"file_id"`
}

type ChatOperationArguments struct {
	ChatId string `parameter:"chat_id"`
}
