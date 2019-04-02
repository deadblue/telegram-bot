package telegroid

import (
	"encoding/json"
	"io"
	"strconv"
)

// The request interface
type ApiRequest interface {
	Finish() (string, io.Reader)
}

// The basic request struct that implements the ApiRequest
type _BasicRequest struct {
	form *_Form
}
func (r *_BasicRequest) _WithString(name, value string) *_BasicRequest {
	if r.form == nil {
		r.form = newForm()
	}
	r.form.AddString(name, value)
	return r
}
func (r *_BasicRequest) _WithFile(name, filename string, filedata io.Reader) *_BasicRequest {
	if r.form == nil {
		r.form = newForm()
	}
	r.form.AddFile(name, filename, filedata)
	return r
}
func (r *_BasicRequest) _WithInt(name string, value int) *_BasicRequest {
	return r._WithString(name, strconv.Itoa(value))
}
func (r *_BasicRequest) _WithFloat(name string, value float64) *_BasicRequest {
	return r._WithString(name, strconv.FormatFloat(value, 'f', -1, 64))
}
func (r *_BasicRequest) _WithBool(name string, value bool) *_BasicRequest {
	return r._WithString(name, strconv.FormatBool(value))
}
func (r *_BasicRequest) _WithJson(name string, value interface{}) *_BasicRequest {
	if value != nil {
		data, err :=json.Marshal(value)
		if err == nil {
			r._WithString(name, string(data))
		}
	}
	return r
}
func (r *_BasicRequest) Finish() (contentType string, data io.Reader) {
	return r.form.Finish()
}
