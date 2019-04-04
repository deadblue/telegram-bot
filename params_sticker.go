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


type UploadStickerFileParameters struct {
	_BasicParameters
}
func (p *UploadStickerFileParameters) UserId(userId int) {
	p.withInt("user_id", userId)
}
func (p *UploadStickerFileParameters) StickerFile(fileName string, fileData io.Reader) {
	p.withFile("png_sticker", fileName, fileData)
}


type AddStickerToSetParameters struct {
	_BasicParameters
}
func (p *AddStickerToSetParameters) UserId(userId int) {
	p.withInt("user_id", userId)
}
func (p *AddStickerToSetParameters) Name(name string) {
	p.withString("name", name)
}
func (p *AddStickerToSetParameters) StickerFile(fileName string, fileData io.Reader) {
	p.withFile("png_sticker", fileName, fileData)
}
func (p *AddStickerToSetParameters) StickerId(fileId string) {
	p.withString("png_sticker", fileId)
}
func (p *AddStickerToSetParameters) Emojis(emojis string) {
	p.withString("emojis", emojis)
}
func (p *AddStickerToSetParameters) MaskPosition(point string, x, y, scale float64) {
	mask := map[string]interface{} {
		"point": point,
		"x_shift": x,
		"y_shift": y,
		"scale": scale,
	}
	p.withJson("mask_position", mask)
}


type CreateNewStickerSetParameters struct {
	AddStickerToSetParameters
}
func (p *CreateNewStickerSetParameters) Title(title string) {
	p.withString("title", title)
}
func (p *CreateNewStickerSetParameters) ContainsMasks() {
	p.withBool("contains_masks", true)
}


type SetStickerPositionInSetParameters struct {
	_BasicParameters
}
func (p *SetStickerPositionInSetParameters) StickerId(fileId string) {
	p.withString("sticker", fileId)
}
func (p *SetStickerPositionInSetParameters) Position(position int) {
	p.withInt("position", position)
}


type DeleteStickerFromSetParameters struct {
	_BasicParameters
}
func (p *DeleteStickerFromSetParameters) StickerId(fileId string) {
	p.withString("sticker", fileId)
}
