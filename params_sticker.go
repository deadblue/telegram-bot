package telegroid

import "io"

type SendStickerParameters struct {
	_CommenSendParameters
}
func (p *SendStickerParameters) StickerId(stickerId string) {
	p.withString("sticker", stickerId)
}
func (p *SendStickerParameters) StickerFile(fileName string, fileData io.Reader) {
	p.withFile("sticker", fileName, fileData)
}


type GetStickerSetParameters struct {
	_BasicParameters
}
func (p *GetStickerSetParameters) Name(name string) {
	p.withString("name", name)
}


