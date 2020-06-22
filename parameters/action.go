package parameters

type SendChatActionParams struct {
	ChatParams
}

func (p *SendChatActionParams) Typing() {
	p.set("action", "typing")
}
func (p *SendChatActionParams) UploadPhoto() {
	p.set("action", "upload_photo")
}
func (p *SendChatActionParams) RecordVideo() {
	p.set("action", "record_video")
}
func (p *SendChatActionParams) UploadVideo() {
	p.set("action", "upload_video")
}
func (p *SendChatActionParams) RecordAudio() {
	p.set("action", "record_audio")
}
func (p *SendChatActionParams) UploadAudio() {
	p.set("action", "upload_audio")
}
func (p *SendChatActionParams) UploadDocument() {
	p.set("action", "upload_document")
}
func (p *SendChatActionParams) FindLocation() {
	p.set("action", "find_location")
}
func (p *SendChatActionParams) RecordVideoNote() {
	p.set("action", "record_video_note")
}
func (p *SendChatActionParams) UploadVideoNote() {
	p.set("action", "upload_video_note")
}
