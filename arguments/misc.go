package arguments

import "io"

type AllowedUpdatesBuilder struct {
	holder _BasicArgs
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
	b.values[updateEditedChannelPost] = true
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
func (b *AllowedUpdatesBuilder) Poll() {
	b.values[updatePoll] = true
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


type GetUpdatesArgs struct {
	_BasicArgs
}
func (a *GetUpdatesArgs) Offset(offset int) {
	a.withInt("offset", offset)
}
func (a *GetUpdatesArgs) Limit(limit int) {
	a.withInt("limit", limit)
}
func (a *GetUpdatesArgs) AllowedUpdates() *AllowedUpdatesBuilder {
	return &AllowedUpdatesBuilder{
		holder: a._BasicArgs,
		values: make(map[string]bool),
	}
}


type SetWebhookArgs struct {
	_BasicArgs
}
func (a *SetWebhookArgs) Url(url string) {
	a.withString("url", url)
}
func (a *SetWebhookArgs) Certificate(fileName string, fileData io.Reader) {
	a.withFile("certificate", fileName, fileData)
}
func (a *SetWebhookArgs) MaxConnections(maxConns int) {
	a.withInt("max_connections", maxConns)
}
func (a *SetWebhookArgs) AllowedUpdates() *AllowedUpdatesBuilder {
	return &AllowedUpdatesBuilder{
		holder: a._BasicArgs,
		values: make(map[string]bool),
	}
}


type GetUserProfilePhotosArgs struct {
	_BasicArgs
}
func (a *GetUserProfilePhotosArgs) UserId(userId int) {
	a.withInt("user_id", userId)
}
func (a *GetUserProfilePhotosArgs) Offset(offset int) {
	a.withInt("offset", offset)
}
func (a *GetUserProfilePhotosArgs) Limit(limit int) {
	a.withInt("limit", limit)
}


type GetFileArgs struct {
	_BasicArgs
}
func (a *GetFileArgs) FileId(fileId string) {
	a.withString("file_id", fileId)
}


type AnswerCallbackQueryArgs struct {
	_BasicArgs
}
func (a *AnswerCallbackQueryArgs) CallbackQueryId(callbackQueryId int) {
	a.withInt("callback_query_id", callbackQueryId)
}
func (a *AnswerCallbackQueryArgs) Text(text string) {
	a.withString("text", text)
}
func (a *AnswerCallbackQueryArgs) Url(url string) {
	a.withString("url", url)
}
func (a *AnswerCallbackQueryArgs) CacheTime(seconds int) {
	a.withInt("cache_time", seconds)
}
func (a *AnswerCallbackQueryArgs) ShowAlert() {
	a.withBool("show_alert", true)
}
