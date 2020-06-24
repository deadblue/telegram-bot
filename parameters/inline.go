package parameters

// Parameters for answerInlineQuery method.
// Reference: https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQueryParams struct {
	implApiParameters
}

func (p *AnswerInlineQueryParams) QueryId(queryId string) {
	p.set("inline_query_id", queryId)
}
func (p *AnswerInlineQueryParams) CacheTime(seconds int) {
	p.setInt("cache_time", seconds)
}
func (p *AnswerInlineQueryParams) IsPersonal(value bool) {
	p.setBool("is_personal", value)
}
func (p *AnswerInlineQueryParams) NextOffset(offset string) {
	p.set("next_offset", offset)
}
func (p *AnswerInlineQueryParams) SwitchPm(text, parameter string) {
	p.set("switch_pm_text", text)
	p.set("switch_pm_parameter", parameter)
}
func (p *AnswerInlineQueryParams) Results(result ...interface{}) {
	// TODO: Make this method friendly.
	p.setJson("results", result)
}
