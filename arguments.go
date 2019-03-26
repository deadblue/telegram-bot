package telegroid

type ParseMode string
type ChatAction string

const (
	ParseModeNormal   ParseMode = ""
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

type InputFileOrString struct {
	FileValue   *InputFile
	StringValue string
}

type SetWebhookArguments struct {
	Url            string     `parameter:"url"`
	Certificate    *InputFile `parameter:"certificate"`
	MaxConnections int        `parameter:"max_connections"`
	AllowedUpdates []string   `parameter:"allowed_updates"`
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
	Photo     InputFileOrString `parameter:"photo"`
	Caption   string            `parameter:"caption"`
	ParseMode ParseMode         `parameter:"parse_mode"`
}

type SendAudioArguments struct {
	AbstractSendArguments
	Audio     InputFileOrString `parameter:"audio"`
	Caption   string            `parameter:"caption"`
	ParseMode ParseMode         `parameter:"parse_mode"`
	Duration  int               `parameter:"duration"`
	Performer string            `parameter:"performer"`
	Title     string            `parameter:"title"`
}

type SendDocumentArguments struct {
	AbstractSendArguments
	Document  InputFileOrString `parameter:"document"`
	Caption   string            `parameter:"caption"`
	ParseMode ParseMode         `parameter:"parse_mode"`
}

type SendVideoArguments struct {
	AbstractSendArguments
	Video             InputFileOrString `parameter:"video"`
	Duration          int               `parameter:"duration"`
	Width             int               `parameter:"width"`
	Height            int               `parameter:"height"`
	Caption           string            `parameter:"caption"`
	ParseMode         ParseMode         `parameter:"parse_mode"`
	SupportsStreaming bool              `parameter:"supports_streaming"`
}

type SendVoiceArguments struct {
	AbstractSendArguments
	Voice     InputFileOrString `parameter:"voice"`
	Caption   string            `parameter:"caption"`
	ParseMode ParseMode         `parameter:"parse_mode"`
	Duration  int               `parameter:"duration"`
}

type SendVideoNoteArguments struct {
	AbstractSendArguments
	VideoNote InputFileOrString `parameter:"video_note"`
	Duration  int               `parameter:"duration"`
	Length    int               `parameter:"length"`
}

type SendMediaGroupArguments struct {
	ChatId              string        `parameter:"chat_id"`
	Media               []interface{} `parameter:"media"`
	DisableNotification bool          `parameter:"disable_notification"`
	ReplyToMessageId    int           `parameter:"reply_to_message_id"`
	Attachments         map[string]InputFile
}

type SendLocationArguments struct {
	AbstractSendArguments
	Latitude   float64 `parameter:"latitude"`
	Longitude  float64 `parameter:"longitude"`
	LivePeriod int     `parameter:"live_period"`
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

type ChatArguments struct {
	ChatId string `parameter:"chat_id"`
}

type ChatMemberArguments struct {
	ChatId string `parameter:"chat_id"`
	UserId int    `parameter:"user_id"`
}

type KickChatMemberArguments struct {
	ChatId    string `parameter:"chat_id"`
	UserId    int    `parameter:"user_id"`
	UntilDate int    `parameter:"until_date"`
}

type RestrictChatMemberArguments struct {
	ChatId                string `parameter:"chat_id"`
	UserId                int    `parameter:"user_id"`
	UntilDate             int    `parameter:"until_date"`
	CanSendMessages       bool   `parameter:"can_send_messages"`
	CanSendMediaMessages  bool   `parameter:"can_send_media_messages"`
	CanSendOtherMessages  bool   `parameter:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `parameter:"can_add_web_page_previews"`
}

type PromoteChatMemberArguments struct {
	ChatId             string `parameter:"chat_id"`
	UserId             int    `parameter:"user_id"`
	CanChangeInfo      bool   `parameter:"can_change_info"`
	CanPostMessages    bool   `parameter:"can_post_messages"`
	CanEditMessages    bool   `parameter:"can_edit_messages"`
	CanDeleteMessages  bool   `parameter:"can_delete_messages"`
	CanInviteUsers     bool   `parameter:"can_invite_users"`
	CanRestrictMembers bool   `parameter:"can_restrict_members"`
	CanPinMessages     bool   `parameter:"can_pin_messages"`
	CanPromoteMembers  bool   `parameter:"can_promote_members"`
}

type SetChatPhotoArguments struct {
	ChatId string    `parameter:"chat_id"`
	Photo  InputFile `parameter:"photo"`
}

type SetChatTitleArguments struct {
	ChatId string `parameter:"chat_id"`
	Title  string `parameter:"title"`
}

type SetChatDescriptionArguments struct {
	ChatId      string `parameter:"chat_id"`
	Description string `parameter:"description"`
}

type PinChatMessageArguments struct {
	ChatId              string `parameter:"chat_id"`
	MessageId           int    `parameter:"message_id"`
	DisableNotification bool   `parameter:"disable_notification"`
}

type SetChatStickerSetArguments struct {
	ChatId         string `parameter:"chat_id"`
	StickerSetName string `parameter:"sticker_set_name"`
}

type AnwserCallbackQueryArguments struct {
	CallbackQueryId string `parameter:"callback_query_id"`
	Text            string `parameter:"text"`
	ShowAlert       bool   `parameter:"show_alert"`
	Url             string `parameter:"url"`
	CacheTime       int    `parameter:"cache_time"`
}

type EditMessageTextArguments struct {
	ChatId                string                `parameter:"chat_id"`
	MessageId             int                   `parameter:"message_id"`
	InlineMessageId       string                `parameter:"inline_message_id"`
	Text                  string                `parameter:"text"`
	ParseMode             ParseMode             `parameter:"parse_mode"`
	DisableWebPagePreview bool                  `parameter:"disable_web_page_preview"`
	ReplyMarkup           *InlineKeyboardMarkup `parameter:"reply_markup"`
}

type EditMessageCaptionArguments struct {
	ChatId          string                `parameter:"chat_id"`
	MessageId       int                   `parameter:"message_id"`
	InlineMessageId string                `parameter:"inline_message_id"`
	Caption         string                `parameter:"caption"`
	ParseMode       ParseMode             `parameter:"parse_mode"`
	ReplyMarkup     *InlineKeyboardMarkup `parameter:"reply_markup"`
}

type EditMessageReplyMarkupArguments struct {
	ChatId          string                `parameter:"chat_id"`
	MessageId       int                   `parameter:"message_id"`
	InlineMessageId string                `parameter:"inline_message_id"`
	ReplyMarkup     *InlineKeyboardMarkup `parameter:"reply_markup"`
}

type DeleteMessageArguments struct {
	ChatId    string `parameter:"chat_id"`
	MessageId int    `parameter:"message_id"`
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

type SendStickerArguments struct {
	AbstractSendArguments
	Sticker InputFileOrString `parameter:"sticker"`
}

type GetStickerSetArguments struct {
	Name string `parameter:"name"`
}

type UploadStickerFileArguments struct {
	UserId     int       `parameter:"user_id"`
	PngSticker InputFile `parameter:"png_sticker"`
}

type CreateNewStickerSetArguments struct {
	UserId        int           `parameter:"user_id"`
	Name          string        `parameter:"name"`
	Title         string        `parameter:"title"`
	PngStricker   InputFile     `parameter:"png_stricker"`
	Emojis        string        `parameter:"emojis"`
	ContainsMasks bool          `parameter:"contains_masks"`
	MaskPosition  *MaskPosition `parameter:"mask_position"`
}

type AddStickerToSetArguments struct {
	UserId       int           `parameter:"user_id"`
	Name         string        `parameter:"name"`
	PngSticker   InputFile     `parameter:"png_sticker"`
	Emojis       string        `parameter:"emojis"`
	MaskPosition *MaskPosition `parameter:"mask_position"`
}

type SetStickerPositionInSetArguments struct {
	Sticker  string `parameter:"sticker"`
	Position int    `parameter:"position"`
}

type DeleteStickerFromSetArguments struct {
	Sticker string `parameter:"sticker"`
}
