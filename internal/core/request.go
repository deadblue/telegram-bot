package core

import (
	"encoding/json"
	"github.com/deadblue/gostream/multipart"
	"io"
	"net/http"
	urllib "net/url"
	"strconv"
	"strings"
)

type ApiParameters interface {
	RequestFor(url string) (req *http.Request, err error)
}

type file struct {
	name string
	size int64
	data io.Reader
}

type BaseApiParameters struct {
	// Values
	values urllib.Values
	// Files
	files map[string]file
}

func (ar *BaseApiParameters) Set(name string, value string) {
	if ar.values == nil {
		ar.values = urllib.Values{}
	}
	ar.values.Set(name, value)
}

func (ar *BaseApiParameters) SetInt(name string, value int) {
	ar.Set(name, strconv.Itoa(value))
}

func (ar *BaseApiParameters) SetInt64(name string, value int64) {
	ar.Set(name, strconv.FormatInt(value, 10))
}

func (ar *BaseApiParameters) SetBool(name string, value bool) {
	ar.Set(name, strconv.FormatBool(value))
}

func (ar *BaseApiParameters) SetJson(name string, value interface{}) {
	if value == nil {
		return
	}
	data, _ := json.Marshal(value)
	ar.Set(name, string(data))
}

func (ar *BaseApiParameters) SetFile(name string, filename string, size int64, data io.Reader) {
	if ar.files == nil {
		ar.files = make(map[string]file)
	}
	ar.files[name] = file{
		name: filename,
		size: size,
		data: data,
	}
}

func (ar *BaseApiParameters) RequestFor(url string) (req *http.Request, err error) {
	if ar.values == nil || len(ar.values) == 0 {
		// Make a GET request
		req, err = http.NewRequest(http.MethodGet, url, nil)
	} else {
		if ar.files == nil || len(ar.files) == 0 {
			// Url-encoded form
			req, err = http.NewRequest(http.MethodPost, url,
				strings.NewReader(ar.values.Encode()))
		} else {
			// Multipart form
			form := multipart.New()
			// Append values
			for name := range ar.values {
				form.AddValue(name, ar.values.Get(name))
			}
			// Append files
			for name := range ar.files {
				f := ar.files[name]
				form.AddFileData(name, f.name, f.size, f.data)
			}
			req, err = multipart.MakeRequest(url, form)
		}
	}
	return
}
