package telegroid

import "fmt"

type ChatRequest struct {
	_BasicRequest
}
func (r *ChatRequest) ChatId(chatId int) {
	r._WithInt("chat_id", chatId)
}
func (r *ChatRequest) Channel(channelName string) {
	r._WithString("chat_id", fmt.Sprintf("@%s", channelName))
}
