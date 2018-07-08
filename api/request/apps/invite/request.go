package invite

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type Request struct {
	UserNamesOrEmails []string
	DeveloperRole     bool
}

type Key string

const (
	keyInvitees      Key = "users"
	keyDeveloperRole Key = "role"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		keyInvitees: strings.NewReader(strings.Join(req.UserNamesOrEmails, ",")),
	}

	if req.DeveloperRole {
		parts[keyDeveloperRole] = strings.NewReader("1")
	} else {
		parts[keyDeveloperRole] = strings.NewReader("2")
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
