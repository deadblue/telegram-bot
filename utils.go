package telegroid

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}
func toLower(ch rune) rune {
	if isUpper(ch) {
		return ch - ('A' - 'a')
	} else {
		return ch
	}
}
func toMethodName(name string) string {
	if len(name) == 0 {
		return ""
	} else {
		runes := []rune(name)
		runes[0] = toLower(runes[0])
		return string(runes)
	}
}

type _Form struct {
	isMultipart bool
	partBuffer  *bytes.Buffer
	partWriter  *multipart.Writer
	values      url.Values
}
// By default, consinder as a simple form
func newForm() *_Form {
	return &_Form{
		isMultipart: false,
		values: url.Values{},
	}
}
func (f *_Form) AddString(name, value string) *_Form {
	if f.isMultipart {
		_ = f.partWriter.WriteField(name, value)
	} else {
		f.values.Set(name, value)
	}
	return f
}
func (f *_Form) AddFile(name, filename string, filedata io.Reader) *_Form {
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
// Once the `Finish` method be called, the form can not be modified anymore.
// This method will return content-type and request data
func (f *_Form) Finish() (contentType string, data io.Reader) {
	if f.isMultipart {
		_ = f.partWriter.Close()
		contentType = f.partWriter.FormDataContentType()
		data = f.partBuffer
	} else {
		contentType = "application/x-www-form-urlencoded"
		data = strings.NewReader(f.values.Encode())
	}
	return
}

