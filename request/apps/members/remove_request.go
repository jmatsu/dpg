package members

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type RemoveRequest struct {
	UserNamesOrEmails []string
}

func (req RemoveRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"users": strings.NewReader(strings.Join(req.UserNamesOrEmails, ",")),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req RemoveRequest) Verify() error {
	if len(req.UserNamesOrEmails) == 0 {
		return fmt.Errorf("the number of removees must be greater than 0")
	}

	return nil
}
