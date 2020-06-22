package parameters

import (
	"github.com/deadblue/telegroid/internal/core"
)

type AnswerCallbackQueryParams struct {
	core.BaseApiParameters
}

func (p *AnswerCallbackQueryParams) CallbackQueryId(queryId string) *AnswerCallbackQueryParams {
	p.Set("callback_query_id", queryId)
	return p
}
func (p *AnswerCallbackQueryParams) Text(text string) *AnswerCallbackQueryParams {
	p.Set("text", text)
	return p
}
func (p *AnswerCallbackQueryParams) Url(url string) *AnswerCallbackQueryParams {
	p.Set("url", url)
	return p
}
func (p *AnswerCallbackQueryParams) CacheTime(seconds int) *AnswerCallbackQueryParams {
	p.SetInt("cache_time", seconds)
	return p
}
func (p *AnswerCallbackQueryParams) ShowAlert() *AnswerCallbackQueryParams {
	p.SetBool("show_alert", true)
	return p
}
