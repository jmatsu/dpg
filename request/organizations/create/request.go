package create

import (
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"strings"
)

type Request struct {
	OrganizationName string
	Description      null.String
}

type Key string

const (
	keyOrganizationName Key = "name"
	keyDescription      Key = "description"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyOrganizationName: strings.NewReader(req.OrganizationName),
	}

	if description := req.Description; description.Valid {
		parts[keyDescription] = strings.NewReader(description.String)
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
