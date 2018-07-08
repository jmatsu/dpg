package add

import (
	"io"
	"strings"
	"github.com/jmatsu/dpg/util"
)

type Request struct {
	TeamName string
}

type Key string

const (
	keyTeamName Key = "team"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyTeamName: strings.NewReader(req.TeamName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
