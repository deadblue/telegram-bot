package telegroid

import "fmt"

type ChatRequest struct {
	_BasicRequest
}
func (r *ChatRequest) WithChatId(chatId int) *ChatRequest {
	r.withInt("chat_id", chatId)
	return r
}
func (r *ChatRequest) WithChannel(channelName string) *ChatRequest {
	r.withString("chat_id", fmt.Sprintf("@%s", channelName))
	return r
}

