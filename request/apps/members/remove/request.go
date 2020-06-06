package remove

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
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

func (req Request) Verify() error {
	if len(req.UserNamesOrEmails) == 0 {
		return fmt.Errorf("the number of removees must be greater than 0")
	}

	return nil
}
