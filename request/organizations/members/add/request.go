package add

import (
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"strings"
)

type Request struct {
	UserName  null.String
	UserEmail null.String
}

type Key string

const (
	keyUserName  Key = "username"
	keyUserEmail Key = "email"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{}

	if req.UserName.Valid {
		parts[keyUserName] = strings.NewReader(req.UserName.String)
	} else {
		parts[keyUserEmail] = strings.NewReader(req.UserEmail.String)
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
