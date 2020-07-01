package core

import (
	"context"
	"errors"
	"net/http"
)

var (
	errInvalidParams = errors.New("invalid API parameters")
)

// All API paramaeters should implement this interface.
type ApiParameters interface {
	// Make HTTP request for specific URL.
	RequestFor(ctx context.Context, url string) (req *http.Request, err error)
}

func makeRequest(ctx context.Context, url string, params interface{}) (req *http.Request, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if params == nil {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	} else {
		if ap, ok := params.(ApiParameters); ok {
			req, err = ap.RequestFor(ctx, url)
		} else {
			err = errInvalidParams
		}
	}
	return
}
