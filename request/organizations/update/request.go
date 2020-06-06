package update

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type Request struct {
	Description string
}

type Key string

const (
	keyDescription Key = "description"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyDescription: strings.NewReader(req.Description),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
