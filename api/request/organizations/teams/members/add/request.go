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
	keyUser Key = "user"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	var userKey string

	if req.UserName.Valid {
		userKey = req.UserName.String
	} else {
		userKey = req.UserEmail.String
	}

	parts := map[Key]io.Reader{
		keyUser: strings.NewReader(userKey),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
