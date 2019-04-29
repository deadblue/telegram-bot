package telegroid

import (
	"fmt"
	"reflect"
)

const (
	apiTemplate = "https://api.telegram.org/bot%s/%s"
	tagMethod   = "method"
)

func (b *Bot) setup(token string) {
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
		fv.Set(b.createInvokerFunction(url, fv.Type()))
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
