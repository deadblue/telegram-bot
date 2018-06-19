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
	InputFileOrStringType = reflect.TypeOf(InputFileOrString{})
)

type holder struct {
	Parameters map[string]string
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
		Parameters: make(map[string]string),
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
			if ft.Type == InputFileOrStringType {
				pv := fv.Interface().(InputFileOrString)
				if pv.FileValue != nil {
					holder.Files[pn] = *pv.FileValue
				} else {
					holder.Parameters[pn] = pv.StringValue
				}
			} else if ft.Type == InputFileType {
				file := fv.Interface().(InputFile)
				if file.Reader != nil {
					holder.Files[pn] = file
				}
			} else {
				holder.Parameters[pn] = getStringValue(fv)
			}
		}
	}
}

func getStringValue(value reflect.Value) string {
	result := ""
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Bool:
		boolVal := value.Bool()
		if boolVal {
			result = strconv.FormatBool(boolVal)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal := value.Int()
		if intVal != 0 {
			result = strconv.FormatInt(intVal, 10)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal := value.Uint()
		if uintVal != 0 {
			result = strconv.FormatUint(uintVal, 10)
		}
	case reflect.Float32, reflect.Float64:
		result = strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.String:
		result = value.String()
	case reflect.Map, reflect.Slice, reflect.Interface:
		if !value.IsNil() {
			jv, err := json.Marshal( value.Interface() )
			if err != nil {
				result = string(jv)
			}
		}
	case reflect.Array, reflect.Struct:
		jv, err := json.Marshal(value.Interface())
		if err == nil {
			result = string(jv)
		}
	}
	return result
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
		if len(pv) != 0 {
			encoder.AddString(pn, pv)
		}
	}
	data := encoder.Finish()
	return encoder.ContentType(), ioutil.NopCloser(bytes.NewReader(data))
}
