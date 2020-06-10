package organizations

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"strings"
)

type CreateRequest struct {
	OrganizationName string
	Description      null.String
}

func (req CreateRequest) IoReaderMap() (*map[string]io.Reader, error) {
	parts := map[string]io.Reader{
		"name": strings.NewReader(req.OrganizationName),
	}

	if description := req.Description; description.Valid {
		parts["description"] = strings.NewReader(description.String)
	}

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req CreateRequest) Verify() error {
	if req.OrganizationName == "" {
		return fmt.Errorf("organization name must be present")
	}

	return nil
}