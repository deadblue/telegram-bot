package telegroid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)


// The request interface
type ApiArguments interface {

	// When this method be called, it close the inner form,
	// then return the content type and request data
	// which will be used to call the telegram bot API
	//
	// The method is always be called by invoker function,
	// developer can call it for debugging, but should not call
	// when you really use the parameters object
	Archive() (string, io.Reader)

}

const (
	apiTemplate = "https://api.telegram.org/bot%s/%s"
	tagMethod = "method"
)


type _InvokeFunction func(args []reflect.Value) (results []reflect.Value)

type _GenericResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

func bindInvoker(tg *Bot, token string) {
	rv := reflect.Indirect(reflect.ValueOf(tg))
	rt := reflect.TypeOf(tg).Elem()
	client := &http.Client{}
	// search api function
	for i := 0; i < rv.NumField(); i++ {
		fv, ft := rv.Field(i), rt.Field(i)
		if fv.Kind() != reflect.Func {
			continue
		}
		// get method name from tag
		methodName := ft.Tag.Get(tagMethod)
		if len(methodName) == 0 {
			// if not set, use field name
			methodName = toMethodName(ft.Name)
		}
		// bind invoker
		funcType := fv.Type()
		invoker := createInvoker(client, token, methodName, funcType)
		fv.Set(reflect.MakeFunc(funcType, invoker))
	}
}

func createInvoker(client *http.Client, token string, methodName string, funcType reflect.Type) _InvokeFunction {
	return func(args []reflect.Value) (results []reflect.Value) {
		// API function should has two results
		resultVal := reflect.New(funcType.Out(0))
		errVal := reflect.New(funcType.Out(1))
		// Call API
		var err error
		if len(args) == 0 {
			err = invokeAPI(client, token, methodName, nil, resultVal.Interface())
		} else {
			err = invokeAPI(client, token, methodName, args[0].Interface(), resultVal.Interface())
		}
		// Pass error
		if err != nil {
			errVal.Elem().Set(reflect.ValueOf(err))
		}
		return []reflect.Value{resultVal.Elem(), errVal.Elem()}
	}
}

func invokeAPI(client *http.Client, token string, methodName string, args, result interface{}) (err error) {
	// build URL
	url := fmt.Sprintf(apiTemplate, token, methodName)
	// build request
	method, contentType, body := http.MethodGet, "", io.Reader(nil)
	if args != nil {
		if aa, ok := args.(ApiArguments); ok {
			contentType, body = aa.Archive()
			method = http.MethodPost
		}
	}
	req, _ := http.NewRequest(method, url, body)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	// send request
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// parse response
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
