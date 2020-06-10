package teams

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type CreateRequest struct {
	TeamName string
}

func (req CreateRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"team": strings.NewReader(req.TeamName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req CreateRequest) Verify() error {

	if req.TeamName == "" {
		return fmt.Errorf("team name must not be empty")
	}

	return nil
}
