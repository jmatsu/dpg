package organization_members

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type AddRequest struct {
	UserName string
}

func (req AddRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"user": strings.NewReader(req.UserName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req AddRequest) Verify() error {
	if req.UserName == "" {
		return fmt.Errorf("user name must not be empty")
	}

	return nil
}
