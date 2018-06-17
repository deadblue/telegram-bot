package telegroid

import (
	"reflect"
	"bytes"
	"mime/multipart"
	"net/url"
	"strconv"
	"encoding/json"
	"io"
	"io/ioutil"
)

const (
	TagParameter = "parameter"
)

var (
	InputFileType = reflect.TypeOf(InputFile{})
)

type holder struct {
	Parameters map[string]reflect.Value
	Files      map[string]InputFile
}

type encoder struct {
	IsMultipart     bool
	MultipartBuffer *bytes.Buffer
	MultipartWriter *multipart.Writer
	UrlencodeForm   url.Values
}

func (e *encoder) AddString(name, value string) {
	if e.IsMultipart {
		e.MultipartWriter.WriteField(name, value)
	} else {
		e.UrlencodeForm.Set(name, value)
	}
}

func (e *encoder) AddFile(name, filename string, file io.Reader) {
	if e.IsMultipart {
		w, _ := e.MultipartWriter.CreateFormFile(name, filename)
		io.Copy(w, file)
	}
}

func (e *encoder) ContentType() string {
	if e.IsMultipart {
		return e.MultipartWriter.FormDataContentType()
	} else {
		return "application/x-www-form-urlencoded"
	}
}

func (e *encoder) Finish() []byte {
	if e.IsMultipart {
		e.MultipartWriter.Close()
		return e.MultipartBuffer.Bytes()
	} else {
		return []byte(e.UrlencodeForm.Encode())
	}
}

func newHolder() holder {
	return holder{
		Parameters: make(map[string]reflect.Value),
		Files:      make(map[string]InputFile),
	}
}

func newEncoder(isMultipart bool) encoder {
	encoder := encoder{IsMultipart: isMultipart}
	if isMultipart {
		encoder.MultipartBuffer = &bytes.Buffer{}
		encoder.MultipartWriter = multipart.NewWriter(encoder.MultipartBuffer)
	} else {
		encoder.UrlencodeForm = url.Values{}
	}
	return encoder
}

func encodeArguments(args interface{}) (string, io.ReadCloser) {
	// scan parameters
	holder := newHolder()
	scanParameters(holder, args)
	// create encoder
	encoder := newEncoder( len(holder.Files) != 0 )
	// add files
	for pn, pv := range holder.Files {
		encoder.AddFile(pn, pv.Name, pv.Reader)
	}
	// add parameters
	for pn, pv := range holder.Parameters {
		for pv.Kind() == reflect.Ptr {
			pv = pv.Elem()
		}
		switch pv.Kind() {
		case reflect.Bool:
			boolVal := pv.Bool()
			if boolVal {
				encoder.AddString(pn, strconv.FormatBool(boolVal))
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal := pv.Int()
			if intVal != 0 {
				encoder.AddString(pn, strconv.FormatInt(intVal, 10))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal := pv.Uint()
			if uintVal != 0 {
				encoder.AddString(pn, strconv.FormatUint(uintVal, 10))
			}
		case reflect.Float32, reflect.Float64:
			encoder.AddString(pn, strconv.FormatFloat(pv.Float(), 'f', -1, 64))
		case reflect.String:
			strVal := pv.String()
			if len(strVal) > 0 {
				encoder.AddString(pn, strVal)
			}
		case reflect.Map, reflect.Slice, reflect.Interface:
			if !pv.IsNil() {
				jv, err := json.Marshal(pv.Interface())
				if err != nil {
					encoder.AddString(pn, string(jv))
				}
			}
		case reflect.Array, reflect.Struct:
			jv, err := json.Marshal(pv.Interface())
			if err == nil {
				encoder.AddString(pn, string(jv))
			}
		}
	}
	data := encoder.Finish()
	return encoder.ContentType(), ioutil.NopCloser(bytes.NewReader(data))
}

func scanParameters(holder holder, entity interface{}) {
	rv, rt := reflect.ValueOf(entity), reflect.TypeOf(entity)
	if rv.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rv.NumField(); i ++ {
		fv, ft := rv.Field(i), rt.Field(i)
		if ft.Anonymous {
			// for embedded struct, deeply scan
			scanParameters(holder, fv.Interface())
		} else {
			pn := ft.Tag.Get(TagParameter)
			if len(pn) == 0 { continue }
			if ft.Type == InputFileType {
				file := fv.Interface().(InputFile)
				if file.Reader != nil {
					holder.Files[pn] = file
				}
			} else {
				holder.Parameters[pn] = fv
			}
		}
	}
}
