package arguments

type ShippingOptionsBuilder interface {
	ArgumentBuilder
	ShippingOptionId(id string) ShippingOptionsBuilder
	Title(title string) ShippingOptionsBuilder
	AddPrice(label string, amount int) ShippingOptionsBuilder
}

type implShippingOptionsBuilder struct {
	form   *_Form
	data   _MapValue
	prices []_MapValue
}

func (b *implShippingOptionsBuilder) Init() ShippingOptionsBuilder {
	if b.data == nil {
		b.data = make(_MapValue)
	}
	if b.prices == nil {
		b.prices = make([]_MapValue, 0)
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
	b.prices = append(b.prices, _MapValue{
		"label":  label,
		"amount": amount,
	})
	return b
}
func (b *implShippingOptionsBuilder) Finish() {
	b.data["prices"] = b.prices
	b.form.WithJson("shipping_options", b.data)
}
