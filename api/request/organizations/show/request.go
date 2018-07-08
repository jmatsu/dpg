package show

import (
	"github.com/jmatsu/dpg/util"
)

type Request struct {
}

type Key string

func (req Request) StringMap() (*map[string]string, error) {
	parts := map[Key]string{}

	out, err := util.StringifyKeysAndValues(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
