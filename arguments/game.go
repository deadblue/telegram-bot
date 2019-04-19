package arguments

type SendGameArgs struct {
	_BasicArgs
}
func (a *SendGameArgs) ChatId(chatId int) {
	a.withInt("chat_id", chatId)
}
func (a *SendGameArgs) ShortName(shortName string) {
	a.withString("game_short_name", shortName)
}
func (a *SendGameArgs) ReplyToMessageId(messageId int) {
	a.withInt("reply_to_message_id", messageId)
}
func (a *SendGameArgs) DisableNotification() {
	a.withBool("disable_notification", true)
}
func (a *SendGameArgs) InlineKeyboard() {
	// TODO construct inline keyboard
}

type GetGameHighScoresArgs struct {
	_BasicArgs
}
func (a *GetGameHighScoresArgs) UserId(userId int) {
	a.withInt("user_id", userId)
}
func (a *GetGameHighScoresArgs) ChatMessageId(chatId int, messageId int) {
	a.withInt("chat_id", chatId).
		withInt("message_id", messageId)
}
func (a *GetGameHighScoresArgs) InlineMessageId(messageId int) {
	a.withInt("inline_message_id", messageId)
}

type SetGameScoreArgs struct {
	GetGameHighScoresArgs
}
func (a *SetGameScoreArgs) Score(score int) {
	a.withInt("score", score)
}
func (a *SetGameScoreArgs) Force() {
	a.withBool("force", true)
}
func (a *SetGameScoreArgs) DisableEditMessage() {
	a.withBool("disable_edit_message", true)
}
