package util

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"reflect"
)

func Merge(m1, m2 map[string]io.Reader) map[string]io.Reader {
	newMap := map[string]io.Reader{}

	for k, v := range m1 {
		newMap[k] = v
	}

	for k, v := range m2 {
		newMap[k] = v
	}

	return newMap
}

func StringifyKeys(mustBeMap interface{}) (*map[string]io.Reader, error) {
	m := reflect.ValueOf(mustBeMap)

	output := make(map[string]io.Reader)

	for _, e := range m.MapKeys() {
		v := m.MapIndex(e)

		if reader, ok := v.Interface().(io.Reader); ok {
			if e.Kind() != reflect.String {
				return nil, errors.New("the type of keys must be string")
			}

			output[e.String()] = reader
		} else {
			return nil, errors.New("all type of values must be io.Reader")
		}
	}

	return &output, nil
}

func Buffering(readerMap map[string]io.Reader) (data bytes.Buffer, contentType string, failure error) {
	writer := multipart.NewWriter(&data)

	for key, reader := range readerMap {
		func() {
			if closer, ok := reader.(io.Closer); ok {
				defer closer.Close()
			}

			if file, ok := reader.(*os.File); ok {
				if fw, err := writer.CreateFormFile(key, file.Name()); err != nil {
					failure = err
					return
				} else if _, err := io.Copy(fw, reader); err != nil {
					failure = err
					return
				}
			} else {
				if fw, err := writer.CreateFormField(key); err != nil {
					failure = err
					return
				} else if _, err := io.Copy(fw, reader); err != nil {
					failure = err
					return
				}
			}
		}()
	}

	if err := writer.Close(); err != nil {
		return data, "", err
	}

	return data, writer.FormDataContentType(), nil
}
