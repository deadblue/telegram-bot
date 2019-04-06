package params

import "io"

type AllowedUpdatesBuilder struct {
	holder _BasicParameters
	values map[string]bool
}
func (b *AllowedUpdatesBuilder) Message() {
	b.values[updateMessage] = true
}
func (b *AllowedUpdatesBuilder) EditedMessage() {
	b.values[updateEditedMessage] = true
}
func (b *AllowedUpdatesBuilder) ChannelPost() {
	b.values[updateChannelPost] = true
}
func (b *AllowedUpdatesBuilder) EditedChannelPost() {
	b.values[updateEditedMessage] = true
}
func (b *AllowedUpdatesBuilder) InlineQuery() {
	b.values[updateInlineQuery] = true
}
func (b *AllowedUpdatesBuilder) ChosenInlineResult() {
	b.values[updateChosenInlineResult] = true
}
func (b *AllowedUpdatesBuilder) CallbackQuery() {
	b.values[updateCallbackQuery] = true
}
func (b *AllowedUpdatesBuilder) ShippingQuery() {
	b.values[updateShippingQuery] = true
}
func (b *AllowedUpdatesBuilder) PreCheckoutQuery() {
	b.values[updatePreCheckoutQuery] = true
}
// Call `Finish` to apply you choice
func (b *AllowedUpdatesBuilder) Finish() {
	allowed, count := make([]string, len(b.values)), 0
	for k, v := range b.values {
		if v {
			allowed[count] = k
			count += 1
		}
	}
	b.holder.withJson("allowed_updates", allowed[:count])
}


type GetUpdatesParameters struct {
	_BasicParameters
}
func (p *GetUpdatesParameters) Offset(offset int) {
	p.withInt("offset", offset)
}
func (p *GetUpdatesParameters) Limit(limit int) {
	p.withInt("limit", limit)
}
func (p *GetUpdatesParameters) AllowedUpdates() *AllowedUpdatesBuilder {
	return &AllowedUpdatesBuilder{
		holder: p._BasicParameters,
		values:  make(map[string]bool),
	}
}


type SetWebhookParameters struct {
	_BasicParameters
}
func (p *SetWebhookParameters) Url(url string) {
	p.withString("url", url)
}
func (p *SetWebhookParameters) Certificate(fileName string, fileData io.Reader) {
	p.withFile("certificate", fileName, fileData)
}
func (p *SetWebhookParameters) MaxConnections(maxConns int) {
	p.withInt("max_connections", maxConns)
}
func (p *SetWebhookParameters) AllowedUpdates() *AllowedUpdatesBuilder {
	return &AllowedUpdatesBuilder{
		holder: p._BasicParameters,
		values:  make(map[string]bool),
	}
}


type GetUserProfilePhotosParameters struct {
	_BasicParameters
}
func (p *GetUserProfilePhotosParameters) UserId(userId int) {
	p.withInt("user_id", userId)
}
func (p *GetUserProfilePhotosParameters) Offset(offset int) {
	p.withInt("offset", offset)
}
func (p *GetUserProfilePhotosParameters) Limit(limit int) {
	p.withInt("limit", limit)
}


type GetFileParameters struct {
	_BasicParameters
}
func (p *GetFileParameters) FileId(fileId string) {
	p.withString("file_id", fileId)
}


type AnswerCallbackQueryParameters struct {
	_BasicParameters
}
func (p *AnswerCallbackQueryParameters) CallbackQueryId(callbackQueryId int) {
	p.withInt("callback_query_id", callbackQueryId)
}
func (p *AnswerCallbackQueryParameters) Text(text string) {
	p.withString("text", text)
}
func (p *AnswerCallbackQueryParameters) Url(url string) {
	p.withString("url", url)
}
func (p *AnswerCallbackQueryParameters) CacheTime(seconds int) {
	p.withInt("cache_time", seconds)
}
func (p *AnswerCallbackQueryParameters) ShowAlert() {
	p.withBool("show_alert", true)
}
