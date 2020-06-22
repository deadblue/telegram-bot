package telegroid

import (
	"fmt"
	"github.com/deadblue/telegroid/internal/core"
	"reflect"
)

const (
	apiTemplate = "https://api.telegram.org/bot%s/%s"
	tagMethod   = "method"
)

// Make all API functions for Bot, this function should be called ONLY ONCE for each Bot.
func upUpDownDownLeftRightLeftRightBA(b *Bot, token string) {
	// Agent for performing API request
	agent := core.New(nil)
	// Scan functions from the bot
	rv := reflect.Indirect(reflect.ValueOf(b))
	rt := reflect.TypeOf(b).Elem()
	for i := 0; i < rv.NumField(); i++ {
		fv, ft := rv.Field(i), rt.Field(i)
		if fv.Kind() != reflect.Func {
			continue
		}
		// Get API method name
		methodName := ft.Tag.Get(tagMethod)
		if len(methodName) == 0 {
			// If not set, use field name
			methodName = toMethodName(ft.Name)
		}
		// Make API url
		url := fmt.Sprintf(apiTemplate, token, methodName)
		// Create invoker
		fv.Set(createFunction(agent, url, fv.Type()))
	}
}

func createFunction(agent *core.Agent, url string, funcType reflect.Type) reflect.Value {
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
