package core

import (
	"errors"
	"net/http"
)

var (
	errInvalidParams = errors.New("invalid API parameters")
)

// All API paramaeters should implement this interface.
type ApiParameters interface {
	// Make HTTP request for specific URL.
	RequestFor(url string) (req *http.Request, err error)
}

func makeRequest(url string, params interface{}) (req *http.Request, err error) {
	if params == nil {
		req, err = http.NewRequest(http.MethodGet, url, nil)
	} else {
		if ap, ok := params.(ApiParameters); ok {
			req, err = ap.RequestFor(url)
		} else {
			err = errInvalidParams
		}
	}
	return
}
