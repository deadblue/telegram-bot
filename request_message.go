package telegroid

import (
	"fmt"
	"io"
)

const (
	parseModeMarkdown = "Markdown"
	parseModeHTML     = "HTML"
)


type ForwardMessageRequest struct {
	_BasicRequest
}
func (r *ForwardMessageRequest) WithChatId(chatId int) *ForwardMessageRequest {
	r.withInt("chat_id", chatId)
	return r
}
func (r *ForwardMessageRequest) WithChannel(channelName string) *ForwardMessageRequest {
	r.withString("chat_id", fmt.Sprintf("@%s", channelName))
	return r
}
func (r *ForwardMessageRequest) WithFromChatId(chatId int) *ForwardMessageRequest {
	r.withInt("from_chat_id", chatId)
	return r
}
func (r *ForwardMessageRequest) WithFromChannel(channelName string) *ForwardMessageRequest {
	r.withString("from_chat_id", fmt.Sprintf("@%s", channelName))
	return r
}
func (r *ForwardMessageRequest) WithMessageId(messageId int) *ForwardMessageRequest {
	r.withInt("message_id", messageId)
	return r
}
func (r *ForwardMessageRequest) DisableNotification() *ForwardMessageRequest {
	r.withBool("disable_notification", true)
	return r
}


type SendRequestBuilder struct {
	_BasicRequest
}
func (r *SendRequestBuilder) WithChatId(chatId int) *SendRequestBuilder {
	r.withInt("chat_id", chatId)
	return r
}
func (r *SendRequestBuilder) WithChannel(channel string) *SendRequestBuilder {
	r.withString("chat_id", fmt.Sprintf("@%s", channel))
	return r
}
func (r *SendRequestBuilder) WithReplyToMessageId(messageId int) *SendRequestBuilder {
	r.withInt("reply_to_message_id", messageId)
	return r
}
func (r *SendRequestBuilder) DisableNotification() *SendRequestBuilder {
	r.withBool("disable_notification", true)
	return r
}


type SendMessageRequest struct {
	_BasicRequest
}
func (r *SendRequestBuilder) ForText() *SendMessageRequest {
	return &SendMessageRequest{
		_BasicRequest: r._BasicRequest,
	}
}
func (r *SendMessageRequest) WithText(text string) *SendMessageRequest {
	r.withString("text", text)
	return r
}
func (r *SendMessageRequest) WithMarkdown(text string) *SendMessageRequest {
	r.withString("text", text).
		withString("parse_mode", parseModeMarkdown)
	return r
}
func (r *SendMessageRequest) WithHTML(text string) *SendMessageRequest {
	r.withString("text", text).
		withString("parse_mode", parseModeHTML)
	return r
}
func (r *SendMessageRequest) DisableWebPagePreview() *SendMessageRequest {
	r.withBool("disable_web_page_preview", true)
	return r
}


type SendPhotoRequest struct {
	_BasicRequest
}
func (r *SendRequestBuilder) ForPhoto() *SendPhotoRequest {
	return &SendPhotoRequest{
		_BasicRequest: r._BasicRequest,
	}
}
func (r *SendPhotoRequest) WithCaption(caption string) *SendPhotoRequest {
	r.withString("caption", caption)
	return r
}
func (r *SendPhotoRequest) WithCaptionMarkdown(caption string) *SendPhotoRequest {
	r.withString("caption", caption).
		withString("parse_mode", parseModeMarkdown)
	return r
}
func (r *SendPhotoRequest) WithCaptionHTML(caption string) *SendPhotoRequest {
	r.withString("caption", caption).
		withString("parse_mode", parseModeHTML)
	return r
}
func (r *SendPhotoRequest) WithPhotoFile(filename string, filedata io.Reader) *SendPhotoRequest {
	r.withFile("photo", filename, filedata)
	return r
}
func (r *SendPhotoRequest) WithPhotoFileId(fileId string) *SendPhotoRequest {
	r.withString("photo", fileId)
	return r
}
func (r *SendPhotoRequest) WithPhotoUrl(url string) *SendPhotoRequest {
	r.withString("photo", url)
	return r
}
