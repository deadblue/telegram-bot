package arguments


type SendMediaGroupArgs struct {
	ChatArgs
}
func (a *SendMediaGroupArgs) Media() {
	// TODO construct medias
}
func (a *SendMediaGroupArgs) ReplyToMessageId(messageId int) {
	a.getForm().WithInt("reply_to_message_id", messageId)
}
func (a *SendMediaGroupArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}


type EditMessageMediaArgs struct {
	EditMessageReplyMarkupArgs
}
func (p *EditMessageMediaArgs) Media() {
	// TODO construct media
}
