package params

// Interface for boolean parameter
type Switch interface {

	// Set the value to true
	On()

	// Set the value to false
	Off()
}

type _ParameterSwitch struct {
	holder _BasicParameters
	name   string
}
func (p *_ParameterSwitch) On() {
	p.holder.withBool(p.name, true)
}
func (p *_ParameterSwitch) Off() {
	p.holder.withBool(p.name, false)
}
