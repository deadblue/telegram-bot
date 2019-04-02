package telegroid


type SetWebhookRequest struct {
	_BasicRequest
}
func (r *SetWebhookRequest) Url(url string) {
	r._WithString("url", url)
}
func (r *SetWebhookRequest) MaxConnections(maxConns int) {
	r._WithInt("max_connections", maxConns)
}


type GetFileRequest struct {
	_BasicRequest
}
func (r *GetFileRequest) FileId(fileId string) {
	r._WithString("file_id", fileId)
}


type GetUpdatesRequest struct {
	_BasicRequest
}
func (r *GetUpdatesRequest) Offset(offset int) {
	r._WithInt("offset", offset)
}
func (r *GetUpdatesRequest) Limit(limit int) {
	r._WithInt("limit", limit)
}
