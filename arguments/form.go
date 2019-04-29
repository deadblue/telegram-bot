package arguments

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
	isMultipart bool
	partBuffer  *bytes.Buffer
	partWriter  *multipart.Writer
	values      url.Values
}

func (f *_Form) WithString(name, value string) *_Form {
	if f.isMultipart {
		_ = f.partWriter.WriteField(name, value)
	} else {
		f.values.Set(name, value)
	}
	return f
}
func (f *_Form) WithInt(name string, value int) *_Form {
	return f.WithString(name, strconv.Itoa(value))
}
func (f *_Form) WithInt64(name string, value int64) *_Form {
	return f.WithString(name, strconv.FormatInt(value, 10))
}
func (f *_Form) WithFloat(name string, value float64) *_Form {
	return f.WithString(name, strconv.FormatFloat(value, 'f', -1, 64))
}
func (f *_Form) WithBool(name string, value bool) *_Form {
	return f.WithString(name, strconv.FormatBool(value))
}
func (f *_Form) WithJson(name string, value interface{}) *_Form {
	if value != nil {
		data, err := json.Marshal(value)
		if err == nil {
			f.WithString(name, string(data))
		}
	}
	return f
}
func (f *_Form) WithFile(name string, file InputFile) *_Form {
	// If this is not a multipart from, convert it to
	if !f.isMultipart {
		f.isMultipart = true
		f.partBuffer = &bytes.Buffer{}
		f.partWriter = multipart.NewWriter(f.partBuffer)
		// Fill multiparts with values
		for key := range f.values {
			_ = f.partWriter.WriteField(key, f.values.Get(key))
			f.values.Del(key)
		}
		// Release values
		f.values = nil
	}
	// Create a file part and fill it
	w, _ := f.partWriter.CreateFormFile(name, file.Name())
	rc := file.Data()
	if rc != nil {
		defer rc.Close()
		_, _ = io.Copy(w, rc)
	}
	return f
}
func (f *_Form) Close() (contentType string, data io.Reader) {
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

// One `_Form` instance CAN NOT be used across multi-goroutines
func newForm() *_Form {
	// By default, consinder as a simple form
	return &_Form{
		isMultipart: false,
		values:      url.Values{},
	}
}
