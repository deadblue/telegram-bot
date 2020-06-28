package parameters

import (
	"encoding/json"
	"github.com/deadblue/gostream/multipart"
	"github.com/deadblue/telegram-bot/internal/urlencoded"
	"io"
	"net/http"
	urllib "net/url"
	"strconv"
)

type _file struct {
	name string
	size int64
	data io.Reader
}

type implApiParameters struct {
	// Values
	values urllib.Values
	// Files
	files map[string]_file
}

func (ar *implApiParameters) set(name string, value string) {
	if ar.values == nil {
		ar.values = urllib.Values{}
	}
	ar.values.Set(name, value)
}

func (ar *implApiParameters) setInt(name string, value int) {
	ar.set(name, strconv.Itoa(value))
}

func (ar *implApiParameters) setInt64(name string, value int64) {
	ar.set(name, strconv.FormatInt(value, 10))
}

func (ar *implApiParameters) setBool(name string, value bool) {
	ar.set(name, strconv.FormatBool(value))
}

func (ar *implApiParameters) setFloat(name string, value float64) {
	ar.set(name, strconv.FormatFloat(value, 'f', -1, 64))
}

func (ar *implApiParameters) setJson(name string, value interface{}) {
	if value == nil {
		return
	}
	data, _ := json.Marshal(value)
	ar.set(name, string(data))
}

func (ar *implApiParameters) setFile(name string, file *InputFile) {
	if file.fileIdOrUrl != "" {
		ar.set(name, file.fileIdOrUrl)
	} else {
		if ar.files == nil {
			ar.files = make(map[string]_file)
		}
		ar.files[name] = _file{
			name: file.name,
			size: file.size,
			data: file.data,
		}
	}
}

func (ar *implApiParameters) RequestFor(url string) (req *http.Request, err error) {
	if ar.values == nil || len(ar.values) == 0 {
		// Make a GET request
		req, err = http.NewRequest(http.MethodGet, url, nil)
	} else {
		if ar.files == nil || len(ar.files) == 0 {
			// Url-encoded form
			req, err = urlencoded.NewRequest(url, ar.values)
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
			req, err = multipart.NewRequest(url, form)
		}
	}
	return
}
