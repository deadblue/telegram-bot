package telegroid

import (
	"github.com/deadblue/telegroid/parameters"
	"github.com/deadblue/telegroid/types"
)

// Telegram Bot API wrapper
type Bot struct {
	GetMe func() (*types.User, error)

	SendMessage func(args *parameters.SendMessageParams) (*types.Message, error)
}

// Create a Bot instance.
//
// The New() process contains some reflect operations, so I do not suggest
// to new more than one Bot for one bot account. Also the Bot is safe
// for using in multiple goroutines, so it is unnecessary to new many.
func New(token string) *Bot {
	bot := new(Bot)
	upUpDownDownLeftRightLeftRightBA(bot, token)
	return bot
}
