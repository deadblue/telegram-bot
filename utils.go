package telegroid

import "io"

func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}
func toLower(ch rune) rune {
	if isUpper(ch) {
		return ch - ('A' - 'a')
	} else {
		return ch
	}
}
func toMethodName(name string) string {
	if len(name) == 0 {
		return ""
	} else {
		runes := []rune(name)
		runes[0] = toLower(runes[0])
		return string(runes)
	}
}

func quietlyClose(c io.Closer) {
	_ = c.Close()
}
