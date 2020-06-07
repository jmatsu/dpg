package organizations

import (
	"github.com/jmatsu/dpg/util"
)

type ShowRequest struct {
}

func (req ShowRequest) StringMap() (*map[string]string, error) {
	parts := map[string]string{}

	out, err := util.StringifyKeysAndValues(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
