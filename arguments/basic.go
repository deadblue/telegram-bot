package arguments

import (
	"io"
)

// The basic argument struct that implements the ApiArguments
type _BasicArgs struct {
	form *_Form
}
func (a *_BasicArgs) getForm() *_Form {
	if a.form == nil {
		a.form = newForm()
	}
	return a.form
}
func (a *_BasicArgs) Finish() (contentType string, data io.Reader) {
	return a.getForm().Close()
}
