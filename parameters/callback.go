package parameters

type AnswerCallbackQueryParams struct {
	implApiParameters
}

func (p *AnswerCallbackQueryParams) QueryId(queryId string) {
	p.set("callback_query_id", queryId)
}
func (p *AnswerCallbackQueryParams) Notification(text string, alert bool) {
	p.set("text", text)
	p.setBool("show_alert", alert)
}
func (p *AnswerCallbackQueryParams) Url(url string) {
	p.set("url", url)
}
func (p *AnswerCallbackQueryParams) CacheTime(seconds int) {
	p.setInt("cache_time", seconds)
}
