package telegroid

import (
	"io"
	"os"
	"bytes"
	"path"
)

type InputFile struct {
	Name   string
	Reader io.Reader
}

func FromFile(filepath string) (value *InputFile, err error) {
	file, err := os.Open(filepath)
	if err == nil {
		value = FromReader(path.Base(file.Name()), file)
	}
	return
}

func FromData(name string, data []byte) *InputFile {
	return &InputFile{
		Name: name,
		Reader: bytes.NewReader(data),
	}
}

func FromReader(name string, reader io.Reader) *InputFile {
	return &InputFile{
		Name: name,
		Reader: reader,
	}
}
