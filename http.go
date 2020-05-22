package telegroid

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	defaultDialTimeout  = 30 * time.Second
	defaultIdleConnTime = 10 * time.Minute
	defaultIdleConns    = 100
)

// All arguments structs should implements this interface, so
// the internal code can get content-type and body for performing request.
type _ApiArguments interface {
	// Making the argument values unmodifiable, then return
	// the content-type and request data.
	Archive() (string, io.Reader)
}

// Generic struct for all Bot API responses
type _GenericResponse struct {
	Ok          bool                 `json:"ok"`
	Result      json.RawMessage      `json:"result"`
	Description string               `json:"description"`
	ErrorCode   int                  `json:"error_code"`
	Parameters  *_ResponseParameters `json:"parameters"`
}
type _ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

// Initialize http client with default config
func (b *Bot) initHttpClient() {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: defaultDialTimeout,
		}).DialContext,
		TLSHandshakeTimeout: defaultDialTimeout,
		MaxIdleConnsPerHost: defaultIdleConns,
		IdleConnTimeout:     defaultIdleConnTime,
	}
	b.hc = &http.Client{Transport: transport}
}

// The core of the core to perform the communication
func (b *Bot) invokeApi(url string, args, result interface{}) (err error) {
	// Build request
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if args != nil {
		if aa, ok := args.(_ApiArguments); ok {
			contentType, body := aa.Archive()
			// Update request
			req.Method = http.MethodPost
			req.Header.Set("Content-Type", contentType)
			if rc, ok := body.(io.ReadCloser); ok {
				req.Body = rc
			} else {
				req.Body = ioutil.NopCloser(body)
			}
		}
	}
	// Send request with retry
	for retry := true; retry; {
		resp, err := b.hc.Do(req)
		// Retry when request error
		if err != nil {
			continue
		}
		// Parse response
		gr := _GenericResponse{}
		err = json.NewDecoder(resp.Body).Decode(&gr)
		quietlyClose(resp.Body)
		// Retry when response invalid
		if err != nil {
			continue
		}
		// Handle error
		if !gr.Ok {
			if params := gr.Parameters; params != nil && params.RetryAfter > 0 {
				// Sleep before retry
				time.Sleep(time.Second * time.Duration(params.RetryAfter))
			} else {
				err = errors.New(gr.Description)
				retry = false
			}
		} else {
			err = json.Unmarshal(gr.Result, result)
			retry = false
		}
	}
	return
}
