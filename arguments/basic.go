package arguments

import (
	"io"
)

// The basic argument struct that implements the ApiArguments
type _BasicArgs struct {

	// The container that holds all arguments.
	form *_Form

	// If this function is not nil, _BasicArgs will call it
	// at the begin in Archive() method. With this, the XxxArgs
	// struct need not override Archive function to do some work.
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

// Alias of the "map[string]interface{}" type.
type _MapValue map[string]interface{}
