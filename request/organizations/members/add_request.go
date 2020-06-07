package members

import (
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"strings"
)

type AddRequest struct {
	UserName  null.String
	UserEmail null.String
}

func (req AddRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{}

	if req.UserName.Valid {
		parts["username"] = strings.NewReader(req.UserName.String)
	} else {
		parts["email"] = strings.NewReader(req.UserEmail.String)
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
