package shared_team_relations

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type CreateRequest struct {
	SharedTeamName string
}

func (req CreateRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"team": strings.NewReader(req.SharedTeamName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req CreateRequest) Verify() error {
	if req.SharedTeamName == "" {
		return fmt.Errorf("shared team name must be present")
	}

	return nil
}
