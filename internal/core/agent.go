package core

import (
	"encoding/json"
	"errors"
	"github.com/deadblue/gostream/quietly"
	"io"
	"net/http"
	"sync/atomic"
	"time"
)

// The underlying API performer
type Agent struct {
	// HTTP client
	hc *http.Client
	// Cooling down time
	cdTime int64
}

// Invoke the API.
func (a *Agent) Invoke(url string, params, result interface{}) (err error) {
	// Make request
	req, err := makeRequest(url, params)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", UserAgent)
	// Wait for CD
	if cdTime, now := atomic.LoadInt64(&a.cdTime), time.Now().Unix(); cdTime > now {
		time.Sleep(time.Duration(cdTime-now) * time.Second)
	}
	// Send request
	ar := &ApiResponse{}
	if err = a.httpSend(req, ar); err != nil {
		return
	} else if !ar.Ok {
		// Update CD time
		if ar.Parameters != nil && ar.Parameters.RetryAfter > 0 {
			atomic.StoreInt64(&a.cdTime,
				time.Now().Unix()+int64(ar.Parameters.RetryAfter))
		}
		return errors.New(ar.Description)
	}
	// Parse response.result
	if result != nil {
		err = json.Unmarshal(ar.Result, result)
	}
	return
}

func (a *Agent) httpSend(req *http.Request, ar *ApiResponse) (err error) {
	resp := (*http.Response)(nil)
	// Send request with retry
	for {
		resp, err = a.hc.Do(req)
		// Break retry when success
		if err == nil {
			break
		}
		// Check if the request can be re-send or not.
		if req.Body != nil {
			s, ok := req.Body.(io.Seeker)
			// CAN NOT retry for unseekable request body.
			if !ok {
				break
			}
			// CAN NOT retry when seek failed.
			if _, err := s.Seek(0, io.SeekStart); err != nil {
				break
			}
		}
	}
	if err != nil {
		return
	}
	defer quietly.Close(resp.Body)
	return json.NewDecoder(resp.Body).Decode(ar)
}

func NewAgent(client *http.Client) *Agent {
	if client == nil {
		client = defaultHttpClient()
	}
	return &Agent{
		hc:     client,
		cdTime: 0,
	}
}
