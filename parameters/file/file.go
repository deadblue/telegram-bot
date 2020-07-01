package file

import (
	"bytes"
	"io"
	"os"
)

type File struct {
	name string
	size int64
	data io.Reader
}

func (f *File) Id() string {
	return ""
}
func (f *File) File() (name string, size int64, data io.Reader) {
	return f.name, f.size, f.data
}

func FromStream(name string, size int64, data io.Reader) *File {
	return &File{
		name: name,
		size: size,
		data: data,
	}
}

func FromBytes(name string, data []byte) *File {
	return &File{
		name: name,
		size: int64(len(data)),
		data: bytes.NewReader(data),
	}
}

func FromLocal(path string) (file *File, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	info, err := f.Stat()
	if err != nil {
		return
	}
	file = &File{
		name: info.Name(),
		size: info.Size(),
		data: f,
	}
	return
}
