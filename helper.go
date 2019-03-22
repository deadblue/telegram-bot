package telegroid

import (
	"fmt"
	"strconv"
)

type MessageBuilder struct {
	args AbstractSendArguments
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{args: AbstractSendArguments{}}
}
func (b *MessageBuilder) ChatId(chatId int) *MessageBuilder {
	b.args.ChatId = strconv.Itoa(chatId)
	return b
}
func (b *MessageBuilder) Channel(channel string) *MessageBuilder {
	b.args.ChatId = fmt.Sprintf("@%s", channel)
	return b
}
func (b *MessageBuilder) DisbaleNotification() *MessageBuilder {
	b.args.DisableNotificate = true
	return b
}
func (b *MessageBuilder) ReplyTo(messageId int) *MessageBuilder {
	b.args.ReplyToMessageId = messageId
	return b
}
func (b *MessageBuilder) Keyboard(buttons [][]string, oneTime bool) *MessageBuilder {
	rowCount := len(buttons)
	rows := make([][]KeyboardButton, rowCount)
	for i := 0; i < rowCount; i++ {
		colCount := len(buttons[i])
		row := make([]KeyboardButton, colCount)
		for j := 0; j < colCount; j++ {
			row[j] = KeyboardButton{Text: buttons[i][j]}
		}
		rows[i] = row
	}
	b.args.ReplyMarkup = ReplyKeyboardMarkup{Keyboard: rows, OneTimeKeyboard: oneTime}
	return b
}
func (b *MessageBuilder) RemoveKeyboard() *MessageBuilder {
	b.args.ReplyMarkup = ReplyKeyboardRemove{RemoveKeyboard: true}
	return b
}
func (b *MessageBuilder) InlineKeyboard(buttons [][]string) *MessageBuilder {
	rowCount := len(buttons)
	rows := make([][]InlineKeyboardButton, rowCount)
	for i := 0; i < rowCount; i++ {
		colCount := len(buttons[i])
		row := make([]InlineKeyboardButton, colCount)
		for j := 0; j < colCount; j++ {
			callbackData := fmt.Sprintf("%d_%d", i, j)
			row[j] = InlineKeyboardButton{
				Text:         buttons[i][j],
				CallbackData: callbackData,
			}
		}
		rows[i] = row
	}
	b.args.ReplyMarkup = InlineKeyboardMarkup{InlineKeyboard: rows}
	return b
}
func (b *MessageBuilder) ForceReply() *MessageBuilder {
	b.args.ReplyMarkup = ForceReply{ForceReply: true}
	return b
}

// Text message builder
type textMessageBuilder struct {
	args SendMessageArguments
}

func (b *MessageBuilder) ForText() *textMessageBuilder {
	return &textMessageBuilder{
		args: SendMessageArguments{
			AbstractSendArguments: b.args,
		},
	}
}
func (tb *textMessageBuilder) Text(text string) *textMessageBuilder {
	tb.args.ParseMode, tb.args.Text = ParseModeNormal, text
	return tb
}
func (tb *textMessageBuilder) Markdown(markdown string) *textMessageBuilder {
	tb.args.ParseMode, tb.args.Text = ParseModeMarkdown, markdown
	return tb
}
func (tb *textMessageBuilder) Done() SendMessageArguments {
	return tb.args
}

// Photo message builder
type photoMessaageBuilder struct {
	args SendPhotoArguments
}

func (b *MessageBuilder) ForPhoto() *photoMessaageBuilder {
	return &photoMessaageBuilder{
		args: SendPhotoArguments{
			AbstractSendArguments: b.args,
		},
	}
}
func (pb *photoMessaageBuilder) FileId(fileId string) *photoMessaageBuilder {
	pb.args.Photo = InputFileOrString{StringValue: fileId}
	return pb
}
func (pb *photoMessaageBuilder) Url(url string) *photoMessaageBuilder {
	pb.args.Photo = InputFileOrString{StringValue: url}
	return pb
}
func (pb *photoMessaageBuilder) LocalFile(filepath string) *photoMessaageBuilder {
	file, err := FromFile(filepath)
	if err == nil {
		pb.args.Photo = InputFileOrString{FileValue: file}
	}
	return pb
}
func (pb *photoMessaageBuilder) Caption(caption string) *photoMessaageBuilder {
	pb.args.ParseMode, pb.args.Caption = ParseModeNormal, caption
	return pb
}
func (pb *photoMessaageBuilder) CaptionMarkdown(markdown string) *photoMessaageBuilder {
	pb.args.ParseMode, pb.args.Caption = ParseModeMarkdown, markdown
	return pb
}
func (pb *photoMessaageBuilder) Done() SendPhotoArguments {
	return pb.args
}

// Sticker message builder
type stickerMessageBuilder struct {
	args SendStickerArguments
}

func (b *MessageBuilder) ForSticker() *stickerMessageBuilder {
	return &stickerMessageBuilder{
		args: SendStickerArguments{
			AbstractSendArguments: b.args,
		},
	}
}
func (sb *stickerMessageBuilder) FileId(fileId string) *stickerMessageBuilder {
	sb.args.Sticker = InputFileOrString{StringValue: fileId}
	return sb
}
func (sb *stickerMessageBuilder) Done() SendStickerArguments {
	return sb.args
}

// Document message builder
type documentMessageBuilder struct {
	args SendDocumentArguments
}

func (b *MessageBuilder) ForDocument() *documentMessageBuilder {
	return &documentMessageBuilder{
		args: SendDocumentArguments{AbstractSendArguments: b.args},
	}
}
func (db *documentMessageBuilder) FileId(fileId string) *documentMessageBuilder {
	db.args.Document = InputFileOrString{StringValue: fileId}
	return db
}
func (db *documentMessageBuilder) LocalFile(filepath string) *documentMessageBuilder {
	file, err := FromFile(filepath)
	if err == nil {
		db.args.Document = InputFileOrString{FileValue: file}
	}
	return db
}
func (db *documentMessageBuilder) Caption(caption string) *documentMessageBuilder {
	db.args.ParseMode, db.args.Caption = ParseModeNormal, caption
	return db
}
func (db *documentMessageBuilder) CaptionMarkdown(markdown string) *documentMessageBuilder {
	db.args.ParseMode, db.args.Caption = ParseModeMarkdown, markdown
	return db
}
func (db *documentMessageBuilder) Done() SendDocumentArguments {
	return db.args
}

// Audio message builder
type audioMessageBuilder struct {
	args SendAudioArguments
}

func (b *MessageBuilder) ForAudio() *audioMessageBuilder {
	return &audioMessageBuilder{
		args: SendAudioArguments{AbstractSendArguments: b.args},
	}
}
func (ab *audioMessageBuilder) FileId(fileId string) *audioMessageBuilder {
	ab.args.Audio = InputFileOrString{StringValue: fileId}
	return ab
}
func (ab *audioMessageBuilder) LocalFile(filepath string) *audioMessageBuilder {
	file, err := FromFile(filepath)
	if err == nil {
		ab.args.Audio = InputFileOrString{FileValue: file}
	}
	return ab
}
func (ab *audioMessageBuilder) Caption(caption string) *audioMessageBuilder {
	ab.args.ParseMode, ab.args.Caption = ParseModeNormal, caption
	return ab
}
func (ab *audioMessageBuilder) CaptionMarkdown(markdown string) *audioMessageBuilder {
	ab.args.ParseMode, ab.args.Caption = ParseModeMarkdown, markdown
	return ab
}
func (ab *audioMessageBuilder) Performer(performer string) *audioMessageBuilder {
	ab.args.Performer = performer
	return ab
}
func (ab *audioMessageBuilder) Title(title string) *audioMessageBuilder {
	ab.args.Title = title
	return ab
}
func (ab *audioMessageBuilder) Done() SendAudioArguments {
	return ab.args
}

// Video message builder
type videoMessageBuilder struct {
	args SendVideoArguments
}

func (b *MessageBuilder) ForVideo() *videoMessageBuilder {
	return &videoMessageBuilder{
		args: SendVideoArguments{AbstractSendArguments: b.args},
	}
}
func (vb *videoMessageBuilder) FileId(fileId string) *videoMessageBuilder {
	vb.args.Video = InputFileOrString{StringValue: fileId}
	return vb
}
func (vb *videoMessageBuilder) LocalFile(filepath string) *videoMessageBuilder {
	file, err := FromFile(filepath)
	if err == nil {
		vb.args.Video = InputFileOrString{FileValue: file}
	}
	return vb
}
func (vb *videoMessageBuilder) Caption(caption string) *videoMessageBuilder {
	vb.args.ParseMode, vb.args.Caption = ParseModeNormal, caption
	return vb
}
func (vb *videoMessageBuilder) CaptionMarkdown(markdown string) *videoMessageBuilder {
	vb.args.ParseMode, vb.args.Caption = ParseModeMarkdown, markdown
	return vb
}
func (vb *videoMessageBuilder) Size(width, height int) *videoMessageBuilder {
	vb.args.Width, vb.args.Height = width, height
	return vb
}
func (vb *videoMessageBuilder) Done() SendVideoArguments {
	return vb.args
}
