package arguments

type SendInvoiceArgs struct {
	_BasicArgs
	prices []map[string]interface{}
}

func (a *SendInvoiceArgs) tryInit() {
	if a.prices == nil {
		a.prices = make([]map[string]interface{}, 0)
	}
	if a.beforeArchive == nil {
		a.beforeArchive = func() {
			a.getForm().WithJson("prices", a.prices)
		}
	}
}
func (a *SendInvoiceArgs) ChatId(chatId int) {
	a.getForm().WithInt("chat_id", chatId)
}
func (a *SendInvoiceArgs) Title(title string) {
	a.getForm().WithString("title", title)
}
func (a *SendInvoiceArgs) Description(description string) {
	a.getForm().WithString("description", description)
}
func (a *SendInvoiceArgs) Payload(payload string) {
	a.getForm().WithString("payload", payload)
}
func (a *SendInvoiceArgs) ProviderToken(token string) {
	a.getForm().WithString("provider_token", token)
}
func (a *SendInvoiceArgs) StartParameter(parameter string) {
	a.getForm().WithString("start_parameter", parameter)
}
func (a *SendInvoiceArgs) Currency(currency string) {
	a.getForm().WithString("currency", currency)
}
func (a *SendInvoiceArgs) AddPrice(label string, amount int) {
	a.tryInit()
	a.prices = append(a.prices, map[string]interface{}{
		"label":  label,
		"amount": amount,
	})
}
func (a *SendInvoiceArgs) ProviderData(data string) {
	a.getForm().WithString("provider_data", data)
}
func (a *SendInvoiceArgs) PhotoUrl(url string) {
	a.getForm().WithString("photo_url", url)
}
func (a *SendInvoiceArgs) PhotoSize(size int) {
	a.getForm().WithInt("photo_size", size)
}
func (a *SendInvoiceArgs) PhotoMeasure(width, height int) {
	a.getForm().
		WithInt("photo_width", width).
		WithInt("photo_height", height)
}
func (a *SendInvoiceArgs) NeedName() {
	a.getForm().WithBool("need_name", true)
}
func (a *SendInvoiceArgs) NeedPhoneNumber() {
	a.getForm().WithBool("need_phone_number", true)
}
func (a *SendInvoiceArgs) NeedEmail() {
	a.getForm().WithBool("need_email", true)
}
func (a *SendInvoiceArgs) NeedShippingAddress() {
	a.getForm().WithBool("need_shipping_address", true)
}
func (a *SendInvoiceArgs) SendPhoneNumberToProvider() {
	a.getForm().WithBool("send_phone_number_to_provider", true)
}
func (a *SendInvoiceArgs) SendEmailToProvider() {
	a.getForm().WithBool("send_email_to_provider", true)
}
func (a *SendInvoiceArgs) Flexible() {
	a.getForm().WithBool("is_flexible", true)
}
func (a *SendInvoiceArgs) DisableNotification() {
	a.getForm().WithBool("disable_notification", true)
}
func (a *SendInvoiceArgs) ReplyToMessageId(messageId int) {
	a.getForm().WithInt("reply_to_message_id", messageId)
}
func (a *SendInvoiceArgs) InlineKeyboard() InlineKeyboardBuilder {
	return &implInlineKeyboardBuilder{
		form: a.getForm(),
		name: "reply_markup",
	}
}

type AnswerShippingQueryArgs struct {
	_BasicArgs
}

func (a *AnswerShippingQueryArgs) QueryId(queryId string) {
	a.getForm().WithString("shipping_query_id", queryId)
}
func (a *AnswerShippingQueryArgs) Ok() {
	a.getForm().WithBool("ok", true)
}
func (a *AnswerShippingQueryArgs) ErrorMessage(message string) {
	a.getForm().WithString("error_message", message)
}
func (a *AnswerShippingQueryArgs) ShippingOptions() ShippingOptionsBuilder {
	return (&implShippingOptionsBuilder{
		form: a.getForm(),
	}).Init()
}

type ShippingOptionsBuilder interface {
	ArgumentBuilder
	ShippingOptionId(id string) ShippingOptionsBuilder
	Title(title string) ShippingOptionsBuilder
	AddPrice(label string, amount int) ShippingOptionsBuilder
}

type implShippingOptionsBuilder struct {
	form   *_Form
	data   map[string]interface{}
	prices []map[string]interface{}
}

func (b *implShippingOptionsBuilder) Init() ShippingOptionsBuilder {
	if b.data == nil {
		b.data = make(map[string]interface{})
	}
	if b.prices == nil {
		b.prices = make([]map[string]interface{}, 0)
	}
	return b
}
func (b *implShippingOptionsBuilder) ShippingOptionId(id string) ShippingOptionsBuilder {
	b.data["id"] = id
	return b
}
func (b *implShippingOptionsBuilder) Title(title string) ShippingOptionsBuilder {
	b.data["title"] = title
	return b
}
func (b *implShippingOptionsBuilder) AddPrice(label string, amount int) ShippingOptionsBuilder {
	b.prices = append(b.prices, map[string]interface{}{
		"label":  label,
		"amount": amount,
	})
	return b
}
func (b *implShippingOptionsBuilder) Finish() {
	b.data["prices"] = b.prices
	b.form.WithJson("shipping_options", b.data)
}

type AnswerPreCheckoutQueryArgs struct {
	_BasicArgs
}

func (a *AnswerPreCheckoutQueryArgs) QueryId(queryId string) {
	a.getForm().WithString("pre_checkout_query_id", queryId)
}
func (a *AnswerPreCheckoutQueryArgs) Ok() {
	a.getForm().WithBool("ok", true)
}
func (a *AnswerPreCheckoutQueryArgs) ErrorMessage(message string) {
	a.getForm().WithString("error_message", message)
}
