package request

import (
	"bytes"
	"github.com/jmatsu/dpg/util"
	"io"
)

type Body interface {
	IoReaderMap() (*map[string]io.Reader, error)
}

func ToMultiFormPart(body Body) (bytes.Buffer, string, error) {
	ioMap, err := body.IoReaderMap()

	if err != nil {
		return bytes.Buffer{}, "", err
	}

	return util.Buffering(*ioMap)
}
