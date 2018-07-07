package api

import (
	"fmt"
	"github.com/jmatsu/dpg/api/request/app/upload"
	"github.com/jmatsu/dpg/api/request/app/invite"
	"github.com/jmatsu/dpg/api/request/app/users"
	"github.com/jmatsu/dpg/api/request/organizations/teams/list"
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

func (e AppUploadEndpoint) MultiPartFormRequest(authority Authority, requestBody upload.Request, verbose bool) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody, verbose)
}

// https://docs.deploygate.com/reference#invite
// https://docs.deploygate.com/reference#apps-members-index

type AppMemberEndpoint struct {
	BaseURL      string
	AppOwnerName string
	AppPlatform  string
	AppId        string
}

func (e AppMemberEndpoint) ToURL() string {
	return fmt.Sprintf("%s/api/users/%s/platforms/%s/apps/%s/members", e.BaseURL, e.AppOwnerName, e.AppPlatform, e.AppId)
}

func (e AppMemberEndpoint) MultiPartFormRequest(authority Authority, requestBody invite.Request, verbose bool) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody, verbose)
}

func (e AppMemberEndpoint) GetQueryRequest(authority Authority, requestParams users.Request, verbose bool) ([]byte, error) {
	return getRequest(e, authority, requestParams, verbose)
}

// https://docs.deploygate.com/reference#apps-teams-index

type OrganizationTeamEndpoint struct {
	BaseURL          string
	OrganizationName string
	AppPlatform      string
	AppId            string
}

func (e OrganizationTeamEndpoint) ToURL() string {
	return fmt.Sprintf("%s/api/organizations/%s/platforms/%s/apps/%s/teams", e.BaseURL, e.OrganizationName, e.AppPlatform, e.AppId)
}

func (e OrganizationTeamEndpoint) GetQueryRequest(authority Authority, requestParams list.Request, verbose bool) ([]byte, error) {
	return getRequest(e, authority, requestParams, verbose)
}
