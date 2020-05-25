package types

type Message struct {
	MessageId             int                   `json:"message_id"`
	From                  *User                 `json:"from"`
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat"`
	ForwardFromMessageId  int                   `json:"forward_from_message_id"`
	ForwardSignature      string                `json:"forward_signature"`
	ForwardSenderName     string                `json:"forward_sender_name"`
	ForwardDate           int                   `json:"forward_date"`
	ReplyToMessage        *Message              `json:"reply_to_message"`
	EditDate              int                   `json:"edit_date"`
	MediaGroupId          string                `json:"media_group_id"`
	AuthorSignature       string                `json:"author_signature"`
	Text                  string                `json:"text"`
	Entities              []*MessageEntity      `json:"entities"`
	Animation             *Animation            `json:"animation"`
	Audio                 *Audio                `json:"audio"`
	Document              *Document             `json:"document"`
	Photo                 []*PhotoSize          `json:"photo"`
	Sticker               *Sticker              `json:"sticker"`
	Video                 *Video                `json:"video"`
	VideoNote             *VideoNote            `json:"video_note"`
	Voice                 *Voice                `json:"voice"`
	Caption               string                `json:"caption"`
	CaptionEntities       []*MessageEntity      `json:"caption_entities"`
	Contact               *Contact              `json:"contact"`
	Dice                  *Dice                 `json:"dice"`
	Game                  *Game                 `json:"game"`
	Poll                  *Poll                 `json:"poll"`
	Venue                 *Venue                `json:"venue"`
	Location              *Location             `json:"location"`
	NewChatMembers        []*User               `json:"new_chat_members"`
	LeftChatMember        *User                 `json:"left_chat_member"`
	NewChatTitle          string                `json:"new_chat_title"`
	NewChatPhoto          []*PhotoSize          `json:"new_chat_photo"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo"`
	GroupChatCreated      bool                  `json:"group_chat_created"`
	SupergroupChatCreated bool                  `json:"supergroup_chat_created"`
	ChatChatCreated       bool                  `json:"chat_chat_created"`
	MigrateToChatId       int                   `json:"migrate_to_chat_id"`
	MigrateFromChatId     int                   `json:"migrate_from_chat_id"`
	PinnedMessage         *Message              `json:"pinned_message"`
	Invoice               *Invoice              `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment"`
	ConnectedWebsite      string                `json:"connected_website"`
	PassportData          *PassportData         `json:"passport_data"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
}

type MessageEntity struct {
	Type     EntityType `json:"type"`
	Offset   int        `json:"offset"`
	Lenght   int        `json:"lenght"`
	Url      string     `json:"url"`
	User     *User      `json:"user"`
	Language string     `json:"language"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
	Vcard       string `json:"vcard"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FoursquareId   string    `json:"foursquare_id"`
	FoursquareType string    `json:"foursquare_type"`
}
