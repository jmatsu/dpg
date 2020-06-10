package team_members

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type CreateRequest struct {
	UserName string
}

func (req CreateRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"user": strings.NewReader(req.UserName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req CreateRequest) Verify() error {
	if req.UserName == "" {
		return fmt.Errorf("user name must be present")
	}

	return nil
}