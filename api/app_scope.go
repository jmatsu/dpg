package api

import (
	"github.com/jmatsu/dpg/request/apps/distributions"
)

func (c Client) DestroyDistributionByName(app App, distributionName string) (string, error) {
	request := distributions.DestroyRequest {
		DistributionName:distributionName,
	}

	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := AppDistributionsEndpoint{
		BaseURL:      c.baseURL,
		AppOwnerName: app.OwnerName,
		AppId:        app.Id,
		AppPlatform:  app.Platform,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
