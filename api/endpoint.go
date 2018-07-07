package api

import (
	"fmt"
	requestAppUpload "github.com/jmatsu/dpg/api/request/app/upload"
	requestAppInvite "github.com/jmatsu/dpg/api/request/app/invite"
)

type Endpoint interface {
	ToURL() string
}

// https://docs.deploygate.com/reference#upload

type AppUploadEndpoint struct {
	BaseURL      string
	AppOwnerName string
}

func (e AppUploadEndpoint) ToURL() string {
	return fmt.Sprintf("%s/api/users/%s/apps", e.BaseURL, e.AppOwnerName)
}

func (e AppUploadEndpoint) MultiPartFormRequest(authority Authority, requestBody requestAppUpload.Request, verbose bool) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody, verbose)
}

// https://docs.deploygate.com/reference#invite

type AppInviteEndpoint struct {
	BaseURL      string
	AppOwnerName string
	AppPlatform  string
	AppId        string
}

func (e AppInviteEndpoint) ToURL() string {
	return fmt.Sprintf("%s/api/users/%s/platforms/%s/apps/%s/members", e.BaseURL, e.AppOwnerName, e.AppPlatform, e.AppId)
}

func (e AppInviteEndpoint) MultiPartFormRequest(authority Authority, requestBody requestAppInvite.Request, verbose bool) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody, verbose)
}
