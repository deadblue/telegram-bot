package arguments

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// InputFile interface that presents name and data reader
type InputFile interface {

	// The base name of the file
	Name() string

	// The reader to read the file data,
	// will be closed by the code who use it.
	Data() io.ReadCloser

}


type dataInputFile struct {
	name string
	data io.ReadCloser
}
func (f *dataInputFile) Name() string {
	return f.name
}
func (f *dataInputFile) Data() io.ReadCloser {
	return f.data
}

// Create `InputFile` instance by pesudo name and memory data
func InputFileFromData(name string, data []byte) InputFile {
	return &dataInputFile{
		name: name,
		data: ioutil.NopCloser( bytes.NewReader(data) ),
	}
}


type localInputFile struct {
	name string
	data io.ReadCloser
}
func (f *localInputFile) Name() string {
	return f.name
}
func (f *localInputFile) Data() io.ReadCloser {
	return f.data
}

// Create `InputFile` instance by local file
func InputFileFromLocal(localFile string) (f InputFile, err error) {
	file, err := os.Open(localFile)
	if err != nil {
		return
	}
	return &localInputFile{
		name: file.Name(),
		data: file,
	}, nil
}
