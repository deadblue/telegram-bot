package builder

import (
	"fmt"
	tg "github.com/deadblue/telegroid"
	"strconv"
)

type MessageBuilder struct {
	args tg.AbstractSendArguments
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{args: tg.AbstractSendArguments{}}
}
func (b *MessageBuilder) ChatId(chatId int) *MessageBuilder {
	b.args.ChatId = strconv.Itoa(chatId)
	return b
}
func (b *MessageBuilder) Channel(channel string) *MessageBuilder {
	b.args.ChatId = fmt.Sprintf("@%s", channel)
	return b
}
func (b *MessageBuilder) DisbaleNotification() *MessageBuilder {
	b.args.DisableNotificate = true
	return b
}
func (b *MessageBuilder) ReplyTo(messageId int) *MessageBuilder {
	b.args.ReplyToMessageId = messageId
	return b
}

func (b *MessageBuilder) ForceReply() *MessageBuilder {
	b.args.ReplyMarkup = &tg.ForceReply{ForceReply: true}
	return b
}
func (b *MessageBuilder) RemoveKeyboard() *MessageBuilder {
	b.args.ReplyMarkup = &tg.ReplyKeyboardRemove{RemoveKeyboard: true}
	return b
}

type _KeyboardHolder struct {
	parent   *MessageBuilder
	rowCount int
	oneTime  bool
	rows     [][]*tg.KeyboardButton
}

func (b *MessageBuilder) Keyboard(rowCount int, oneTime bool) *_KeyboardHolder {
	return &_KeyboardHolder{
		parent:   b,
		rowCount: rowCount,
		oneTime:  oneTime,
		rows:     make([][]*tg.KeyboardButton, rowCount),
	}
}
func (h *_KeyboardHolder) AppendButton(rowIndex int, text string) *_KeyboardHolder {
	if rowIndex < 0 {
		rowIndex = 0
	}
	if rowIndex >= h.rowCount {
		rowIndex = h.rowCount - 1
	}
	row := h.rows[rowIndex]
	if row == nil {
		row = []*tg.KeyboardButton{
			{Text: text},
		}
	} else {
		row = append(row, &tg.KeyboardButton{Text: text})
	}
	h.rows[rowIndex] = row
	return h
}
func (h *_KeyboardHolder) Done() *MessageBuilder {
	h.parent.args.ReplyMarkup = &tg.ReplyKeyboardMarkup{
		Keyboard:        h.rows,
		OneTimeKeyboard: h.oneTime,
	}
	return h.parent
}

type _InlineKeyboardHolder struct {
	parent   *MessageBuilder
	rowCount int
	rows     [][]*tg.InlineKeyboardButton
}

func (b *MessageBuilder) InlineKeyboard(rowCount int) *_InlineKeyboardHolder {
	return &_InlineKeyboardHolder{
		parent:   b,
		rowCount: rowCount,
		rows:     make([][]*tg.InlineKeyboardButton, rowCount),
	}
}

func (h *_InlineKeyboardHolder) AppendButton(rowIndex int, text, data string) *_InlineKeyboardHolder {
	if rowIndex < 0 {
		rowIndex = 0
	}
	if rowIndex >= h.rowCount {
		rowIndex = h.rowCount - 1
	}
	row := h.rows[rowIndex]
	if row == nil {
		row = []*tg.InlineKeyboardButton{
			{Text: text, CallbackData: &data},
		}
	} else {
		row = append(row, &tg.InlineKeyboardButton{
			Text: text, CallbackData: &data,
		})
	}
	h.rows[rowIndex] = row
	return h
}

func (h *_InlineKeyboardHolder) Done() *MessageBuilder {
	// set keyboard to parent
	h.parent.args.ReplyMarkup = &tg.InlineKeyboardMarkup{
		InlineKeyboard: h.rows,
	}
	return h.parent
}

// Text message builder
type _TextMessageBuilder struct {
	args *tg.SendMessageArguments
}

func (b *MessageBuilder) ForText() *_TextMessageBuilder {
	return &_TextMessageBuilder{
		args: &tg.SendMessageArguments{
			AbstractSendArguments: b.args,
		},
	}
}
func (tb *_TextMessageBuilder) Text(text string) *_TextMessageBuilder {
	tb.args.ParseMode, tb.args.Text = tg.ParseModeNormal, text
	return tb
}
func (tb *_TextMessageBuilder) Markdown(markdown string) *_TextMessageBuilder {
	tb.args.ParseMode, tb.args.Text = tg.ParseModeMarkdown, markdown
	return tb
}
func (tb *_TextMessageBuilder) Done() *tg.SendMessageArguments {
	return tb.args
}
