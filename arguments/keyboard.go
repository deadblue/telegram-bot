package arguments

// The basic keyboard builder for both
// ReplyKeyboardBuilder and InlineKeyboardBuilder
type _BasicKeyboardBuilder struct {
	layout  []int
	buttons []interface{}
}
func (b *_BasicKeyboardBuilder) setLayout(rowSize ...int) {
	b.layout = rowSize
}
func (b *_BasicKeyboardBuilder) appendButton(buttons ...interface{}) {
	if b.buttons == nil {
		b.buttons = buttons
	} else {
		b.buttons = append(b.buttons, buttons...)
	}
}
func (b *_BasicKeyboardBuilder) getKeyboard() [][]interface{} {
	// If user does not set layout, all the buttons
	// will be shown in one row
	if b.layout == nil {
		return [][]interface{} {
			b.buttons,
		}
	}
	// Arrange buttons with the layout
	btnUsed, btnCount := 0, len(b.buttons)
	rowIndex, rowCount := 0, len(b.layout)
	keyboard := make([][]interface{}, 0)
	for btnCount > 0 {
		// How many buttons should be in this row
		rowSize := 0
		if rowIndex < rowCount {
			rowSize = b.layout[rowIndex]
		} else {
			rowSize = b.layout[rowCount - 1]
		}
		if rowSize > btnCount {
			rowSize = btnCount
		}
		// Append the row to keyboard
		keyboard = append(keyboard, b.buttons[btnUsed: btnUsed+rowSize])
		rowIndex += 1
		btnUsed += rowSize
		btnCount -= rowSize
	}
	return keyboard
}

// The builder for ReplyKeyboard
// You should call `Finish` at the end
type ReplyKeyboardBuilder struct {
	_BasicKeyboardBuilder
	holder _BasicArgs
	name   string

	resize   bool
	oneTime  bool
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
func (b *ReplyKeyboardBuilder) AddButtons(text ...string) *ReplyKeyboardBuilder {
	buttons := make([]interface{}, len(text))
	for i, t := range text {
		buttons[i] = map[string]string{
			"text": t,
		}
	}
	b.appendButton(buttons...)
	return b
}
func (b *ReplyKeyboardBuilder) AddContactButton(text string) *ReplyKeyboardBuilder {
	button := map[string]interface{}{
		"text":            text,
		"request_contact": true,
	}
	b.appendButton(button)
	return b
}
func (b *ReplyKeyboardBuilder) AddLocationButton(text string) *ReplyKeyboardBuilder {
	button := map[string]interface{}{
		"text":             text,
		"request_location": true,
	}
	b.appendButton(button)
	return b
}
func (b *ReplyKeyboardBuilder) Layout(rowSize ...int) *ReplyKeyboardBuilder {
	b.setLayout(rowSize...)
	return b
}
func (b *ReplyKeyboardBuilder) Finish() {
	data := map[string]interface{}{
		"resize_keyboard":   b.resize,
		"one_time_keyboard": b.oneTime,
		"selective":         b.seletive,
		"keyboard":          b.getKeyboard(),
	}
	b.holder.withJson(b.name, data)
}

// The builder for InlineKeyboard
// You should call `Finish` at the end
type InlineKeyboardBuilder struct {
	_BasicKeyboardBuilder
	holder _BasicArgs
	name   string
}
func (b *InlineKeyboardBuilder) AddUrlButton(text, url string) *InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text": text,
		"url":  url,
	}
	b.appendButton(button)
	return b
}
func (b *InlineKeyboardBuilder) AddCallbackButton(text, data string) *InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text":          text,
		"callback_data": data,
	}
	b.appendButton(button)
	return b
}
func (b *InlineKeyboardBuilder) AddGameButton(text string) *InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text":          text,
		"callback_game": struct{}{},
	}
	b.appendButton(button)
	return b
}
func (b *InlineKeyboardBuilder) AddPayButton(text string) *InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text": text,
		"pay":  true,
	}
	b.appendButton(button)
	return b
}
func (b *InlineKeyboardBuilder) Layout(rowSize...int) *InlineKeyboardBuilder {
	b.setLayout(rowSize...)
	return b
}
func (b *InlineKeyboardBuilder) Finish() {
	data := map[string]interface{}{
		"inline_keyboard": b.getKeyboard(),
	}
	b.holder.withJson(b.name, data)
}
