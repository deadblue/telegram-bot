package core

import (
	"encoding/json"
	"errors"
	"github.com/deadblue/gostream/quietly"
	"net/http"
	"sync/atomic"
	"time"
)

// The underlying API performer
type Agent struct {
	// HTTP client
	hc *http.Client
	// Cooling down time
	cdTime *int64
}

// Send API request, and parse the response.
func (a *Agent) Send(url string, params, result interface{}) (err error) {
	// Make request
	req, err := internalMakeRequest(url, params)
	if err != nil {
		return
	}
	// Wait for CD
	if cdTime, now := atomic.LoadInt64(a.cdTime), time.Now().Unix(); cdTime > now {
		time.Sleep(time.Duration(cdTime-now) * time.Second)
	}
	// Send request
	resp, err := a.hc.Do(req)
	if err != nil {
		return
	}
	defer quietly.Close(resp.Body)
	// Parse response
	ar := &ApiResponse{}
	if err = json.NewDecoder(resp.Body).Decode(ar); err != nil {
		return
	} else if !ar.Ok {
		// Update CD time
		if ar.Parameters != nil && ar.Parameters.RetryAfter > 0 {
			atomic.StoreInt64(a.cdTime,
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

var (
	errInvalidParams = errors.New("invalid API parameters")
)

func internalMakeRequest(url string, params interface{}) (req *http.Request, err error) {
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

func New(client *http.Client) *Agent {
	if client == nil {
		client = defaultHttpClient()
	}
	cdTime := int64(0)
	return &Agent{
		hc:     client,
		cdTime: &cdTime,
	}
}
