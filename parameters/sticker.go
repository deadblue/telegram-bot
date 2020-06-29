package parameters

import (
	"github.com/deadblue/telegram-bot/types"
)

type SendStickerParams struct {
	baseSendParams
}

func (p *SendStickerParams) Sticker(file InputFile) {
	p.setFile("sticker", file)
}

type GetStickerSetParams struct {
	implApiParameters
}

func (p *GetStickerSetParams) Name(name string) {
	p.set("name", name)
}

type UploadStickerFileParams struct {
	implApiParameters
}

func (p *UploadStickerFileParams) UserId(userId int) {
	p.setInt("user_id", userId)
}
func (p *UploadStickerFileParams) PngSticker(file InputFile) {
	p.setFile("png_sticker", file)
}

type AddStickerToSetParams struct {
	GetStickerSetParams
}

func (p *AddStickerToSetParams) UserId(userId int) {
	p.setInt("user_id", userId)
}
func (p *AddStickerToSetParams) Emojis(emojis string) {
	p.set("emojis", emojis)
}
func (p *AddStickerToSetParams) PngSticker(file InputFile) {
	p.setFile("png_sticker", file)
}
func (p *AddStickerToSetParams) TgsSticker(file InputFile) {
	p.setFile("tgs_sticker", file)
}
func (p *AddStickerToSetParams) MaskPosition(position *types.MaskPosition) {
	p.setJson("mask_position", position)
}

type CreateNewStickerSetParams struct {
	AddStickerToSetParams
}

func (p *CreateNewStickerSetParams) Title(title string) {
	p.set("title", title)
}
func (p *CreateNewStickerSetParams) ContainsMasks() {
	p.setBool("contains_masks", true)
}

type DeleteStickerFromSetParams struct {
	implApiParameters
}

func (p *DeleteStickerFromSetParams) Sticker(fileId string) {
	p.set("sticker", fileId)
}

type SetStickerPositionInSetParams struct {
	DeleteStickerFromSetParams
}

func (p *SetStickerPositionInSetParams) Position(position int) {
	p.setInt("position", position)
}

type SetStickerSetThumbParams struct {
	GetStickerSetParams
}

func (p *SetStickerSetThumbParams) UserId(userId int) {
	p.setInt("user_id", userId)
}
func (p *SetStickerSetThumbParams) Thumb(file InputFile) {
	p.setFile("thumb", file)
}
