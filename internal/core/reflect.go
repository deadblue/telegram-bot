package core

import (
	"fmt"
	"reflect"
)

const (
	apiTemplate = "https://api.telegram.org/bot%s/%s"
)

// Bind create invoke functions for all function type fields on bot.
func Bind(bot interface{}, token string) {
	// HTTP agent for perform HTTP request
	agent := NewAgent(nil)
	// Scan functions from bot
	rv := reflect.Indirect(reflect.ValueOf(bot))
	rt := reflect.TypeOf(bot).Elem()
	for i := 0; i < rv.NumField(); i++ {
		fv, ft := rv.Field(i), rt.Field(i)
		if fv.Kind() != reflect.Func {
			continue
		}
		// Make API URL
		url := fmt.Sprintf(apiTemplate, token, toMethodName(ft.Name))
		// Create invoker
		fv.Set(createFunction(agent, url, fv.Type()))
	}
}

func toMethodName(name string) string {
	runes := []rune(name)
	if runes[0] >= 'A' && runes[0] <= 'Z' {
		runes[0] -= 'A' - 'a'
	}
	return string(runes)
}

func createFunction(agent *Agent, url string, funcType reflect.Type) reflect.Value {
	apiFunc := func(args []reflect.Value) (results []reflect.Value) {
		// API function MUST has two out
		resultVal := reflect.New(funcType.Out(0))
		errVal := reflect.New(funcType.Out(1))
		// Call API
		var err error
		if len(args) == 0 {
			err = agent.Send(url, nil, resultVal.Interface())
		} else {
			err = agent.Send(url, args[0].Interface(), resultVal.Interface())
		}
		// Check error
		if err != nil {
			errVal.Elem().Set(reflect.ValueOf(err))
		}
		return []reflect.Value{resultVal.Elem(), errVal.Elem()}
	}
	return reflect.MakeFunc(funcType, apiFunc)
}
