package api

import (
	"fmt"
	"github.com/jmatsu/dpg/api/request"
)

type Endpoint interface {
	ToURL() string
}

type MultiFormRequestEndpoint interface {
	MultiPartFormRequest(authority Authority, requestBody request.Body, verbose bool) ([]byte, error)
}

// https://docs.deploygate.com/reference#upload

type UploadAppEndpoint struct {
	BaseURL      string
	AppOwnerName string
}

func (e UploadAppEndpoint) ToURL() string {
	return fmt.Sprintf("%s/api/users/%s/apps", e.BaseURL, e.AppOwnerName)
}

func (e UploadAppEndpoint) MultiPartFormRequest(authority Authority, requestBody request.Body, verbose bool) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody, verbose)
}

// https://docs.deploygate.com/reference#invite
