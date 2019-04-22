package arguments

type AnswerInlineQueryArgs struct {
	_BasicArgs
}
func (a *AnswerInlineQueryArgs) QueryId(queryId string) {
	a.getForm().WithString("inline_query_id", queryId)
}
func (a *AnswerInlineQueryArgs) CacheTime(seconds int) {
	a.getForm().WithInt("cache_time", seconds)
}
func (a *AnswerInlineQueryArgs) IsPersonal() {
	a.getForm().WithBool("is_personal", true)
}
func (a *AnswerInlineQueryArgs) NextOffset(offset string) {
	a.getForm().WithString("next_offset", offset)
}
func (a *AnswerInlineQueryArgs) SwitchPmText(text string) {
	a.getForm().WithString("switch_pm_text", text)
}
func (a *AnswerInlineQueryArgs) SwitchPmParameter(parameter string) {
	a.getForm().WithString("switch_pm_parameter", parameter)
}
func (a *AnswerInlineQueryArgs) Results() {
	// TODO construct inline result
}
