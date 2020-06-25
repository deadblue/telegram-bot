package parameters

import (
	"fmt"
	"github.com/deadblue/telegram-bot/types"
)

type FormattedText interface {
	fmt.Stringer
	Mode() types.ParseMode
}

type PlainText string

func (t PlainText) String() string {
	return string(t)
}
func (t PlainText) Mode() types.ParseMode {
	return types.ModePlain
}

type HTMLText string

func (t HTMLText) String() string {
	return string(t)
}
func (t HTMLText) Mode() types.ParseMode {
	return types.ModeHtml
}

type MarkdownText string

func (t MarkdownText) String() string {
	return string(t)
}
func (t MarkdownText) Mode() types.ParseMode {
	return types.ModeMarkdown
}

type MarkdownV2Text string

func (t MarkdownV2Text) String() string {
	return string(t)
}
func (t MarkdownV2Text) Mode() types.ParseMode {
	return types.ModeMarkdownV2
}
