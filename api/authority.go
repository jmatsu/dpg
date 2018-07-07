package api

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
	"bytes"
)

type Authority struct {
	Token string
}

type Key string

const (
	Token Key = "token"
)

func (authority Authority) ioReaderMap() (*map[string]io.Reader, error) {
	m := map[Key]io.Reader{
		Token: strings.NewReader(authority.Token),
	}

	out, err := util.StringifyKeys(m)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (authority Authority) Merge(m map[string]io.Reader) (out map[string]io.Reader, err error) {
	m2, err := authority.ioReaderMap()

	if err != nil {
		return
	}

	out = util.Merge(m, *m2)

	return
}

func (authority Authority) MultiPartForm(reqIOMap *map[string]io.Reader) (data bytes.Buffer, contentType string, err error) {
	ioMap, err := authority.Merge(*reqIOMap)

	if err != nil {
		return data, "", err
	}

	data, contentType, err = util.Buffering(ioMap)

	if err != nil {
		return data, "", err
	}

	return data, contentType, nil
}