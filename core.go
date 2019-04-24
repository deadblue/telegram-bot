package telegroid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)


// All arguments structs should implements this interface, so
// the internal code can get content-type and body for performing request.
type ApiArguments interface {

	// Making the argument values unmodifiable, then
	// return the content-type and request data.
	Archive() (string, io.Reader)

}

const (
	apiTemplate = "https://api.telegram.org/bot%s/%s"
	tagMethod = "method"
)

type _GenericResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

func (b *Bot) setup(token string) {
	// Create HTTP client
	// TODO: Adjust http client config
	b.hc = &http.Client{}
	// Scan API invoker
	rv := reflect.Indirect(reflect.ValueOf(b))
	rt := reflect.TypeOf(b).Elem()
	for i := 0; i < rv.NumField(); i++ {
		fv, ft := rv.Field(i), rt.Field(i)
		if fv.Kind() != reflect.Func {
			continue
		}
		// Get API method name
		mn := ft.Tag.Get(tagMethod)
		if len(mn) == 0 {
			// If not set, use field name
			mn = toMethodName(ft.Name)
		}
		// Make API url
		url := fmt.Sprintf(apiTemplate, token, mn)
		// Create invoker
		fv.Set( b.createInvokerFunction(url, fv.Type()) )
	}
}

func (b *Bot) createInvokerFunction(url string, funcType reflect.Type) reflect.Value {
	invoker := func(args []reflect.Value) (results []reflect.Value) {
		// API invoker function MUST has two out
		resultVal := reflect.New(funcType.Out(0))
		errVal := reflect.New(funcType.Out(1))
		// Call API
		var err error
		if len(args) == 0 {
			err = b.invokeApi(url, nil, resultVal.Interface())
		} else {
			err = b.invokeApi(url, args[0].Interface(), resultVal.Interface())
		}
		// Check error
		if err != nil {
			errVal.Elem().Set(reflect.ValueOf(err))
		}
		return []reflect.Value{resultVal.Elem(), errVal.Elem()}
	}
	return reflect.MakeFunc(funcType, invoker)
}

func (b *Bot) invokeApi(url string, args, result interface{}) (err error) {
	// Build request
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if args != nil {
		if aa, ok := args.(ApiArguments); ok {
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
	// Send request
	resp, err := b.hc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Parse response
	gr := _GenericResponse{}
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err == nil && !gr.Ok {
		err = errors.New(gr.Description)
	}
	if err != nil {
		return
	}
	return json.Unmarshal(gr.Result, result)
}
