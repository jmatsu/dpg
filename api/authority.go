package api

import (
	"bytes"
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"net/url"
	"strings"
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

	out, err := util.StringifyKeysOfReaderMap(m)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (authority Authority) stringMap() (*map[string]string, error) {
	m := map[Key]string{
		Token: authority.Token,
	}

	out, err := util.StringifyKeysAndValues(m)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (authority Authority) mergeIOReaderMap(m map[string]io.Reader) (out map[string]io.Reader, err error) {
	m2, err := authority.ioReaderMap()

	if err != nil {
		return
	}

	out = util.MergeIOReaderMap(m, *m2)

	return
}

func (authority Authority) mergeStringMap(m map[string]string) (out map[string]string, err error) {
	m2, err := authority.stringMap()

	if err != nil {
		return
	}

	out = util.MergeStringMap(m, *m2)

	return
}

func (authority Authority) GetParams(stringMap *map[string]string) (string, error) {
	strMap, err := authority.mergeStringMap(*stringMap)

	if err != nil {
		return "", err
	}

	var slices []string

	for k, v := range strMap {
		slices = append(slices, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}

	return strings.Join(slices, "&"), nil
}

func (authority Authority) MultiPartForm(reqIOMap *map[string]io.Reader) (data bytes.Buffer, contentType string, err error) {
	ioMap, err := authority.mergeIOReaderMap(*reqIOMap)

	if err != nil {
		return data, "", err
	}

	data, contentType, err = util.Buffering(ioMap)

	if err != nil {
		return data, "", err
	}

	return data, contentType, nil
}
