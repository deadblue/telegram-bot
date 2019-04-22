package arguments

type ReplyKeyboardBuilder interface {
	// Requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to false, in which case the custom keyboard is
	// always of the same height as the app's standard keyboard.
	ResizeKeyboard() ReplyKeyboardBuilder
	// Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display
	// the usual letter-keyboard in the chat – the user can press a special button
	// in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard() ReplyKeyboardBuilder
	// Use this parameter if you want to show the keyboard to specific users only.
	Selective() ReplyKeyboardBuilder
	// Add one or more text buttons.
	Buttons(text ...string) ReplyKeyboardBuilder
	// Add a contact button, which will send user's phone number when be clicked.
	ContactButton(text string) ReplyKeyboardBuilder
	// Add a location button, which will send user's current location when be clicked.
	LocationButton(text string) ReplyKeyboardBuilder
	// Define the layout of the buttons:
	//     If not set:       All the buttons will be shown in one row.
	//     If set as [2]:    The buttons will be splited to mulitple rows,
	//                       each row contains two buttons at most.
	//     If set as [3, 2]: The buttons will be splited to mulitple rows,
	//                       the first row contains three buttons,
	//                       and the other rows contain two buttons at most.
	Layout(rowSize ...int) ReplyKeyboardBuilder
	// Apply the settings, should be called at the end of the setup.
	Finish()
}

type InlineKeyboardBuilder interface {
	// Add a URL button.
	UrlButton(text, url string) InlineKeyboardBuilder
	// Add a callback button.
	CallbackButton(text, data string) InlineKeyboardBuilder
	// Add a game button
	GameButton(text string) InlineKeyboardBuilder
	// Add a pay button
	PayButton(text string) InlineKeyboardBuilder
	// Define the layout of the buttons:
	//     If not set:       All the buttons will be shown in one row.
	//     If set as [2]:    The buttons will be splited to mulitple rows,
	//                       each row contains two buttons at most.
	//     If set as [3, 2]: The buttons will be splited to mulitple rows,
	//                       the first row contains three buttons,
	//                       and the other rows contain two buttons at most.
	Layout(rowSize ...int) InlineKeyboardBuilder
	// Apply the settings, should be called at the end of the setup.
	Finish()
}


type basicKeyboardBuilder struct {
	layout  []int
	buttons []interface{}
}
func (b *basicKeyboardBuilder) setLayout(rowSize ...int) {
	b.layout = rowSize
}
func (b *basicKeyboardBuilder) appendButton(buttons ...interface{}) {
	if b.buttons == nil {
		b.buttons = buttons
	} else {
		b.buttons = append(b.buttons, buttons...)
	}
}
func (b *basicKeyboardBuilder) makeKeyboard() [][]interface{} {
	// If layout is not set, all buttons will be shown in one row
	if b.layout == nil {
		return [][]interface{}{
			b.buttons,
		}
	}
	// Arrange buttons with layout
	btnUsed, btnCount := 0, len(b.buttons)
	rowIndex, rowCount := 0, len(b.layout)
	keyboard := make([][]interface{}, 0)
	for btnCount > 0 {
		// How many buttons should be in this row
		rowSize := 0
		if rowIndex < rowCount {
			rowSize = b.layout[rowIndex]
		} else {
			rowSize = b.layout[rowCount-1]
		}
		if rowSize > btnCount {
			rowSize = btnCount
		}
		// Append the row to keyboard
		keyboard = append(keyboard, b.buttons[btnUsed:btnUsed+rowSize])
		rowIndex += 1
		btnUsed += rowSize
		btnCount -= rowSize
	}
	return keyboard
}


type implReplyKeyboardBuilder struct {
	basicKeyboardBuilder
	form *_Form
	name string
	data map[string]interface{}
}
func (b *implReplyKeyboardBuilder) getData() map[string]interface{} {
	if b.data == nil {
		b.data = make(map[string]interface{})
	}
	return b.data
}
func (b *implReplyKeyboardBuilder) ResizeKeyboard() ReplyKeyboardBuilder {
	b.getData()["resize_keyboard"] = true
	return b
}
func (b *implReplyKeyboardBuilder) OneTimeKeyboard() ReplyKeyboardBuilder {
	b.getData()["one_time_keyboard"] = true
	return b
}
func (b *implReplyKeyboardBuilder) Selective() ReplyKeyboardBuilder {
	b.getData()["selective"] = true
	return b
}
func (b *implReplyKeyboardBuilder) Buttons(text ...string) ReplyKeyboardBuilder {
	buttons := make([]interface{}, len(text))
	for i, t := range text {
		buttons[i] = map[string]string{
			"text": t,
		}
	}
	b.appendButton(buttons...)
	return b
}
func (b *implReplyKeyboardBuilder) ContactButton(text string) ReplyKeyboardBuilder {
	button := map[string]interface{}{
		"text":            text,
		"request_contact": true,
	}
	b.appendButton(button)
	return b
}
func (b *implReplyKeyboardBuilder) LocationButton(text string) ReplyKeyboardBuilder {
	button := map[string]interface{}{
		"text":             text,
		"request_location": true,
	}
	b.appendButton(button)
	return b
}
func (b *implReplyKeyboardBuilder) Layout(rowSize ...int) ReplyKeyboardBuilder {
	b.setLayout(rowSize...)
	return b
}
func (b *implReplyKeyboardBuilder) Finish() {
	data := b.getData()
	data["keyboard"] = b.makeKeyboard()
	b.form.WithJson(b.name, data)
}


type implInlineKeyboardBuilder struct {
	basicKeyboardBuilder
	form *_Form
	name string
}
func (b *implInlineKeyboardBuilder) UrlButton(text, url string) InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text": text,
		"url":  url,
	}
	b.appendButton(button)
	return b
}
func (b *implInlineKeyboardBuilder) CallbackButton(text, data string) InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text":          text,
		"callback_data": data,
	}
	b.appendButton(button)
	return b
}
func (b *implInlineKeyboardBuilder) GameButton(text string) InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text":          text,
		"callback_game": struct{}{},
	}
	b.appendButton(button)
	return b
}
func (b *implInlineKeyboardBuilder) PayButton(text string) InlineKeyboardBuilder {
	button := map[string]interface{}{
		"text": text,
		"pay":  true,
	}
	b.appendButton(button)
	return b
}
func (b *implInlineKeyboardBuilder) Layout(rowSize ...int) InlineKeyboardBuilder {
	b.setLayout(rowSize...)
	return b
}
func (b *implInlineKeyboardBuilder) Finish() {
	data := map[string]interface{}{
		"inline_keyboard": b.makeKeyboard(),
	}
	b.form.WithJson(b.name, data)
}
