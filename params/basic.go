package params

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
)

type _Form struct {
	closed      bool
	isMultipart bool
	partBuffer  *bytes.Buffer
	partWriter  *multipart.Writer
	values      url.Values
}
func newForm() *_Form {
	// By default, consinder as a simple form
	return &_Form{
		closed:      false,
		isMultipart: false,
		values:      url.Values{},
	}
}
func (f *_Form) AddString(name, value string) *_Form {
	if f.closed {
		return f
	}
	if f.isMultipart {
		_ = f.partWriter.WriteField(name, value)
	} else {
		f.values.Set(name, value)
	}
	return f
}
func (f *_Form) AddFile(name, filename string, filedata io.Reader) *_Form {
	if f.closed {
		return f
	}
	// If this is not a multipart from, convert it
	if !f.isMultipart {
		// TODO: Add mutex to support multi-goroutines
		f.isMultipart = true
		f.partBuffer = &bytes.Buffer{}
		f.partWriter = multipart.NewWriter(f.partBuffer)
		// fill multipars with values
		for key := range f.values {
			_ = f.partWriter.WriteField(key, f.values.Get(key))
			f.values.Del(key)
		}
		// release values
		f.values = nil
	}
	w, _ := f.partWriter.CreateFormFile(name, filename)
	_, _ = io.Copy(w, filedata)
	return f
}
func (f *_Form) Close() (contentType string, data io.Reader) {
	if f.isMultipart {
		if !f.closed {
			f.closed = true
			_ = f.partWriter.Close()
		}
		contentType = f.partWriter.FormDataContentType()
		data = f.partBuffer
	} else {
		contentType = "application/x-www-form-urlencoded"
		data = strings.NewReader(f.values.Encode())
	}
	return
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
func (p *_BasicParameters) withInt64(name string, value int64) *_BasicParameters {
	return p.withString(name, strconv.FormatInt(value, 10))
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
	return p.form.Close()
}
