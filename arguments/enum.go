package arguments

const (
	// update type
	updateMessage            = "message"
	updateEditedMessage      = "edited_message"
	updateChannelPost        = "channel_post"
	updateEditedChannelPost  = "edited_channel_post"
	updateInlineQuery        = "inline_query"
	updateChosenInlineResult = "chosen_inline_result"
	updateCallbackQuery      = "callback_query"
	updateShippingQuery      = "shipping_query"
	updatePreCheckoutQuery   = "pre_checkout_query"
	updatePoll               = "poll"

	// parse mode
	parseModeMarkdown = "Markdown"
	parseModeHTML     = "HTML"

	// chat action
	chatActionTyping          = "typing"
	chatActionUploadPhoto     = "upload_photo"
	chatActionRecordVideo     = "record_video"
	chatActionUploadVideo     = "upload_video"
	chatActionRecordAudio     = "record_audio"
	chatActionUploadAudio     = "upload_audio"
	chatActionUploadDocument  = "upload_document"
	chatActionFindLocation    = "find_location"
	chatActionRecordVideoNote = "record_video_note"
	chatActionUploadVideoNote = "upload_video_note"

	// media type
	mediaPhoto     = "photo"
	mediaVideo     = "video"
	mediaAnimation = "animation"
	mediaAudio     = "audio"
	mediaDocument  = "document"
)
