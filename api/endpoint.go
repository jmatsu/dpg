package api

import (
	"fmt"
	appsMembersAdd "github.com/jmatsu/dpg/api/request/apps/members/add"
	appsMembersList "github.com/jmatsu/dpg/api/request/apps/members/list"
	appsMembersRemove "github.com/jmatsu/dpg/api/request/apps/members/remove"
	appsSharedTeamsAdd "github.com/jmatsu/dpg/api/request/apps/shared_teams/add"
	appsSharedTeamsList "github.com/jmatsu/dpg/api/request/apps/shared_teams/list"
	appsSharedTeamsRemove "github.com/jmatsu/dpg/api/request/apps/shared_teams/remove"
	appsTeamsAdd "github.com/jmatsu/dpg/api/request/apps/teams/add"
	appsTeamsList "github.com/jmatsu/dpg/api/request/apps/teams/list"
	appsTeamsRemove "github.com/jmatsu/dpg/api/request/apps/teams/remove"
	appsUpload "github.com/jmatsu/dpg/api/request/apps/upload"
	distributionsRemove "github.com/jmatsu/dpg/api/request/distributions/destroy"
	organizationsCreate "github.com/jmatsu/dpg/api/request/organizations/create"
	organizationsRemove "github.com/jmatsu/dpg/api/request/organizations/destroy"
	organizationsList "github.com/jmatsu/dpg/api/request/organizations/list"
	organizationsShow "github.com/jmatsu/dpg/api/request/organizations/show"
	organizationsUpdate "github.com/jmatsu/dpg/api/request/organizations/update"
	"github.com/sirupsen/logrus"
	"os"
)

var EndpointURL string

func init() {
	if e := os.Getenv("DPG_ENDPOINT"); e != "" {
		EndpointURL = e
	} else {
		EndpointURL = "https://deploygate.com"
	}
}

type Endpoint interface {
	ToURL() string
}

// https://docs.deploygate.com/reference#upload

type AppUploadEndpoint struct {
	BaseURL      string
	AppOwnerName string
}

func (e AppUploadEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/users/%s/apps", e.BaseURL, e.AppOwnerName)

	logrus.Debugln(url)

	return url
}

func (e AppUploadEndpoint) MultiPartFormRequest(authority Authority, requestBody appsUpload.Request) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody)
}

// https://docs.deploygate.com/reference#invite
// https://docs.deploygate.com/reference#apps-members-index
// https://docs.deploygate.com/reference#apps-members-destroy

type AppMemberEndpoint struct {
	BaseURL      string
	AppOwnerName string
	AppPlatform  string
	AppId        string
}

func (e AppMemberEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/users/%s/platforms/%s/apps/%s/members", e.BaseURL, e.AppOwnerName, e.AppPlatform, e.AppId)

	logrus.Debugln(url)

	return url
}

func (e AppMemberEndpoint) MultiPartFormRequest(authority Authority, requestBody appsMembersAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody)
}

func (e AppMemberEndpoint) GetRequest(authority Authority, requestParams appsMembersList.Request) ([]byte, error) {
	return getRequest(e, authority, requestParams)
}

func (e AppMemberEndpoint) DeleteRequest(authority Authority, requestBody appsMembersRemove.Request) ([]byte, error) {
	return deleteRequest(e, authority, requestBody)
}

// https://docs.deploygate.com/reference#apps-teams-index
// https://docs.deploygate.com/reference#apps-teams-create
// https://docs.deploygate.com/reference#apps-teams-destroy

type OrganizationAppTeamsEndpoint struct {
	BaseURL          string
	OrganizationName string
	AppPlatform      string
	AppId            string
	TeamName         string
}

func (e OrganizationAppTeamsEndpoint) ToURL() string {
	var url string

	if url = fmt.Sprintf("%s/api/organizations/%s/platforms/%s/apps/%s/teams", e.BaseURL, e.OrganizationName, e.AppPlatform, e.AppId); e.TeamName != "" {
		url = fmt.Sprintf("%s/%s", url, e.TeamName)
	}

	logrus.Debugln(url)

	return url
}

func (e OrganizationAppTeamsEndpoint) MultiPartFormRequest(authority Authority, requestBody appsTeamsAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody)
}

func (e OrganizationAppTeamsEndpoint) GetRequest(authority Authority, requestParams appsTeamsList.Request) ([]byte, error) {
	return getRequest(e, authority, requestParams)
}

func (e OrganizationAppTeamsEndpoint) DeleteRequest(authority Authority, requestBody appsTeamsRemove.Request) ([]byte, error) {
	return deleteRequest(e, authority, requestBody)
}

// https://docs.deploygate.com/reference#apps-shared-teams-index
// https://docs.deploygate.com/reference#apps-shared-teams-create
// https://docs.deploygate.com/reference#apps-shared-teams-destroy

type OrganizationAppSharedTeamsEndpoint struct {
	BaseURL          string
	OrganizationName string
	AppPlatform      string
	AppId            string
	TeamName         string
}

func (e OrganizationAppSharedTeamsEndpoint) ToURL() string {
	var url string

	if url = fmt.Sprintf("%s/api/organizations/%s/platforms/%s/apps/%s/shared_teams", e.BaseURL, e.OrganizationName, e.AppPlatform, e.AppId); e.TeamName != "" {
		url = fmt.Sprintf("%s/%s", url, e.TeamName)
	}

	logrus.Debugln(url)

	return url
}

func (e OrganizationAppSharedTeamsEndpoint) MultiPartFormRequest(authority Authority, requestBody appsSharedTeamsAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody)
}

func (e OrganizationAppSharedTeamsEndpoint) GetRequest(authority Authority, requestParams appsSharedTeamsList.Request) ([]byte, error) {
	return getRequest(e, authority, requestParams)
}

func (e OrganizationAppSharedTeamsEndpoint) DeleteRequest(authority Authority, requestBody appsSharedTeamsRemove.Request) ([]byte, error) {
	return deleteRequest(e, authority, requestBody)
}

// https://docs.deploygate.com/reference#%E3%82%A2%E3%83%97%E3%83%AA%E3%81%AE%E9%85%8D%E5%B8%83%E3%83%9A%E3%83%BC%E3%82%B8%E3%82%92%E5%89%8A%E9%99%A4%E3%81%99%E3%82%8B

type DistributionsEndpoint struct {
	BaseURL         string
	DistributionKey string
}

func (e DistributionsEndpoint) ToURL() string {
	var url string

	if url = fmt.Sprintf("%s/api/distributions", e.BaseURL); e.DistributionKey != "" {
		url = fmt.Sprintf("%s/%s", url, e.DistributionKey)
	}

	logrus.Debugln(url)

	return url
}

func (e DistributionsEndpoint) DeleteRequest(authority Authority, requestBody distributionsRemove.Request) ([]byte, error) {
	return deleteRequest(e, authority, requestBody)
}

// https://docs.deploygate.com/reference#organizations-index
// https://docs.deploygate.com/reference#organizations-create
// https://docs.deploygate.com/reference#organizations-show
// https://docs.deploygate.com/reference#organizations-update
// https://docs.deploygate.com/reference#organizations-destroy

type OrganizationsEndpoint struct {
	BaseURL          string
	OrganizationName string
}

func (e OrganizationsEndpoint) ToURL() string {
	var url string

	if url = fmt.Sprintf("%s/api/organizations", e.BaseURL); e.OrganizationName != "" {
		url = fmt.Sprintf("%s/%s", url, e.OrganizationName)
	}

	logrus.Debugln(url)

	return url
}

func (e OrganizationsEndpoint) MultiPartFormRequest(authority Authority, requestBody organizationsCreate.Request) ([]byte, error) {
	return multiPartFormRequest(e, authority, requestBody)
}

func (e OrganizationsEndpoint) GetListRequest(authority Authority, requestParams organizationsList.Request) ([]byte, error) {
	return getRequest(e, authority, requestParams)
}

func (e OrganizationsEndpoint) GetSingleRequest(authority Authority, requestParams organizationsShow.Request) ([]byte, error) {
	return getRequest(e, authority, requestParams)
}

func (e OrganizationsEndpoint) DeleteRequest(authority Authority, requestBody organizationsRemove.Request) ([]byte, error) {
	return deleteRequest(e, authority, requestBody)
}

func (e OrganizationsEndpoint) PatchRequest(authority Authority, requestBody organizationsUpdate.Request) ([]byte, error) {
	return patchRequest(e, authority, requestBody)
}
