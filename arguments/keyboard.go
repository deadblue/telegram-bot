package arguments

// The builder for assembling a reply keyboard
type ReplyKeyboardBuilder interface {
	ArgumentBuilder

	// Requests clients to resize the keyboard vertically for optimal fit
	ResizeKeyboard() ReplyKeyboardBuilder

	// Requests clients hide the keyboard once it's been used,
	OneTimeKeyboard() ReplyKeyboardBuilder

	// Set the keyboard shows to specific users only.
	Selective() ReplyKeyboardBuilder

	// Add one or more text buttons.
	AddButtons(text ...string) ReplyKeyboardBuilder

	// Add a contact button, which will send user's phone number when be clicked.
	AddContactButton(text string) ReplyKeyboardBuilder

	// Add a location button, which will send user's current location when be clicked.
	AddLocationButton(text string) ReplyKeyboardBuilder

	// Define how the buttons are arranged.
	//
	// Each rowSize means how many buttons can be placed in that row, and the
	// last rowSize will be followed if there are more buttons need to be placed.
	//
	// If the layout is not specified, all the buttons will be placed in one row.
	//
	// For example, if the layout is
	//     [2]:    The buttons will be splited to multiple rows,
	//             each row contains two buttons at most.
	//     [3, 2]: The buttons will be splited to multiple rows,
	//             the first row contains three buttons,
	//             and the other rows contain two buttons at most.
	Layout(rowSize ...int) ReplyKeyboardBuilder
}

// The builder for assembling a inline keyboard
type InlineKeyboardBuilder interface {
	ArgumentBuilder

	// Add a URL button.
	AddUrlButton(text, url string) InlineKeyboardBuilder

	// Add a callback button.
	AddCallbackButton(text, data string) InlineKeyboardBuilder

	// Add a game button.
	GameButton(text string) InlineKeyboardBuilder

	// Add a pay button.
	PayButton(text string) InlineKeyboardBuilder

	// Define how the buttons are arranged.
	//
	// Each rowSize means how many buttons can be placed in that row, and the
	// last rowSize will be followed if there are more buttons need to be placed.
	//
	// If the layout is not specified, all the buttons will be placed in one row.
	//
	// For example, if the layout is
	//     [2]:    The buttons will be splited to multiple rows,
	//             each row contains two buttons at most.
	//     [3, 2]: The buttons will be splited to multiple rows,
	//             the first row contains three buttons,
	//             and the other rows contain two buttons at most.
	Layout(rowSize ...int) InlineKeyboardBuilder
}

type basicKeyboardBuilder struct {
	layout  []int
	buttons []_MapValue
	// Special button means a game button or a pay button, which
	// should be placed at first, and can only have one at most.
	sepcialButton _MapValue
}

func (b *basicKeyboardBuilder) setLayout(rowSize []int) {
	b.layout = rowSize
}
func (b *basicKeyboardBuilder) appendButton(buttons ..._MapValue) {
	if b.buttons == nil {
		b.buttons = buttons
	} else {
		b.buttons = append(b.buttons, buttons...)
	}
}
func (b *basicKeyboardBuilder) makeKeyboard() [][]_MapValue {
	// Buttons to be arranged
	buttons := b.buttons
	if b.sepcialButton != nil {
		buttons = make([]_MapValue, len(b.buttons)+1)
		buttons[0] = b.sepcialButton
		copy(buttons[1:], b.buttons)
	}
	// If layout is not set, all buttons will be shown in one row
	if b.layout == nil {
		return [][]_MapValue{
			buttons,
		}
	}
	// Arrange buttons with layout
	btnUsed, btnCount := 0, len(buttons)
	rowIndex, rowCount := 0, len(b.layout)
	keyboard := make([][]_MapValue, 0)
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
		keyboard = append(keyboard, buttons[btnUsed:btnUsed+rowSize])
		rowIndex += 1
		btnUsed += rowSize
		btnCount -= rowSize
	}
	return keyboard
}

type implReplyKeyboardBuilder struct {
	basicKeyboardBuilder
	data _MapValue
	form *_Form
	name string
}

func (b *implReplyKeyboardBuilder) set(name string, value interface{}) {
	if b.data == nil {
		b.data = make(_MapValue)
	}
	b.data[name] = value
}
func (b *implReplyKeyboardBuilder) ResizeKeyboard() ReplyKeyboardBuilder {
	b.set("resize_keyboard", true)
	return b
}
func (b *implReplyKeyboardBuilder) OneTimeKeyboard() ReplyKeyboardBuilder {
	b.set("one_time_keyboard", true)
	return b
}
func (b *implReplyKeyboardBuilder) Selective() ReplyKeyboardBuilder {
	b.set("selective", true)
	return b
}
func (b *implReplyKeyboardBuilder) AddButtons(text ...string) ReplyKeyboardBuilder {
	buttons := make([]_MapValue, len(text))
	for i, t := range text {
		buttons[i] = _MapValue{
			"text": t,
		}
	}
	b.appendButton(buttons...)
	return b
}
func (b *implReplyKeyboardBuilder) AddContactButton(text string) ReplyKeyboardBuilder {
	b.appendButton(_MapValue{
		"text":            text,
		"request_contact": true,
	})
	return b
}
func (b *implReplyKeyboardBuilder) AddLocationButton(text string) ReplyKeyboardBuilder {
	b.appendButton(_MapValue{
		"text":             text,
		"request_location": true,
	})
	return b
}
func (b *implReplyKeyboardBuilder) Layout(rowSize ...int) ReplyKeyboardBuilder {
	b.setLayout(rowSize)
	return b
}
func (b *implReplyKeyboardBuilder) Finish() {
	b.set("keyboard", b.makeKeyboard())
	b.form.WithJson(b.name, b.data)
}

type implInlineKeyboardBuilder struct {
	basicKeyboardBuilder
	form *_Form
	name string
}

func (b *implInlineKeyboardBuilder) AddUrlButton(text, url string) InlineKeyboardBuilder {
	b.appendButton(_MapValue{
		"text": text,
		"url":  url,
	})
	return b
}
func (b *implInlineKeyboardBuilder) AddCallbackButton(text, data string) InlineKeyboardBuilder {
	b.appendButton(_MapValue{
		"text":          text,
		"callback_data": data,
	})
	return b
}
func (b *implInlineKeyboardBuilder) GameButton(text string) InlineKeyboardBuilder {
	b.sepcialButton = _MapValue{
		"text":          text,
		"callback_game": struct{}{},
	}
	return b
}
func (b *implInlineKeyboardBuilder) PayButton(text string) InlineKeyboardBuilder {
	b.sepcialButton = _MapValue{
		"text": text,
		"pay":  true,
	}
	return b
}
func (b *implInlineKeyboardBuilder) Layout(rowSize ...int) InlineKeyboardBuilder {
	b.setLayout(rowSize)
	return b
}
func (b *implInlineKeyboardBuilder) Finish() {
	b.form.WithJson(b.name, _MapValue{
		"inline_keyboard": b.makeKeyboard(),
	})
}
