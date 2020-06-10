package request

import "github.com/jmatsu/dpg/util"

type Params interface {
	StringMap() (*map[string]string, error)
	Verify() error
}

func ToQuery(params Params) (string, error) {
	strMap, err := params.StringMap()

	if err != nil {
		return "", err
	}

	return util.ToQuery(*strMap)
}
