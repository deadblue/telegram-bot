package telegroid

type GetUpdatesRequest struct {
	_BasicParameters
}

func (r *GetUpdatesRequest) Offset(offset int) {
	r.withInt("offset", offset)
}
func (r *GetUpdatesRequest) Limit(limit int) {
	r.withInt("limit", limit)
}
func (r *GetUpdatesRequest) AllowedUpdates() {
	// TODO
}

type SetWebhookRequest struct {
	_BasicParameters
}

func (r *SetWebhookRequest) Url(url string) {
	r.withString("url", url)
}
func (r *SetWebhookRequest) MaxConnections(maxConns int) {
	r.withInt("max_connections", maxConns)
}
func (r *SetWebhookRequest) AllowedUpdates() {
	// TODO
}

type GetFileRequest struct {
	_BasicParameters
}

func (r *GetFileRequest) FileId(fileId string) {
	r.withString("file_id", fileId)
}
