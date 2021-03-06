package add

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type Request struct {
	SharedTeamName string
}

type Key string

const (
	keyTeamName Key = "team"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyTeamName: strings.NewReader(req.SharedTeamName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
