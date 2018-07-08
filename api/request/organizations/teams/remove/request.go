package remove

import (
	"io"
	"github.com/jmatsu/dpg/util"
)

type Request struct{}

type Key string

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
