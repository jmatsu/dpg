package distributions

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"io"
	"strings"
)

type DestroyRequest struct {
	DistributionName string
}

func (req DestroyRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"distribution_name": strings.NewReader(req.DistributionName),
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req DestroyRequest) Verify() error {
	if req.DistributionName == "" {
		return fmt.Errorf("distribution name must be present")
	}

	return nil
}