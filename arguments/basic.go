package arguments

import (
	"io"
)

// The basic argument struct that implements the ApiArguments
type _BasicArgs struct {
	form *_Form
	beforeArchive func()
}
func (a *_BasicArgs) getForm() *_Form {
	if a.form == nil {
		a.form = newForm()
	}
	return a.form
}
func (a *_BasicArgs) Archive() (contentType string, data io.Reader) {
	if a.beforeArchive != nil {
		a.beforeArchive()
	}
	return a.getForm().Close()
}

// For complex argument such as JSON value,
// telegroid provides builder utilities for them.
//
// All the builders implement this interface.
type ArgumentBuilder interface {

	// Apply the settings.
	// Should be called at the end of setup.
	Finish()

}
