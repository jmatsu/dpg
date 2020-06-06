package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/apps/distributions/destroy"
)

func (c Client) DestroyDistributionByName(app App, request destroy.Request) (string, error) {
	if err := app.verify(); err != nil {
		return "", err
	}

	if request.DistributionName == "" {
		return "", fmt.Errorf("distribution name must be present")
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
