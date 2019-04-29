package arguments

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// InputFile interface that hold the file name and content reader
type InputFile interface {

	// Base name of the file
	Name() string

	// Reader to read the file data, will be
	// closed by the code who use it.
	Data() io.ReadCloser
}

type implInputFile struct {
	name string
	data io.ReadCloser
}

func (f *implInputFile) Name() string {
	return f.name
}
func (f *implInputFile) Data() io.ReadCloser {
	return f.data
}

// Create `InputFile` instance by pesudo name and memory data
func InputFileFromData(name string, data []byte) InputFile {
	return &implInputFile{
		name: name,
		data: ioutil.NopCloser(bytes.NewReader(data)),
	}
}

// Create `InputFile` instance by local file
func InputFileFromLocal(localFile string) (f InputFile, err error) {
	file, err := os.Open(localFile)
	if err != nil {
		return
	}
	return &implInputFile{
		name: file.Name(),
		data: file,
	}, nil
}
