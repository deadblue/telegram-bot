package urlencoded

import (
	"context"
	"net/http"
	urllib "net/url"
)

func NewRequest(url string, values urllib.Values) (req *http.Request, err error) {
	return NewRequestWithContext(context.Background(), url, values)
}

func NewRequestWithContext(ctx context.Context, url string, values urllib.Values) (req *http.Request, err error) {
	f := createForm(values)
	req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, f)
	if err == nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.ContentLength = int64(f.Len())
	}
	return
}
