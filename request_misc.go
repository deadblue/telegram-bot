package telegroid


type SetWebhookRequest struct {
	_BasicRequest
}
func (r *SetWebhookRequest) SetUrl(url string) {
	r.withString("url", url)
}
func (r *SetWebhookRequest) SetMaxConnections(maxConns int) {
	r.withInt("max_connections", maxConns)
}


type GetFileRequest struct {
	_BasicRequest
}
func (r *GetFileRequest) WithFileId(fileId string) *GetFileRequest {
	r.withString("file_id", fileId)
	return r
}


type GetUpdatesRequest struct {
	_BasicRequest
}
func (r *GetUpdatesRequest) WithOffset(offset int) *GetUpdatesRequest {
	r.withInt("offset", offset)
	return r
}
func (r *GetUpdatesRequest) WithLimit(limit int) *GetUpdatesRequest {
	r.withInt("limit", limit)
	return r
}
