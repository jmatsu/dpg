package remove

import (
	"io"
	"strings"
	"github.com/jmatsu/dpg/util"
)

type Request struct {
	UserNamesOrEmails []string
}

type Key string

const (
	keyRemovees Key = "users"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyRemovees: strings.NewReader(strings.Join(req.UserNamesOrEmails, ",")),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
