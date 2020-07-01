package text

import "github.com/deadblue/telegram-bot/types"

type Plain string

func (t Plain) String() string {
	return string(t)
}
func (t Plain) Mode() types.ParseMode {
	return types.ModePlain
}
