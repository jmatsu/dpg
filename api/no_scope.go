package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/apps"
	"github.com/jmatsu/dpg/request/distributions/destroy"
)

func (c Client) UploadApp(appOwnerName string, request apps.UploadRequest) (string, error) {
	endpoint := AppsEndpoint{
		BaseURL:      c.baseURL,
		AppOwnerName: appOwnerName,
	}

	if appOwnerName == "" {
		return "", fmt.Errorf("app owner name must be present")
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) DestroyDistributionByKey(distributionKey string, request destroy.Request) (string, error) {
	if distributionKey == "" {
		return "", fmt.Errorf("distribution name must be present")
	}
	endpoint := DistributionsEndpoint{
		BaseURL:         c.baseURL,
		DistributionKey: distributionKey,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
