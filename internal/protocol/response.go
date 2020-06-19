package protocol

import "encoding/json"

type ApiResponse struct {
	// Success flag.
	Ok bool `json:"ok"`
	// Result data when successful.
	Result json.RawMessage `json:"result"`
	// Error code.
	ErrorCode int `json:"error_code"`
	// Human-readable error description.
	Description string `json:"description"`
	// Extend information for unsuccessful request.
	Parameters *ResponseParameters `json:"parameters"`
}

type ResponseParameters struct {
	// Target chat id has been upgrade to the specific id.
	MigrateToChatId int64 `json:"migrate_to_chat_id"`
	// Next request should be schedule after the specific seconds.
	RetryAfter int `json:"retry_after"`
}
