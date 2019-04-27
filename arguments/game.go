package arguments

type SendGameArgs struct {
	_BasicArgs
}

func (a *SendGameArgs) ChatId(chatId int) {
	a.getForm().WithInt("chat_id", chatId)
}
func (a *SendGameArgs) ShortName(shortName string) {
	a.getForm().WithString("game_short_name", shortName)
}
func (a *SendGameArgs) ReplyToMessageId(messageId int) {
	a.getForm().WithInt("reply_to_message_id", messageId)
}
func (a *SendGameArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}
func (a *SendGameArgs) InlineKeyboard() InlineKeyboardBuilder {
	return &implInlineKeyboardBuilder{
		form: a.getForm(),
		name: "reply_markup",
	}
}

type _GetGameHighScoresArgs struct {
	_BasicArgs
}

func (a *_GetGameHighScoresArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}
func (a *_GetGameHighScoresArgs) ChatMessageId(chatId int, messageId int) {
	a.getForm().
		WithInt("chat_id", chatId).
		WithInt("message_id", messageId)
}
func (a *_GetGameHighScoresArgs) InlineMessageId(messageId int) {
	a.getForm().WithInt("inline_message_id", messageId)
}

type GetGameHighScoresArgs _GetGameHighScoresArgs

type SetGameScoreArgs struct {
	_GetGameHighScoresArgs
}

func (a *SetGameScoreArgs) Score(score int) {
	a.getForm().WithInt("score", score)
}
func (a *SetGameScoreArgs) Force() {
	a.getForm().WithBool("force", true)
}
func (a *SetGameScoreArgs) DisableEditMessage() {
	a.getForm().WithBool("disable_edit_message", true)
}
