package telegroid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
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

func bindInvoker(tg *Telegroid, token string) {
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
		methodName := ft.Tag.Get(TagMethod)
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

func createInvoker(client * http.Client, token string, methodName string, funcType reflect.Type) InvokeFunction {
	return func(args []reflect.Value) (results []reflect.Value) {
		// API function should has two results
		resultVal := reflect.New( funcType.Out(0) )
		errVal := reflect.New( funcType.Out(1) )
		// Call API
		var err error
		if len(args) == 0 {
			err = invokeAPI(client, token, methodName, nil, resultVal.Interface())
		} else {
			err = invokeAPI(client, token, methodName, args[0].Interface(), resultVal.Interface())
		}
		// Pass error
		if err != nil {
			errVal.Elem().Set( reflect.ValueOf(err) )
		}
		return []reflect.Value{resultVal.Elem(), errVal.Elem()}
	}
}

func invokeAPI(client *http.Client, token string, methodName string, argument, result interface{}) (err error) {
	// build URL
	url := fmt.Sprintf(APITemplate, token, methodName)
	// build request
	req, _ := http.NewRequest("GET", url, nil)
	if argument != nil {
		// setup post
		contentType, body := encodeArguments(argument)
		req.Method, req.Body = "POST", body
		req.Header.Set("Content-Type", contentType)
	}
	// send request
	resp, err := client.Do(req)
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
