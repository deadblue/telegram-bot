package arguments

type GetUpdatesArgs struct {
	_BasicArgs
}

func (a *GetUpdatesArgs) Offset(offset int) {
	a.getForm().WithInt("offset", offset)
}
func (a *GetUpdatesArgs) Limit(limit int) {
	a.getForm().WithInt("limit", limit)
}
func (a *GetUpdatesArgs) AllowedUpdates() AllowedUpdatesBuilder {
	return new(implAllowedUpdatesBuilder).Init(a.getForm())
}

type SetWebhookArgs struct {
	_BasicArgs
}

func (a *SetWebhookArgs) Url(url string) {
	a.getForm().WithString("url", url)
}
func (a *SetWebhookArgs) Certificate(file InputFile) {
	a.getForm().WithFile("certificate", file)
}
func (a *SetWebhookArgs) MaxConnections(maxConns int) {
	a.getForm().WithInt("max_connections", maxConns)
}
func (a *SetWebhookArgs) AllowedUpdates() AllowedUpdatesBuilder {
	return new(implAllowedUpdatesBuilder).Init(a.getForm())
}

type GetUserProfilePhotosArgs struct {
	_BasicArgs
}

func (a *GetUserProfilePhotosArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}
func (a *GetUserProfilePhotosArgs) Offset(offset int) {
	a.getForm().WithInt("offset", offset)
}
func (a *GetUserProfilePhotosArgs) Limit(limit int) {
	a.getForm().WithInt("limit", limit)
}

type GetFileArgs struct {
	_BasicArgs
}

func (a *GetFileArgs) FileId(fileId string) {
	a.getForm().WithString("file_id", fileId)
}

type AnswerCallbackQueryArgs struct {
	_BasicArgs
}

func (a *AnswerCallbackQueryArgs) QueryId(queryId string) {
	a.getForm().WithString("callback_query_id", queryId)
}
func (a *AnswerCallbackQueryArgs) Text(text string) {
	a.getForm().WithString("text", text)
}
func (a *AnswerCallbackQueryArgs) Url(url string) {
	a.getForm().WithString("url", url)
}
func (a *AnswerCallbackQueryArgs) CacheTime(seconds int) {
	a.getForm().WithInt("cache_time", seconds)
}
func (a *AnswerCallbackQueryArgs) ShowAlert() {
	a.getForm().WithBool("show_alert", true)
}
