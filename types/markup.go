package types

type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard"`
	OneTimeKeyboard bool                `json:"one_time_keyboard"`
	Selective       bool                `json:"selective"`
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact"`
	RequestLocation bool                    `json:"request_location"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`
}

type KeyboardButtonPollType struct {
	Type PollType `json:"type"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	Url                          string        `json:"url"`
	LoginUrl                     *LoginUrl     `json:"login_url"`
	CallbackData                 string        `json:"callback_data"`
	SwitchInlineQuery            string        `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	Pay                          bool          `json:"pay"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageId string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}
