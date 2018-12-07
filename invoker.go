package telegroid

import (
	"reflect"
	"fmt"
	"encoding/json"
	"net/http"
	"errors"
)

const (
	TagMethod   = "method"

	APITemplate = "https://api.telegram.org/bot%s/%s"
)

type InvokeFunction func(args []reflect.Value) (results []reflect.Value)

type GenericResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description *string         `json:"description"`
}

func (d *Telegroid) bindInvoker() {
	rv := reflect.Indirect(reflect.ValueOf(d))
	rt := reflect.TypeOf(d).Elem()
	// search api function
	for i := 0; i < rv.NumField(); i++ {
		fv, ft := rv.Field(i), rt.Field(i)
		if fv.Kind() != reflect.Func {
			continue
		}
		// get method name from tag
		method := ft.Tag.Get(TagMethod)
		if len(method) == 0 {
			// if not set, use field name
			method = toMethodName(ft.Name)
		}
		// bind invoker
		funcType := fv.Type()
		invoker := d.createInvoker(method, funcType)
		fv.Set(reflect.MakeFunc(funcType, invoker))
	}
}

func (d *Telegroid) createInvoker(method string, funcType reflect.Type) InvokeFunction {
	return func(args []reflect.Value) (results []reflect.Value) {
		// API function should has two results
		resultVal := reflect.New( funcType.Out(0) )
		errVal := reflect.New( funcType.Out(1) )
		// Call API
		var err error
		if len(args) == 0 {
			err = d.invokeAPI(method, nil, resultVal.Interface())
		} else {
			err = d.invokeAPI(method, args[0].Interface(), resultVal.Interface())
		}
		// Pass error
		if err != nil {
			errVal.Elem().Set( reflect.ValueOf(err) )
		}
		return []reflect.Value{resultVal.Elem(), errVal.Elem()}
	}
}

func (d *Telegroid) invokeAPI(method string, argument, result interface{}) (err error) {
	// build URL
	url := fmt.Sprintf(APITemplate, d.token, method)
	// build request
	req, _ := http.NewRequest("GET", url, nil)
	if argument != nil {
		// setup post
		contentType, body := encodeArguments(argument)
		req.Method, req.Body = "POST", body
		req.Header.Set("Content-Type", contentType)
	}
	// send request
	resp, err := d.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// parse response
	respObj := GenericResponse{}
	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return
	}
	if !respObj.Ok {
		return errors.New(*respObj.Description)
	}
	return json.Unmarshal(respObj.Result, result)
}
