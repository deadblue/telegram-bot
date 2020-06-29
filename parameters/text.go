package parameters

import (
	"fmt"
	"github.com/deadblue/telegram-bot/types"
)

type FormattedText interface {
	fmt.Stringer
	Mode() types.ParseMode
}
