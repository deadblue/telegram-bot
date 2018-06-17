package telegroid

type WebhookInfo struct {
	Url                  string    `json:"url"`
	HasCustomCertificate bool      `json:"has_custom_certificate"`
	PendingUpdateCount   int       `json:"pending_update_count"`
	LastErrorDate        int       `json:"last_error_date"`
	LastErrorMessage     *string   `json:"last_error_message"`
	MaxConnections       int       `json:"max_connections"`
	AllowedUpdates       *[]string `json:"allowed_updates"`
}

type User struct {
	Id           int     `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name"`
	Username     *string `json:"username"`
	LanguageCode *string `json:"language_code"`
}

type Chat struct {
	Id                          int        `json:"id"`
	Type                        string     `json:"type"`
	Title                       *string    `json:"title"`
	Username                    *string    `json:"username"`
	FirstName                   *string    `json:"first_name"`
	LastName                    *string    `json:"last_name"`
	AllMembersAreAdministartors bool       `json:"all_members_are_administartors"`
	Photo                       *ChatPhoto `json:"photo"`
	Description                 *string    `json:"description"`
	InviteLink                  *string    `json:"invite_link"`
	PinnedMessage               *Message   `json:"pinned_message"`
	StickerSetName              *string    `json:"sticker_set_name"`
	CanSetStickerName           bool       `json:"can_set_sticker_name"`
}

type Message struct {
	MessageId            int              `json:"message_id"`
	From                 *User            `json:"from"`
	Date                 int              `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from"`
	ForwardFromChat      *Chat            `json:"forward_from_chat"`
	ForwardFromMessageId int              `json:"forward_from_message_id"`
	ForwardSignature     string           `json:"forward_signature"`
	ForwardDate          int              `json:"forward_date"`
	ReplyToMessage       *Message         `json:"reply_to_message"`
	EditDate             int              `json:"edit_date"`
	MediaGroupId         string           `json:"media_group_id"`
	AuthorSignature      string           `json:"author_signature"`
	Text                 string           `json:"text"`
	Entities             *[]MessageEntity `json:"entities"`
	CaptionEntities      *[]MessageEntity `json:"caption_entities"`
	Audio                *Audio           `json:"audio"`
	Document             *Document        `json:"document"`
	Game                 *Game            `json:"game"`
	Photo                *[]PhotoSize     `json:"photo"`
	Sticker              *Sticker         `json:"sticker"`
	Video                *Video           `json:"video"`
	Voice                *Voice           `json:"voice"`
	VideoNote            *VideoNote       `json:"video_note"`
	Caption              *string          `json:"caption"`
	//Location
	//Venue
	NewChatMembers        *[]User
	LeftChatMember        *User
	NewChatTitle          *string
	NewChatPhoto          *[]PhotoSize
	DeleteChatPhoto       bool
	GroupChatCreated      bool
	SupergroupChatCreated bool
	ChannelChatCreated    bool
	MigrateToChatId       int
	MigrateFromChatId     int
	PinnedMessage         *Message
	//Invooice
	//SuccessfulPayment
	ConnectedWebsite *string
}

type MessageEntity struct {
	Type   string  `json:"type"`
	Offset int     `json:"offset"`
	Length int     `json:"length"`
	Url    *string `json:"url"`
	User   *User   `json:"user"`
}

type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type UserProfilePhotos struct {
	TotalCount int          `parameter:"total_count"`
	Photos     *[]PhotoSize `parameter:"photos"`
}

type File struct {
	FileId   string `parameter:"file_id"`
	FileSize int    `parameter:"file_size"`
	FilePath string `parameter:"file_path"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        *[][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard"`
	OneTimeKeyboard bool                `json:"one_time_keyboard"`
	Selective       bool                `json:"selective"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	Url                          *string       `json:"url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

type ChatPhoto struct {
	SmallFileId string `json:"small_file_id"`
	BigFileId   string `json:"big_file_id"`
}

type MimeFile struct {
	FileId   string  `json:"file_id"`
	MimeType *string `json:"mime_type"`
	FileSize int     `json:"file_size"`
}

type Audio struct {
	MimeFile
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
}

type Document struct {
	MimeFile
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
}

type Video struct {
	MimeFile
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
}

type Voice struct {
	MimeFile
	Duration string `json:"duration"`
}

type VideoNote struct {
	FileId   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int        `json:"file_size"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Sticker struct {
	FileId       int           `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        *string       `json:"emoji"`
	SetName      *string       `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        *[]PhotoSize     `json:"photo"`
	Text         *string          `json:"text"`
	TextEntities *[]MessageEntity `json:"text_entities"`
	Animation    *Animation       `json:"animation"`
}

type Animation struct {
	MimeFile
	Thumb    *PhotoSize `json:"thumb"`
	FileName *string    `json:"file_name"`
}

type CallbackGame struct{}

type InlineQuery struct {
	Id       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageId string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

type Update struct {
	UpdateId           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
}
