package telegroid

import (
	"reflect"
	"fmt"
	"encoding/json"
	"net/http"
)

const (
	TagMethod    = "method"
	APITemplate = "https://api.telegram.org/bot%s/%s"
)

type InvokeFunction func(args []reflect.Value) (results []reflect.Value)

type GenericResponse struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
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
		methodName := ft.Tag.Get(TagMethod)
		if len(methodName) == 0 {
			continue
		}
		// bind invoker
		funcType := fv.Type()
		invoker := d.createInvoker(methodName, funcType.Out(0))
		fv.Set(reflect.MakeFunc(funcType, invoker))
	}
}

func (d *Telegroid) createInvoker(method string, outType reflect.Type) InvokeFunction {
	return func(args []reflect.Value) (results []reflect.Value) {
		result := reflect.New(outType)
		if len(args) == 0 {
			d.invokeAPI(method, nil, result.Interface())
		} else {
			d.invokeAPI(method, args[0].Interface(), result.Interface())
		}
		return []reflect.Value{result.Elem()}
	}
}

func (d *Telegroid) invokeAPI(method string, argument, result interface{}) (err error) {
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
	// parse response
	respObj := GenericResponse{}
	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return
	}
	return json.Unmarshal(respObj.Result, result)
}
