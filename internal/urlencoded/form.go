package urlencoded

import (
	"net/url"
	"strings"
)

type form struct {
	r *strings.Reader
}

func (f *form) Seek(offset int64, whence int) (int64, error) {
	return f.r.Seek(offset, whence)
}

func (f *form) Read(b []byte) (int, error) {
	return f.r.Read(b)
}

func (f *form) Close() error {
	return nil
}

func (f *form) Len() int {
	return f.r.Len()
}

func createForm(values url.Values) *form {
	return &form{
		r: strings.NewReader(values.Encode()),
	}
}
