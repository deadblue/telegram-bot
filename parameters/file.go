package parameters

import (
	"bytes"
	"io"
	"os"
)

type InputFile struct {
	// file Id on telegram server.
	fileIdOrUrl string
	// file data from local.
	name string
	size int64
	data io.Reader
}

func FileFromId(fileId string) *InputFile {
	return &InputFile{
		fileIdOrUrl: fileId,
	}
}

func FileFromUrl(url string) *InputFile {
	return &InputFile{
		fileIdOrUrl: url,
	}
}

func FileFromData(name string, data []byte) *InputFile {
	return &InputFile{
		fileIdOrUrl: "",
		name:        name,
		size:        int64(len(data)),
		data:        bytes.NewReader(data),
	}
}

func FileFromLocal(path string) (*InputFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &InputFile{
		fileIdOrUrl: "",
		name:        info.Name(),
		size:        info.Size(),
		data:        file,
	}, nil
}
