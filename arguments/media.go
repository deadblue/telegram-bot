package arguments


type SendMediaGroupArgs struct {
	ChatArgs
}
func (a *SendMediaGroupArgs) Media() {
	// TODO construct medias
}
func (a *SendMediaGroupArgs) ReplyToMessageId(messageId int) {
	a.withInt("reply_to_message_id", messageId)
}
func (a *SendMediaGroupArgs) DisableNotification() {
	a.withBool("disable_notification", true)
}


type EditMessageMediaArgs struct {
	EditMessageReplyMarkupArgs
}
func (p *EditMessageMediaArgs) Media() {
	// TODO construct media
}
