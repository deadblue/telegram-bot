package arguments

// The builder to define which updates can received.
type AllowedUpdatesBuilder interface {
	ArgumentBuilder

	// Allow to receive incoming message.
	Message() AllowedUpdatesBuilder

	// Allow to receive new version of a meesage which was edited.
	EditedMessage() AllowedUpdatesBuilder

	// Allow to receive incoming channel post.
	ChannelPost() AllowedUpdatesBuilder

	// Allow to receive new version of a channel post which was edited.
	EditedChannelPost() AllowedUpdatesBuilder

	// Allow to receive incoming inline query.
	InlineQuery() AllowedUpdatesBuilder

	// Allow to receive the result which user chose from an inline query.
	ChosenInlineResult() AllowedUpdatesBuilder

	// Allow to receive incoming callback query.
	CallbackQuery() AllowedUpdatesBuilder

	// Allow to receive incoming shipping query.
	ShippingQuery() AllowedUpdatesBuilder

	// Allow to receive incoming pre-checkout query.
	PreCheckoutQuery() AllowedUpdatesBuilder

	// Allow to receive poll state when it was voted or stopped.
	Poll() AllowedUpdatesBuilder
}

type implAllowedUpdatesBuilder struct {
	form   *_Form
	values map[string]bool
}

func (b *implAllowedUpdatesBuilder) Init(form *_Form) AllowedUpdatesBuilder {
	b.form = form
	b.values = make(map[string]bool)
	return b
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
