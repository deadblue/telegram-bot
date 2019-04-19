package arguments

import "io"

type SendStickerArgs struct {
	_CommonSendArgs
}
func (a *SendStickerArgs) StickerId(stickerId string) {
	a.withString("sticker", stickerId)
}
func (a *SendStickerArgs) StickerFile(fileName string, fileData io.Reader) {
	a.withFile("sticker", fileName, fileData)
}


type GetStickerSetArgs struct {
	_BasicArgs
}
func (a *GetStickerSetArgs) Name(name string) {
	a.withString("name", name)
}


type UploadStickerFileArgs struct {
	_BasicArgs
}
func (a *UploadStickerFileArgs) UserId(userId int) {
	a.withInt("user_id", userId)
}
func (a *UploadStickerFileArgs) StickerFile(fileName string, fileData io.Reader) {
	a.withFile("png_sticker", fileName, fileData)
}


type AddStickerToSetArgs struct {
	_BasicArgs
}
func (a *AddStickerToSetArgs) UserId(userId int) {
	a.withInt("user_id", userId)
}
func (a *AddStickerToSetArgs) Name(name string) {
	a.withString("name", name)
}
func (a *AddStickerToSetArgs) StickerFile(fileName string, fileData io.Reader) {
	a.withFile("png_sticker", fileName, fileData)
}
func (a *AddStickerToSetArgs) StickerId(fileId string) {
	a.withString("png_sticker", fileId)
}
func (a *AddStickerToSetArgs) Emojis(emojis string) {
	a.withString("emojis", emojis)
}
func (a *AddStickerToSetArgs) MaskPosition(point string, x, y, scale float64) {
	mask := map[string]interface{} {
		"point": point,
		"x_shift": x,
		"y_shift": y,
		"scale": scale,
	}
	a.withJson("mask_position", mask)
}


type CreateNewStickerSetArgs struct {
	AddStickerToSetArgs
}
func (a *CreateNewStickerSetArgs) Title(title string) {
	a.withString("title", title)
}
func (a *CreateNewStickerSetArgs) ContainsMasks() {
	a.withBool("contains_masks", true)
}


type SetStickerPositionInSetArgs struct {
	_BasicArgs
}
func (a *SetStickerPositionInSetArgs) StickerId(fileId string) {
	a.withString("sticker", fileId)
}
func (a *SetStickerPositionInSetArgs) Position(position int) {
	a.withInt("position", position)
}


type DeleteStickerFromSetArgs struct {
	_BasicArgs
}
func (p *DeleteStickerFromSetArgs) StickerId(fileId string) {
	p.withString("sticker", fileId)
}
