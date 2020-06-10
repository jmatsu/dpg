package shared_teams

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"strings"
)

type CreateRequest struct {
	SharedTeamName string
	Description    null.String
}

func (req CreateRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"name":        strings.NewReader(req.SharedTeamName),
	}

	if req.Description.Valid {
		parts["description"] = strings.NewReader(req.Description.String)
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
