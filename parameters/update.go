package parameters

import (
	"strings"
)

type AllowedUpdates map[string]bool

func (a AllowedUpdates) AllowMessage() AllowedUpdates {
	a["message"] = true
	return a
}
func (a AllowedUpdates) AllowEditedMessage() AllowedUpdates {
	a["edited_message"] = true
	return a
}
func (a AllowedUpdates) AllowChannelPost() AllowedUpdates {
	a["channel_post"] = true
	return a
}
func (a AllowedUpdates) AllowEditedChannelPost() AllowedUpdates {
	a["edited_channel_post"] = true
	return a
}
func (a AllowedUpdates) AllowInlineQuery() AllowedUpdates {
	a["inline_query"] = true
	return a
}
func (a AllowedUpdates) AllowChosenInlineResult() AllowedUpdates {
	a["chosen_inline_result"] = true
	return a
}
func (a AllowedUpdates) AllowCallbackQuery() AllowedUpdates {
	a["callback_query"] = true
	return a
}
func (a AllowedUpdates) AllowShippingQuery() AllowedUpdates {
	a["shipping_query"] = true
	return a
}
func (a AllowedUpdates) AllowPreCheckoutQuery() AllowedUpdates {
	a["pre_checkout_query"] = true
	return a
}
func (a AllowedUpdates) AllowPoll() AllowedUpdates {
	a["poll"] = true
	return a
}
func (a AllowedUpdates) AllowPollAnswer() AllowedUpdates {
	a["poll_answer"] = true
	return a
}
func (a AllowedUpdates) Value() string {
	buf, empty := strings.Builder{}, true
	buf.WriteString("[")
	for key, value := range a {
		if !value {
			continue
		}
		if !empty {
			buf.WriteString(", ")
		}
		buf.WriteString("\"" + key + "\"")
		empty = false
	}
	buf.WriteString("]")
	return buf.String()
}

type GetUpdateParams struct {
	implApiParameters
}

func (p *GetUpdateParams) Range(offset, limit int) {
	p.setInt("offset", offset)
	p.setInt("limit", limit)
}
func (p *GetUpdateParams) Timeout(timeout int) {
	p.setInt("timeout", timeout)
}
func (p *GetUpdateParams) AllowedUpdates(updates AllowedUpdates) *GetUpdateParams {
	if updates != nil {
		p.set("allowed_updates", updates.Value())
	}
	return p
}

type SetWebhookParams struct {
	implApiParameters
}

func (p *SetWebhookParams) Url(url string) {
	p.set("url", url)
}
func (p *SetWebhookParams) Certificate(file InputFile) {
	p.setFile("certificate", file)
}
func (p *SetWebhookParams) MaxConnections(maxConns int) {
	p.setInt("max_connections", maxConns)
}
func (p *SetWebhookParams) AllowedUpdates(updates AllowedUpdates) *SetWebhookParams {
	if updates != nil {
		p.set("allowed_updates", updates.Value())
	}
	return p
}
