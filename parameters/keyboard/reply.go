package keyboard

import (
	"github.com/deadblue/telegram-bot/types"
)

type ReplyBuilder struct {
	// Buttons on the keyboard
	buttons []*types.KeyboardButton
	// flags
	resize    bool
	oneTime   bool
	selective bool
}

func (b *ReplyBuilder) AddButtons(text ...string) *ReplyBuilder {
	buttons := make([]*types.KeyboardButton, len(text))
	for i, t := range text {
		buttons[i] = &types.KeyboardButton{
			Text: t,
		}
	}
	b.buttons = append(b.buttons, buttons...)
	return b
}
func (b *ReplyBuilder) AddContactButton(text string) *ReplyBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text:           text,
		RequestContact: true,
	})
	return b
}
func (b *ReplyBuilder) AddLocationButton(text string) *ReplyBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text:            text,
		RequestLocation: true,
	})
	return b
}
func (b *ReplyBuilder) AddPollButton(text string) *ReplyBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: "any",
		},
	})
	return b
}
func (b *ReplyBuilder) AddRegularPollButton(text string) *ReplyBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: "regular",
		},
	})
	return b
}
func (b *ReplyBuilder) AddQuizPollButton(text string) *ReplyBuilder {
	b.buttons = append(b.buttons, &types.KeyboardButton{
		Text: text,
		RequestPoll: &types.KeyboardButtonPollType{
			Type: "quiz",
		},
	})
	return b
}
func (b *ReplyBuilder) ResizeKeyboard() *ReplyBuilder {
	b.resize = true
	return b
}
func (b *ReplyBuilder) OneTimeKeyboard() *ReplyBuilder {
	b.oneTime = true
	return b
}
func (b *ReplyBuilder) Selective() *ReplyBuilder {
	b.selective = true
	return b
}
func (b *ReplyBuilder) Build(rowSize ...int) (markup *types.ReplyKeyboardMarkup) {
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
