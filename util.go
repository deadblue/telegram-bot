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
	tb.args.Text = text
	return tb
}
func (tb *textMessageBuilder) ParseMode(mode ParseMode) *textMessageBuilder {
	tb.args.ParseMode = mode
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
func (pb *photoMessaageBuilder) PhotoFileId(fileId string) *photoMessaageBuilder {
	pb.args.Photo = InputFileOrString{ StringValue:fileId }
	return pb
}
func (pb *photoMessaageBuilder) PhotoUrl(url string) *photoMessaageBuilder {
	pb.args.Photo = InputFileOrString{ StringValue:url }
	return pb
}
func (pb *photoMessaageBuilder) PhotoFile(filepath string) *photoMessaageBuilder {
	file, err := FromFile(filepath)
	if err == nil {
		pb.args.Photo = InputFileOrString{ FileValue:&file }
	}
	return pb
}
func (pb *photoMessaageBuilder) Caption(caption string) *photoMessaageBuilder {
	pb.args.Caption = caption
	return pb
}
func (pb *photoMessaageBuilder) Done() SendPhotoArguments {
	return pb.args
}
