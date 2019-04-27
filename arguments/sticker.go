package arguments

type SendStickerArgs struct {
	_CommonSendArgs
}

func (a *SendStickerArgs) StickerId(stickerId string) {
	a.getForm().WithString("sticker", stickerId)
}
func (a *SendStickerArgs) StickerFile(file InputFile) {
	a.getForm().WithFile("sticker", file)
}

type GetStickerSetArgs struct {
	_BasicArgs
}

func (a *GetStickerSetArgs) Name(name string) {
	a.getForm().WithString("name", name)
}

type UploadStickerFileArgs struct {
	_BasicArgs
}

func (a *UploadStickerFileArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}
func (a *UploadStickerFileArgs) StickerFile(file InputFile) {
	a.getForm().WithFile("png_sticker", file)
}

type AddStickerToSetArgs struct {
	_BasicArgs
}

func (a *AddStickerToSetArgs) UserId(userId int) {
	a.getForm().WithInt("user_id", userId)
}
func (a *AddStickerToSetArgs) Name(name string) {
	a.getForm().WithString("name", name)
}
func (a *AddStickerToSetArgs) StickerId(fileId string) {
	a.getForm().WithString("png_sticker", fileId)
}
func (a *AddStickerToSetArgs) StickerFile(file InputFile) {
	a.getForm().WithFile("png_sticker", file)
}
func (a *AddStickerToSetArgs) Emojis(emojis string) {
	a.getForm().WithString("emojis", emojis)
}
func (a *AddStickerToSetArgs) MaskPosition(point string, x, y, scale float64) {
	mask := map[string]interface{}{
		"point":   point,
		"x_shift": x,
		"y_shift": y,
		"scale":   scale,
	}
	a.getForm().WithJson("mask_position", mask)
}


type CreateNewStickerSetArgs struct {
	AddStickerToSetArgs
}

func (a *CreateNewStickerSetArgs) Title(title string) {
	a.getForm().WithString("title", title)
}
func (a *CreateNewStickerSetArgs) ContainsMasks() {
	a.getForm().WithBool("contains_masks", true)
}

type SetStickerPositionInSetArgs struct {
	_BasicArgs
}

func (a *SetStickerPositionInSetArgs) StickerId(fileId string) {
	a.getForm().WithString("sticker", fileId)
}
func (a *SetStickerPositionInSetArgs) Position(position int) {
	a.getForm().WithInt("position", position)
}

type DeleteStickerFromSetArgs struct {
	_BasicArgs
}

func (a *DeleteStickerFromSetArgs) StickerId(fileId string) {
	a.getForm().WithString("sticker", fileId)
}
