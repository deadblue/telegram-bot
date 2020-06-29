package text

import "github.com/deadblue/telegram-bot/types"

type HTML string

func (t HTML) String() string {
	return string(t)
}
func (t HTML) Mode() types.ParseMode {
	return types.ModeHtml
}
