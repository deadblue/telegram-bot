package telegroid

import (
	"io"
	"strconv"
)

// The request interface for all
type ApiRequest interface {
	Done() (string, io.Reader)
}

// The basic request struct that implements the ApiRequest
type _BasicRequest struct {
	form *_Form
}
func (r *_BasicRequest) withString(name, value string) *_BasicRequest {
	if r.form == nil {
		r.form = newForm()
	}
	r.form.AddString(name, value)
	return r
}
func (r *_BasicRequest) withInt(name string, value int) *_BasicRequest {
	return r.withString(name, strconv.Itoa(value))
}
func (r *_BasicRequest) withBool(name string, value bool) *_BasicRequest {
	return r.withString(name, strconv.FormatBool(value))
}
func (r *_BasicRequest) withFile(name, filename string, filedata io.Reader) *_BasicRequest {
	if r.form == nil {
		r.form = newForm()
	}
	r.form.AddFile(name, filename, filedata)
	return r
}
func (r *_BasicRequest) Done() (contentType string, data io.Reader) {
	return r.form.Finish()
}
