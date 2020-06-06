package destroy

import (
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type Request struct {
	DistributionName string
}

type Key string

const (
	distributionName Key = "distribution_name"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[Key]io.Reader{
		distributionName: strings.NewReader(req.DistributionName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
