package members

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type AddRequest struct {
	UserNamesOrEmails []string
	DeveloperRole     bool
}

func (req AddRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"users": strings.NewReader(strings.Join(req.UserNamesOrEmails, ",")),
	}

	if req.DeveloperRole {
		parts["role"] = strings.NewReader("1")
	} else {
		parts["role"] = strings.NewReader("2")
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req AddRequest) Verify() error {
	if len(req.UserNamesOrEmails) == 0 {
		return fmt.Errorf("the number of invitees must be greater than 0")
	}

	return nil
}
