package telegroid

import (
	"net/http"
)

type Telegroid struct {
	token          string
	client         *http.Client
	GetMe          func() *User                                `method:"getMe"`
	GetUpdates     func(request GetUpdatesArguments) *[]Update `method:"getUpdates"`
	SetWebhook     func(args SetWebhookArguments) *bool        `method:"setWebhook"`
	DeleteWebhook  func() *bool                                `method:"deleteWebhook"`
	GetWebhookInfo func() *WebhookInfo                         `method:"getWebhookInfo"`
	ForwardMessage func(args ForwardMessageArguments) *Message `method:"forwardMessage"`
	SendMessage    func(args SendMessageArguments) *Message    `method:"sendMessage"`
	SendPhoto      func(args SendPhotoArguments) *Message      `method:"sendPhoto"`
	SendAudio      func(args SendAudioArguments) *Message      `method:"sendAudio"`
	SendDocument   func(args SendDocumentArguments) *Message   `method:"sendDocument"`
	SendVideo      func(args SendVideoArguments) *Message      `method:"sendVideo"`
	SendVoice      func(args SendVoiceArguments) *Message      `method:"sendVoice"`
	SendVideoNote  func(args SendVideoNoteArguments) *Message  `method:"sendVideoNote"`
	SendMediaGroup func(args SendMediaGroupArguments) *Message `method:"sendMediaGroup"`
	SendLocation   func(args SendLocationArguments) *Message   `method:"sendLocation"`
	SendVenue      func(args SendVenueArguments) *Message      `method:"sendVenue"`
	SendContact    func(args SendContentArguments) *Message    `method:"sendContact"`
	SendChatAction func(args SendChatActionArguments) *bool    `method:"sendChatAction"`
}

func NewTelegroid(token string) *Telegroid {
	var td = &Telegroid{
		token:  token,
		client: &http.Client{},
	}
	td.bindInvoker()
	return td
}
