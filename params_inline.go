package telegroid

type AnswerInlineQueryParameters struct {
	_BasicParameters
}
func (p *AnswerInlineQueryParameters) QueryId(queryId string) {
	p.withString("inline_query_id", queryId)
}
func (p *AnswerInlineQueryParameters) CacheTime(seconds int) {
	p.withInt("cache_time", seconds)
}
func (p *AnswerInlineQueryParameters) IsPersonal() {
	p.withBool("is_personal", true)
}
func (p *AnswerInlineQueryParameters) NextOffset(offset string) {
	p.withString("next_offset", offset)
}
func (p *AnswerInlineQueryParameters) SwitchPmText(text string) {
	p.withString("switch_pm_text", text)
}
func (p *AnswerInlineQueryParameters) SwitchPmParameter(parameter string) {
	p.withString("switch_pm_parameter", parameter)
}
func (p *AnswerInlineQueryParameters) Results() {
	// TODO
}
