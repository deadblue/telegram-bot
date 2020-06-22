package telegroid

import "log"

const (
	telegramBotToken = "your-bot-token"
)

func Example() {
	bot := New(telegramBotToken)
	user, err := bot.GetMe()
	if err != nil {
		log.Fatalf("Call getMe error: %s", err.Error())
	} else {
		log.Printf("Bot name: %s", user.Username)
	}
}
