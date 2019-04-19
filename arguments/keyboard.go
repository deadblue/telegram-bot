package arguments


type ReplyKeyboardBuilder struct {
	holder _BasicArgs
	name   string

	resize bool
	oneTime bool
	seletive bool
}
func (b *ReplyKeyboardBuilder) ResizeKeyboard() *ReplyKeyboardBuilder {
	b.resize = true
	return b
}
func (b *ReplyKeyboardBuilder) OneTimeKeyboard() *ReplyKeyboardBuilder {
	b.oneTime = true
	return b
}
func (b *ReplyKeyboardBuilder) Seletive() *ReplyKeyboardBuilder {
	b.seletive = true
	return b
}
func (b *ReplyKeyboardBuilder) AddButton() *ReplyKeyboardBuilder {
	// TODO: define `ReplyKeyboardBuilder.AddButton` method
	return b
}
func (b *ReplyKeyboardBuilder) Finish() {
	data := map[string]interface{} {
		"resize_keyboard": b.resize,
		"one_time_keyboard": b.oneTime,
		"selective": b.seletive,
		"keyboard": nil,
	}
	b.holder.withJson(b.name, data)
}


// TODO: define `InlineKeyboardBuilder`
type InlineKeyboardBuilder struct {
	holder _BasicArgs
	name   string
}
func (b *InlineKeyboardBuilder) Finish() {
}
