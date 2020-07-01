package file

import (
	"io"
)

type Id string

func (i Id) Id() string {
	return string(i)
}
func (i Id) File() (name string, size int64, data io.Reader) {
	return "", 0, nil
}

type Url string

func (u Url) Id() string {
	return string(u)
}
func (u Url) File() (name string, size int64, data io.Reader) {
	return "", 0, nil
}
