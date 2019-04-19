package params


type SendMediaGroupParameters struct {
	ChatParameters
}
func (p *SendMediaGroupParameters) Media() {
	// TODO construct medias
}
func (p *SendMediaGroupParameters) ReplyToMessageId(messageId int) {
	p.withInt("reply_to_message_id", messageId)
}
func (p *SendMediaGroupParameters) DisableNotification() {
	p.withBool("disable_notification", true)
}


type EditMessageMediaParameters struct {
	EditMessageReplyMarkupParameters
}
func (p *EditMessageMediaParameters) Media() {
	// TODO construct media
}
