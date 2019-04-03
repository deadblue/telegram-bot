package telegroid

import "io"

type GetUpdatesParameters struct {
	_BasicParameters
}
func (p *GetUpdatesParameters) Offset(offset int) {
	p.withInt("offset", offset)
}
func (p *GetUpdatesParameters) Limit(limit int) {
	p.withInt("limit", limit)
}
func (p *GetUpdatesParameters) AllowedUpdates() {
	// TODO
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
func (p *SetWebhookParameters) AllowedUpdates() {
	// TODO
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
func (p *AnswerCallbackQueryParameters) ShowAlert() *SwitchParameter {
	return &SwitchParameter{
		holder: p._BasicParameters,
		name: "show_alert",
	}
}
