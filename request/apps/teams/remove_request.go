package teams

import (
	"github.com/jmatsu/dpg/util"
	"io"
)

type RemoveRequest struct{}

func (req RemoveRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req RemoveRequest) Verify() error {
	return nil
}
