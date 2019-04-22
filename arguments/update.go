package arguments


type AllowedUpdatesBuilder interface {
	Message() AllowedUpdatesBuilder
	EditedMessage() AllowedUpdatesBuilder
	ChannelPost() AllowedUpdatesBuilder
	EditedChannelPost() AllowedUpdatesBuilder
	InlineQuery() AllowedUpdatesBuilder
	ChosenInlineResult() AllowedUpdatesBuilder
	CallbackQuery() AllowedUpdatesBuilder
	ShippingQuery() AllowedUpdatesBuilder
	PreCheckoutQuery() AllowedUpdatesBuilder
	Poll() AllowedUpdatesBuilder
	Finish()
}

type implAllowedUpdatesBuilder struct {
	form   *_Form
	values map[string]bool
}
func (b *implAllowedUpdatesBuilder) Message() AllowedUpdatesBuilder {
	b.values[updateMessage] = true
	return b
}
func (b *implAllowedUpdatesBuilder) EditedMessage() AllowedUpdatesBuilder {
	b.values[updateEditedMessage] = true
	return b
}
func (b *implAllowedUpdatesBuilder) ChannelPost() AllowedUpdatesBuilder {
	b.values[updateChannelPost] = true
	return b
}
func (b *implAllowedUpdatesBuilder) EditedChannelPost() AllowedUpdatesBuilder {
	b.values[updateEditedChannelPost] = true
	return b
}
func (b *implAllowedUpdatesBuilder) InlineQuery() AllowedUpdatesBuilder {
	b.values[updateInlineQuery] = true
	return b
}
func (b *implAllowedUpdatesBuilder) ChosenInlineResult() AllowedUpdatesBuilder {
	b.values[updateChosenInlineResult] = true
	return b
}
func (b *implAllowedUpdatesBuilder) CallbackQuery() AllowedUpdatesBuilder {
	b.values[updateCallbackQuery] = true
	return b
}
func (b *implAllowedUpdatesBuilder) ShippingQuery() AllowedUpdatesBuilder {
	b.values[updateShippingQuery] = true
	return b
}
func (b *implAllowedUpdatesBuilder) PreCheckoutQuery() AllowedUpdatesBuilder {
	b.values[updatePreCheckoutQuery] = true
	return b
}
func (b *implAllowedUpdatesBuilder) Poll() AllowedUpdatesBuilder {
	b.values[updatePoll] = true
	return b
}
func (b *implAllowedUpdatesBuilder) Finish() {
	allowed, count := make([]string, len(b.values)), 0
	for k, v := range b.values {
		if v {
			allowed[count] = k
			count += 1
		}
	}
	b.form.WithJson("allowed_updates", allowed[:count])
}
