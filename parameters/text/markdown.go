package text

import "github.com/deadblue/telegram-bot/types"

type Markdown string

func (t Markdown) String() string {
	return string(t)
}
func (t Markdown) Mode() types.ParseMode {
	return types.ModeMarkdown
}

type MarkdownV2 string

func (t MarkdownV2) String() string {
	return string(t)
}
func (t MarkdownV2) Mode() types.ParseMode {
	return types.ModeMarkdownV2
}
