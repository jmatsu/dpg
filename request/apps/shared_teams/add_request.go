package shared_teams

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type AddRequest struct {
	SharedTeamName string
}

func (req AddRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"team": strings.NewReader(req.SharedTeamName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req AddRequest) Verify() error {
	if req.SharedTeamName == "" {
		return fmt.Errorf("shared team name must be present")
	}

	return nil
}
