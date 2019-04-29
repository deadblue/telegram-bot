package arguments

// Value switch for boolean argument
type Switch interface {

	// Set the argument to true
	On()

	// Set the argument to false
	Off()
}

type implSwitch struct {
	form *_Form
	name string
}

func (s *implSwitch) On() {
	s.form.WithBool(s.name, true)
}
func (s *implSwitch) Off() {
	s.form.WithBool(s.name, false)
}
