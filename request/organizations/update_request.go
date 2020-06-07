package organizations

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type UpdateRequest struct {
	Description string
}

func (req UpdateRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"description": strings.NewReader(req.Description),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
