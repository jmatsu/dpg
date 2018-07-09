package add

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type Request struct {
	UserName string
}

type Key string

const (
	keyUserName Key = "user"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyUserName: strings.NewReader(req.UserName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
