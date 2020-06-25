package parameters

import (
	"github.com/deadblue/telegram-bot/types"
)

type ReplyKeyboardBuilder struct {
	// Buttons on the keyboard
	buttons []*types.KeyboardButton
	// flags
	resize    bool
	oneTime   bool
	selective bool
}

func (b *ReplyKeyboardBuilder) AddButtons(text ...string) *ReplyKeyboardBuilder {
	buttons := make([]*types.KeyboardButton, len(text))
	for i, t := range text {
		buttons[i] = &types.KeyboardButton{
			Text: t,
		}
	}
	b.buttons = append(b.buttons, buttons...)
	return b
}
func (b *ReplyKeyboardBuilder) AddContactButton(text string) *ReplyKeyboardBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text:           text,
		RequestContact: true,
	})
	return b
}
func (b *ReplyKeyboardBuilder) AddLocationButton(text string) *ReplyKeyboardBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text:            text,
		RequestLocation: true,
	})
	return b
}
func (b *ReplyKeyboardBuilder) AddPollButton(text string) *ReplyKeyboardBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: "any",
		},
	})
	return b
}
func (b *ReplyKeyboardBuilder) AddRegularPollButton(text string) *ReplyKeyboardBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: "regular",
		},
	})
	return b
}
func (b *ReplyKeyboardBuilder) AddQuizPollButton(text string) *ReplyKeyboardBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: "quiz",
		},
	})
	return b
}
func (b *ReplyKeyboardBuilder) ResizeKeyboard() *ReplyKeyboardBuilder {
	b.resize = true
	return b
}
func (b *ReplyKeyboardBuilder) OneTimeKeyboard() *ReplyKeyboardBuilder {
	b.oneTime = true
	return b
}
func (b *ReplyKeyboardBuilder) Selective() *ReplyKeyboardBuilder {
	b.selective = true
	return b
}
func (b *ReplyKeyboardBuilder) Build(rowSize ...int) (markup *types.ReplyKeyboardMarkup) {
	// I DO NEED Generics HERE!!!
	btnUsed, btnRemained := 0, len(b.buttons)
	rowIndex, rowCount := 0, len(rowSize)
	keyboard := make([][]*types.KeyboardButton, 0)
	for btnRemained > 0 {
		// Determine row size
		currRowSize := 0
		if rowIndex < rowCount {
			currRowSize = rowSize[rowIndex]
		} else {
			currRowSize = rowSize[rowCount-1]
		}
		if currRowSize > btnRemained {
			currRowSize = btnRemained
		}
		// Set button for current row
		keyboard = append(keyboard, b.buttons[btnUsed:btnUsed+currRowSize])
		rowIndex += 1
		btnUsed += currRowSize
		btnRemained -= currRowSize
	}
	markup = &types.ReplyKeyboardMarkup{
		ResizeKeyboard:  b.resize,
		OneTimeKeyboard: b.oneTime,
		Selective:       b.selective,
		Keyboard:        keyboard,
	}
	return
}

type InlineKeyboardBuilder struct {
	// Special button which should be the first button at the first row.
	// May be a game button or pay button.
	specialButton *types.InlineKeyboardButton
	// Normal buttons
	buttons []*types.InlineKeyboardButton
}

func (b *InlineKeyboardBuilder) AddUrlButton(text, url string) *InlineKeyboardBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text: text,
		Url:  url,
	})
	return b
}
func (b *InlineKeyboardBuilder) AddCallbackButton(text, data string) *InlineKeyboardBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	})
	return b
}
func (b *InlineKeyboardBuilder) AddLoginButton(text string, loginUrl *types.LoginUrl) *InlineKeyboardBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text:     text,
		LoginUrl: loginUrl,
	})
	return b
}
func (b *InlineKeyboardBuilder) AddInlineQueryButton(text, query string) *InlineKeyboardBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	})
	return b
}
func (b *InlineKeyboardBuilder) SetGameButton(text string) *InlineKeyboardBuilder {
	b.specialButton = &types.InlineKeyboardButton{
		Text:         text,
		CallbackGame: &types.CallbackGame{},
	}
	return b
}
func (b *InlineKeyboardBuilder) SetPayButton(text string) *InlineKeyboardBuilder {
	b.specialButton = &types.InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
	return b
}
func (b *InlineKeyboardBuilder) Build(rowSize ...int) *types.InlineKeyboardMarkup {
	// Prepend special button to button slice.
	buttons, btnCount := b.buttons, len(b.buttons)
	if b.specialButton != nil {
		btnCount += 1
		buttons = make([]*types.InlineKeyboardButton, btnCount)
		buttons[0] = b.specialButton
		copy(buttons[1:], b.buttons)
	}
	if btnCount == 0 {
		return nil
	}
	// Build keyboard
	keyboard := make([][]*types.InlineKeyboardButton, 0)
	if rowIndex, rowCount := 0, len(rowSize); rowCount == 0 {
		keyboard = [][]*types.InlineKeyboardButton{
			buttons,
		}
	} else {
		btnUsed, btnRemained := 0, btnCount
		for btnRemained > 0 {
			// Determine row size
			currRowSize := 0
			if rowIndex < rowCount {
				currRowSize = rowSize[rowIndex]
			} else {
				currRowSize = rowSize[rowCount-1]
			}
			if currRowSize > btnRemained {
				currRowSize = btnRemained
			}
			// Set button for current row
			keyboard = append(keyboard, b.buttons[btnUsed:btnUsed+currRowSize])
			rowIndex += 1
			btnUsed += currRowSize
			btnRemained -= currRowSize
		}
	}
	return &types.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}
