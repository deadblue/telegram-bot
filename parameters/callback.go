package parameters

type AnswerCallbackQueryParams struct {
	implApiParameters
}

func (p *AnswerCallbackQueryParams) CallbackQueryId(queryId string) *AnswerCallbackQueryParams {
	p.set("callback_query_id", queryId)
	return p
}
func (p *AnswerCallbackQueryParams) Text(text string) *AnswerCallbackQueryParams {
	p.set("text", text)
	return p
}
func (p *AnswerCallbackQueryParams) Url(url string) *AnswerCallbackQueryParams {
	p.set("url", url)
	return p
}
func (p *AnswerCallbackQueryParams) CacheTime(seconds int) *AnswerCallbackQueryParams {
	p.setInt("cache_time", seconds)
	return p
}
func (p *AnswerCallbackQueryParams) ShowAlert() *AnswerCallbackQueryParams {
	p.setBool("show_alert", true)
	return p
}
