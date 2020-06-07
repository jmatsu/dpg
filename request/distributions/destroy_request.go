package distributions

import (
	"github.com/jmatsu/dpg/util"
	"io"
)

type DestroyRequest struct{}

func (req DestroyRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
