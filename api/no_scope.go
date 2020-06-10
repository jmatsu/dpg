package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/apps"
	"github.com/jmatsu/dpg/request/distributions"
	"github.com/jmatsu/dpg/request/organizations"
	"gopkg.in/guregu/null.v3"
)

func (c Client) UploadApp(
	appOwnerName string,
	appFilePath string,
	appVisible bool,
	enableNotification bool,
	shortMessage null.String,
	distributionKey null.String,
	distributionName null.String,
	releaseNote null.String,
	) (string, error) {
		request := apps.UploadRequest{
			AppFilePath:        appFilePath,
			AppVisible:         appVisible,
			EnableNotification: enableNotification,
			ShortMessage:       shortMessage,
			DistributionKey:    distributionKey,
			DistributionName:   distributionName,
			ReleaseNote:        releaseNote,
		}

	if appOwnerName == "" {
		return "", fmt.Errorf("app owner name must be present")
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := AppsEndpoint{
		BaseURL:      c.baseURL,
		AppOwnerName: appOwnerName,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) DestroyDistributionByKey(distributionKey string) (string, error) {
	request :=  distributions.DestroyRequest{}

	if distributionKey == "" {
		return "", fmt.Errorf("distribution key must be present")
	}

	if err:= request.Verify(); err != nil {
		return "", err
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

func (c Client) CreateOrganization(organizationName string, description null.String) (string, error) {
	request := organizations.CreateRequest {
		OrganizationName:organizationName,
		Description:description,
	}

	if err:= request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationsEndpoint{
		BaseURL:c.baseURL,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}