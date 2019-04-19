package arguments

type AnswerInlineQueryArgs struct {
	_BasicArgs
}
func (a *AnswerInlineQueryArgs) QueryId(queryId string) {
	a.withString("inline_query_id", queryId)
}
func (a *AnswerInlineQueryArgs) CacheTime(seconds int) {
	a.withInt("cache_time", seconds)
}
func (a *AnswerInlineQueryArgs) IsPersonal() {
	a.withBool("is_personal", true)
}
func (a *AnswerInlineQueryArgs) NextOffset(offset string) {
	a.withString("next_offset", offset)
}
func (a *AnswerInlineQueryArgs) SwitchPmText(text string) {
	a.withString("switch_pm_text", text)
}
func (a *AnswerInlineQueryArgs) SwitchPmParameter(parameter string) {
	a.withString("switch_pm_parameter", parameter)
}
func (a *AnswerInlineQueryArgs) Results() {
	// TODO construct inline result
}
