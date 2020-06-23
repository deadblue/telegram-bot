package parameters

import "github.com/deadblue/telegroid/types"

type SendInvoiceParams struct {
	implApiParameters
}

func (p *SendInvoiceParams) ChatId(chatId int64) {
	p.setInt64("chat_id", chatId)
}
func (p *SendInvoiceParams) Product(title, description string) {
	p.set("title", title)
	p.set("description", description)
}
func (p *SendInvoiceParams) Provider(token string, data interface{}) {
	p.set("provider_token", token)
	if data != nil {
		p.setJson("provider_data", data)
	}
}
func (p *SendInvoiceParams) Payload(payload string) {
	p.set("payload", payload)
}
func (p *SendInvoiceParams) StartParameter(parameter string) {
	p.set("start_parameter", parameter)
}
func (p *SendInvoiceParams) Prices(currency string, price ...*types.LabeledPrice) {
	p.set("currency", currency)
	p.setJson("prices", price)
}
func (p *SendInvoiceParams) Photo(url string, size int, width int, height int) {
	p.set("photo_url", url)
	if size > 0 {
		p.setInt("photo_size", size)
	}
	if width > 0 && height > 0 {
		p.setInt("photo_width", width)
		p.setInt("photo_height", height)
	}
}
func (p *SendInvoiceParams) NeedName() {
	p.setBool("need_name", true)
}
func (p *SendInvoiceParams) NeedPhoneNumber() {
	p.setBool("need_phone_number", true)
}
func (p *SendInvoiceParams) NeedEmail() {
	p.setBool("need_email", true)
}
func (p *SendInvoiceParams) NeedShippingAddress() {
	p.setBool("need_shipping_address", true)
}
func (p *SendInvoiceParams) SendPhoneNumberToProvider() {
	p.setBool("send_phone_number_to_provider", true)
}
func (p *SendInvoiceParams) SendEmailToProvider() {
	p.setBool("send_email_to_provider", true)
}
func (p *SendInvoiceParams) Flexible() {
	p.setBool("is_flexible", true)
}
func (p *SendInvoiceParams) DisableNotification() {
	p.setBool("disable_notification", true)
}
func (p *SendInvoiceParams) ReplyToMessage(messageId int) {
	p.setInt("reply_to_message_id", messageId)
}
func (p *SendInvoiceParams) InlineKeyboard(markup *types.InlineKeyboardMarkup) {
	p.setJson("reply_markup", markup)
}

type AnswerShippingQueryParams struct {
	implApiParameters
}

func (p *AnswerShippingQueryParams) QueryId(queryId string) {
	p.set("shipping_query_id", queryId)
}
func (p *AnswerShippingQueryParams) Ok(options ...*types.ShippingOption) {
	p.setBool("ok", true)
	p.setJson("shipping_options", options)
}
func (p *AnswerShippingQueryParams) Nok(message string) {
	p.setBool("ok", false)
	p.set("error_message", message)
}

type AnswerPreCheckoutQueryParams struct {
	implApiParameters
}

func (p *AnswerPreCheckoutQueryParams) QueryId(queryId string) {
	p.set("pre_checkout_query_id", queryId)
}
func (p *AnswerPreCheckoutQueryParams) Ok() {
	p.setBool("ok", true)
}
func (p *AnswerPreCheckoutQueryParams) Nok(message string) {
	p.setBool("ok", false)
	p.set("error_message", message)
}
