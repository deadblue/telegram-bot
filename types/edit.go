package types

import "encoding/json"

type EditMessageResult struct {
	// True when successfully update message.
	Ok bool
	// The edited message if message sent by bot, otherwise nil.
	Message *Message
}

func (r *EditMessageResult) UnmarshalJSON(data []byte) (err error) {
	if data[0] == '{' {
		r.Ok, r.Message = true, &Message{}
		err = json.Unmarshal(data, r.Message)
	} else {
		err = json.Unmarshal(data, &(r.Ok))
	}
	return
}
