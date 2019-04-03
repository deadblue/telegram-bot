package telegroid

import (
	"encoding/json"
	"io"
	"strconv"
)

// The request interface
type ApiParameters interface {
	Finish() (string, io.Reader)
}

// The basic request struct that implements the ApiParameters
type _BasicParameters struct {
	form *_Form
}
func (p *_BasicParameters) withString(name, value string) *_BasicParameters {
	if p.form == nil {
		p.form = newForm()
	}
	p.form.AddString(name, value)
	return p
}
func (p *_BasicParameters) withFile(name, filename string, filedata io.Reader) *_BasicParameters {
	if p.form == nil {
		p.form = newForm()
	}
	p.form.AddFile(name, filename, filedata)
	return p
}
func (p *_BasicParameters) withInt(name string, value int) *_BasicParameters {
	return p.withString(name, strconv.Itoa(value))
}
func (p *_BasicParameters) withFloat(name string, value float64) *_BasicParameters {
	return p.withString(name, strconv.FormatFloat(value, 'f', -1, 64))
}
func (p *_BasicParameters) withBool(name string, value bool) *_BasicParameters {
	return p.withString(name, strconv.FormatBool(value))
}
func (p *_BasicParameters) withJson(name string, value interface{}) *_BasicParameters {
	if value != nil {
		data, err :=json.Marshal(value)
		if err == nil {
			p.withString(name, string(data))
		}
	}
	return p
}
func (p *_BasicParameters) Finish() (contentType string, data io.Reader) {
	return p.form.Finish()
}
