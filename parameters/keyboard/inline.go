package keyboard

import (
	"github.com/deadblue/telegram-bot/types"
)

type InlineBuilder struct {
	// Special button which should be the first button at the first row.
	// May be a game button or pay button.
	specialButton *types.InlineKeyboardButton
	// Normal buttons
	buttons []*types.InlineKeyboardButton
}

func (b *InlineBuilder) AddUrlButton(text, url string) *InlineBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text: text,
		Url:  url,
	})
	return b
}
func (b *InlineBuilder) AddCallbackButton(text, data string) *InlineBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	})
	return b
}
func (b *InlineBuilder) AddLoginButton(text string, loginUrl *types.LoginUrl) *InlineBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text:     text,
		LoginUrl: loginUrl,
	})
	return b
}
func (b *InlineBuilder) AddInlineQueryButton(text, query string) *InlineBuilder {
	b.buttons = append(b.buttons, &types.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: query,
	})
	return b
}
func (b *InlineBuilder) SetGameButton(text string) *InlineBuilder {
	b.specialButton = &types.InlineKeyboardButton{
		Text:         text,
		CallbackGame: &types.CallbackGame{},
	}
	return b
}
func (b *InlineBuilder) SetPayButton(text string) *InlineBuilder {
	b.specialButton = &types.InlineKeyboardButton{
		Text: text,
		Pay:  true,
	}
	return b
}
func (b *InlineBuilder) Build(rowSize ...int) *types.InlineKeyboardMarkup {
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
