package parameters

import "github.com/deadblue/telegroid/types"

type FormattedText struct {
	text string
	mode types.ParseMode
}

func PlainText(text string) FormattedText {
	return FormattedText{
		text: text,
		mode: types.ModePlain,
	}
}

func HTMLText(text string) FormattedText {
	return FormattedText{
		text: text,
		mode: types.ModeHtml,
	}
}

func MarkdownText(text string) FormattedText {
	return FormattedText{
		text: text,
		mode: types.ModeMarkdown,
	}
}

func MarkdownV2Text(text string) FormattedText {
	return FormattedText{
		text: text,
		mode: types.ModeMarkdownV2,
	}
}
