package arguments

type SendInvoiceArgs struct {
	_BasicArgs
	prices []_MapValue
}

func (a *SendInvoiceArgs) receivePrice(price _MapValue) {
	if a.beforeArchive == nil {
		a.beforeArchive = func() {
			a.getForm().WithJson("prices", a.prices)
		}
	}
	if a.prices == nil {
		a.prices = []_MapValue{price}
	} else {
		a.prices = append(a.prices, price)
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
	a.receivePrice(_MapValue{
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

type AnswerShippingQueryArgs struct {
	_BasicArgs
	options []_MapValue
}

func (a *AnswerShippingQueryArgs) receiveOption(option _MapValue) {
	if a.beforeArchive == nil {
		a.beforeArchive = func() {
			a.getForm().WithJson("shipping_options", a.options)
		}
	}
	if a.options == nil {
		a.options = []_MapValue{option}
	} else {
		a.options = append(a.options, option)
	}
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
func (a *AnswerShippingQueryArgs) AddShippingOption() ShippingOptionBuilder {
	return new(implShippingOptionBuilder).Init(a.receiveOption)
}

type ShippingOptionBuilder interface {
	ArgumentBuilder

	// Set option identifier
	Id(id string) ShippingOptionBuilder

	// Set option title
	Title(title string) ShippingOptionBuilder

	// Add a labeled price
	AddPrice(label string, amount int) ShippingOptionBuilder
}

type implShippingOptionBuilder struct {
	receiver func(_MapValue)
	data     _MapValue
	prices   []_MapValue
}

func (b *implShippingOptionBuilder) Init(receiver func(_MapValue)) ShippingOptionBuilder {
	b.receiver = receiver
	b.data = make(_MapValue)
	b.prices = make([]_MapValue, 0)
	return b
}
func (b *implShippingOptionBuilder) Id(id string) ShippingOptionBuilder {
	b.data["id"] = id
	return b
}
func (b *implShippingOptionBuilder) Title(title string) ShippingOptionBuilder {
	b.data["title"] = title
	return b
}
func (b *implShippingOptionBuilder) AddPrice(label string, amount int) ShippingOptionBuilder {
	b.prices = append(b.prices, _MapValue{
		"label":  label,
		"amount": amount,
	})
	return b
}
func (b *implShippingOptionBuilder) Finish() {
	b.data["prices"] = b.prices
	b.receiver(b.data)
}
