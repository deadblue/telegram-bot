package parameters

import (
	"fmt"
	"github.com/deadblue/telegram-bot/types"
	"io"
)

// Reference: https://core.telegram.org/bots/api#inputfile
type InputFile interface {
	Id() string
	File() (name string, size int64, data io.Reader)
}

// Reference: https://core.telegram.org/bots/api#formatting-options
type FormattedText interface {
	fmt.Stringer
	Mode() types.ParseMode
}
