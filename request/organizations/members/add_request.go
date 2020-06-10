package members

import (
	"fmt"
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

func (req AddRequest) Verify() error {
	if req.UserEmail.Valid && req.UserName.Valid {
		return fmt.Errorf("user email and name cannot be specified at once")
	}

	if !req.UserName.Valid {
		return fmt.Errorf("one of user email or name is required")
	}

	return nil
}
