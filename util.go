package telegroid

import "strconv"

type MessageBuilder struct {
	args AbstractSendArguments
}
func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{ args: AbstractSendArguments{} }
}
func (b *MessageBuilder) ChatId(chatId int) *MessageBuilder {
	b.args.ChatId = strconv.Itoa(chatId)
	return b
}
func (b *MessageBuilder) ChannelUser(channelUser string) *MessageBuilder {
	b.args.ChatId = channelUser
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
			row[j] = KeyboardButton{ Text: buttons[i][j] }
		}
		rows[i] = row
	}
	b.args.ReplyMarkup = ReplyKeyboardMarkup{ Keyboard:rows, OneTimeKeyboard:oneTime }
	return b
}
func (b *MessageBuilder) RemoveKeyboard() *MessageBuilder {
	b.args.ReplyMarkup = ReplyKeyboardRemove{ RemoveKeyboard:true }
	return b
}
func (b *MessageBuilder) ForceReply() *MessageBuilder {
	b.args.ReplyMarkup = ForceReply{ ForceReply:true }
	return b
}

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
	tb.args.ParseMode = ""
	tb.args.Text = text
	return tb
}
func (tb *textMessageBuilder) Markdown(markdown string) *textMessageBuilder {
	tb.args.ParseMode = ParseModeMarkdown
	tb.args.Text = markdown
	return tb
}
func (tb *textMessageBuilder) Done() SendMessageArguments {
	return tb.args
}

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
	pb.args.Photo = InputFileOrString{ StringValue:fileId }
	return pb
}
func (pb *photoMessaageBuilder) Url(url string) *photoMessaageBuilder {
	pb.args.Photo = InputFileOrString{ StringValue:url }
	return pb
}
func (pb *photoMessaageBuilder) LocalFile(filepath string) *photoMessaageBuilder {
	file, err := FromFile(filepath)
	if err == nil {
		pb.args.Photo = InputFileOrString{ FileValue:&file }
	}
	return pb
}
func (pb *photoMessaageBuilder) Caption(caption string) *photoMessaageBuilder {
	pb.args.ParseMode = ""
	pb.args.Caption = caption
	return pb
}
func (pb *photoMessaageBuilder) MarkdownCaption(markdown string) *photoMessaageBuilder {
	pb.args.ParseMode = ParseModeMarkdown
	pb.args.Caption = markdown
	return pb
}
func (pb *photoMessaageBuilder) Done() SendPhotoArguments {
	return pb.args
}

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
	sb.args.Sticker = InputFileOrString{ StringValue:fileId }
	return sb
}
func (sb *stickerMessageBuilder) Done() SendStickerArguments {
	return sb.args
}
