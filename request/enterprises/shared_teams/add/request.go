package add

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type Request struct {
	SharedTeamName string
	Description    string
}

type Key string

const (
	keyTeamName    Key = "name"
	keyDescription Key = "description"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyTeamName:    strings.NewReader(req.SharedTeamName),
		keyDescription: strings.NewReader(req.Description),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
