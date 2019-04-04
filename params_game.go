package telegroid

type SendGameParameters struct {
	_BasicParameters
}
func (p *SendGameParameters) ChatId(chatId int) {
	p.withInt("chat_id", chatId)
}
func (p *SendGameParameters) ShortName(shortName string) {
	p.withString("game_short_name", shortName)
}
func (p *SendGameParameters) ReplyToMessageId(messageId int) {
	p.withInt("reply_to_message_id", messageId)
}
func (p *SendGameParameters) DisableNotification() {
	p.withBool("disable_notification", true)
}
func (p *SendGameParameters) InlineKeyboard() {
	// TODO
}

type GetGameHighScoresParameters struct {
	_BasicParameters
}
func (p *GetGameHighScoresParameters) UserId(userId int) {
	p.withInt("user_id", userId)
}
func (p *GetGameHighScoresParameters) ChatMessageId(chatId int, messageId int) {
	p.withInt("chat_id", chatId).
		withInt("message_id", messageId)
}
func (p *GetGameHighScoresParameters) InlineMessageId(messageId int) {
	p.withInt("inline_message_id", messageId)
}

type SetGameScoreParameters struct {
	GetGameHighScoresParameters
}
func (p *SetGameScoreParameters) Score(score int) {
	p.withInt("score", score)
}
func (p *SetGameScoreParameters) Force() {
	p.withBool("force", true)
}
func (p *SetGameScoreParameters) DisableEditMessage() {
	p.withBool("disable_edit_message", true)
}
