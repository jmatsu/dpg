package invite

import (
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"strings"
)

type Request struct {
	UserNamesOrEmails []string
	DeveloperRole     null.Bool
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

	if devRole := req.DeveloperRole; devRole.Valid {
		if devRole.Bool {
			parts[keyDeveloperRole] = strings.NewReader("1")
		} else {
			parts[keyDeveloperRole] = strings.NewReader("2")
		}
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
