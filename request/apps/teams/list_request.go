package teams

import (
	"github.com/jmatsu/dpg/util"
)

type ListRequest struct {
}

func (req ListRequest) StringMap() (*map[string]string, error) {
	parts := map[string]string{}

	out, err := util.StringifyKeysAndValues(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req ListRequest) Verify() error {
	return nil
}
