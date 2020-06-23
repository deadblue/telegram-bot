package parameters

import "github.com/deadblue/telegroid/types"

type SendGameParams struct {
	implApiParameters
}

func (p *SendGameParams) ChatId(chatId int64) {
	p.setInt64("chat_id", chatId)
}
func (p *SendGameParams) ShortName(name string) {
	p.set("game_short_name", name)
}
func (p *SendGameParams) DisableNotification() {
	p.setBool("disable_notification", true)
}
func (p *SendGameParams) ReplyToMessage(messageId int) {
	p.setInt("reply_to_message_id", messageId)
}
func (p *SendGameParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) {
	p.setJson("reply_markup", markup)
}

type GetGameHighScoresParams struct {
	implApiParameters
}

func (p *GetGameHighScoresParams) UserId(userId int) {
	p.setInt("user_id", userId)
}
func (p *GetGameHighScoresParams) ChatMessage(chatId, messageId int) {
	p.setInt("chat_id", chatId)
	p.setInt("message_id", messageId)
}
func (p *GetGameHighScoresParams) InlineMessage(messageId string) {
	p.set("inline_message_id", messageId)
}

type SetGameScoreParams struct {
	GetGameHighScoresParams
}

func (p *SetGameScoreParams) Score(score int) {
	p.setInt("score", score)
}
func (p *SetGameScoreParams) Force() {
	p.setBool("force", true)
}
func (p *SetGameScoreParams) DisableEditMessage() {
	p.setBool("disable_edit_message", true)
}
