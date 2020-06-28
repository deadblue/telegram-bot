package urlencoded

import (
	"net/http"
	urllib "net/url"
)

func NewRequest(url string, values urllib.Values) (req *http.Request, err error) {
	f := createForm(values)
	req, err = http.NewRequest(http.MethodPost, url, f)
	if err == nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.ContentLength = int64(f.Len())
	}
	return
}
