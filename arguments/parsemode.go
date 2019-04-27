package arguments

// Switch parse mode for text/caption
type ParseModeSelector interface {

	// Set parse mode to Markdown
	Markdown()

	// Set parse mode to HTML
	HTML()
}

type implParseModeSelector struct {
	form *_Form
}

func (m *implParseModeSelector) Markdown() {
	m.form.WithString("parse_mode", parseModeMarkdown)
}
func (m *implParseModeSelector) HTML() {
	m.form.WithString("parse_mode", parseModeHTML)
}
