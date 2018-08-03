package api

import (
	"fmt"
	appsDistributionsDelete "github.com/jmatsu/dpg/api/request/apps/distributions/destroy"
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
	enterprisesMembersAdd "github.com/jmatsu/dpg/api/request/enterprises/members/add"
	enterprisesMembersList "github.com/jmatsu/dpg/api/request/enterprises/members/list"
	enterprisesMembersRemove "github.com/jmatsu/dpg/api/request/enterprises/members/remove"
	enterprisesOrganizationsMembersAdd "github.com/jmatsu/dpg/api/request/enterprises/organizations/members/add"
	enterprisesOrganizationsMembersList "github.com/jmatsu/dpg/api/request/enterprises/organizations/members/list"
	enterprisesOrganizationsMembersRemove "github.com/jmatsu/dpg/api/request/enterprises/organizations/members/remove"
	enterprisesSharedTeamsAdd "github.com/jmatsu/dpg/api/request/enterprises/shared_teams/add"
	enterprisesSharedTeamsList "github.com/jmatsu/dpg/api/request/enterprises/shared_teams/list"
	enterprisesSharedTeamsRemove "github.com/jmatsu/dpg/api/request/enterprises/shared_teams/remove"
	organizationsCreate "github.com/jmatsu/dpg/api/request/organizations/create"
	organizationsRemove "github.com/jmatsu/dpg/api/request/organizations/destroy"
	organizationsList "github.com/jmatsu/dpg/api/request/organizations/list"
	organizationsMembersAdd "github.com/jmatsu/dpg/api/request/organizations/members/add"
	organizationsMembersList "github.com/jmatsu/dpg/api/request/organizations/members/list"
	organizationsMembersRemove "github.com/jmatsu/dpg/api/request/organizations/members/remove"
	organizationsShow "github.com/jmatsu/dpg/api/request/organizations/show"
	organizationsTeamsMembersAdd "github.com/jmatsu/dpg/api/request/organizations/teams/members/add"
	organizationsTeamsMembersList "github.com/jmatsu/dpg/api/request/organizations/teams/members/list"
	organizationsTeamsMembersRemove "github.com/jmatsu/dpg/api/request/organizations/teams/members/remove"
	organizationsUpdate "github.com/jmatsu/dpg/api/request/organizations/update"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	netUrl "net/url"
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

type AppsEndpoint struct {
	BaseURL      string
	AppOwnerName string
}

func (e AppsEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/users/%s/apps", e.BaseURL, e.AppOwnerName)

	logrus.Debugln(url)

	return url
}

func (e AppsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsUpload.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#invite
// https://docs.deploygate.com/reference#apps-members-index
// https://docs.deploygate.com/reference#apps-members-destroy

type AppMembersEndpoint struct {
	BaseURL      string
	AppOwnerName string
	AppPlatform  string
	AppId        string
}

func (e AppMembersEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/users/%s/platforms/%s/apps/%s/members", e.BaseURL, e.AppOwnerName, e.AppPlatform, e.AppId)

	logrus.Debugln(url)

	return url
}

func (e AppMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsMembersAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e AppMembersEndpoint) GetListRequest(authorization Authorization, requestParams appsMembersList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e AppMembersEndpoint) DeleteRequest(authorization Authorization, requestBody appsMembersRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#invite

type AppDistributionsEndpoint struct {
	BaseURL      string
	AppOwnerName string
	AppPlatform  string
	AppId        string
}

func (e AppDistributionsEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/users/%s/platforms/%s/apps/%s/distributions", e.BaseURL, e.AppOwnerName, e.AppPlatform, e.AppId)

	logrus.Debugln(url)

	return url
}

func (e AppDistributionsEndpoint) DeleteRequest(authorization Authorization, requestBody appsDistributionsDelete.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
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

func (e OrganizationAppTeamsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsTeamsAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationAppTeamsEndpoint) GetListRequest(authorization Authorization, requestParams appsTeamsList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationAppTeamsEndpoint) DeleteRequest(authorization Authorization, requestBody appsTeamsRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#apps-shared-teams-index
// https://docs.deploygate.com/reference#apps-shared-teams-create
// https://docs.deploygate.com/reference#apps-shared-teams-destroy

type EnterpriseOrganizationAppSharedTeamsEndpoint struct {
	BaseURL          string
	OrganizationName string
	AppPlatform      string
	AppId            string
	SharedTeamName   string
}

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) ToURL() string {
	var url string

	if url = fmt.Sprintf("%s/api/organizations/%s/platforms/%s/apps/%s/shared_teams", e.BaseURL, e.OrganizationName, e.AppPlatform, e.AppId); e.SharedTeamName != "" {
		url = fmt.Sprintf("%s/%s", url, e.SharedTeamName)
	}

	logrus.Debugln(url)

	return url
}

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsSharedTeamsAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) GetListRequest(authorization Authorization, requestParams appsSharedTeamsList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) DeleteRequest(authorization Authorization, requestBody appsSharedTeamsRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
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

func (e DistributionsEndpoint) DeleteRequest(authorization Authorization, requestBody distributionsRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
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

func (e OrganizationsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody organizationsCreate.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationsEndpoint) GetListRequest(authorization Authorization, requestParams organizationsList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationsEndpoint) GetSingleRequest(authorization Authorization, requestParams organizationsShow.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationsEndpoint) DeleteRequest(authorization Authorization, requestBody organizationsRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

func (e OrganizationsEndpoint) PatchRequest(authorization Authorization, requestBody organizationsUpdate.Request) ([]byte, error) {
	return patchRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#organizations-members-index
// https://docs.deploygate.com/reference#organizations-members-create
// https://docs.deploygate.com/reference#organizations-members-destroy

type OrganizationMembersEndpoint struct {
	BaseURL          string
	OrganizationName string
	UserName         null.String
	UserEmail        null.String
}

func (e OrganizationMembersEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/organizations/%s/members", e.BaseURL, e.OrganizationName)

	if e.UserName.Valid {
		url = fmt.Sprintf("%s/%s", url, e.UserName.String)
	} else if e.UserEmail.Valid {
		url = fmt.Sprintf("%s/%s", url, netUrl.QueryEscape(e.UserEmail.String))
	}

	logrus.Debugln(url)

	return url
}

func (e OrganizationMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody organizationsMembersAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationMembersEndpoint) GetListRequest(authorization Authorization, requestParams organizationsMembersList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationMembersEndpoint) DeleteRequest(authorization Authorization, requestBody organizationsMembersRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#organizations-teams-users-index
// https://docs.deploygate.com/reference#organizations-teams-users-create
// https://docs.deploygate.com/reference#organizations-teams-users-destroy

type OrganizationTeamsMembersEndpoint struct {
	BaseURL          string
	OrganizationName string
	TeamName         string
	UserName         string
}

func (e OrganizationTeamsMembersEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/organizations/%s/teams/%s/users", e.BaseURL, e.OrganizationName, e.TeamName)

	if e.UserName != "" {
		url = fmt.Sprintf("%s/%s", url, e.UserName)
	}

	logrus.Debugln(url)

	return url
}

func (e OrganizationTeamsMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody organizationsTeamsMembersAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationTeamsMembersEndpoint) GetListRequest(authorization Authorization, requestParams organizationsTeamsMembersList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationTeamsMembersEndpoint) DeleteRequest(authorization Authorization, requestBody organizationsTeamsMembersRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#enterprises-users-index
// https://docs.deploygate.com/reference#enterprises-users-create
// https://docs.deploygate.com/reference#enterprises-users-destroy

type EnterpriseMembersEndpoint struct {
	BaseURL        string
	EnterpriseName string
	UserName       string
}

func (e EnterpriseMembersEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/enterprises/%s/users", e.BaseURL, e.EnterpriseName)

	if e.UserName != "" {
		url = fmt.Sprintf("%s/%s", url, e.UserName)
	}

	logrus.Debugln(url)

	return url
}

func (e EnterpriseMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody enterprisesMembersAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseMembersEndpoint) GetListRequest(authorization Authorization, requestParams enterprisesMembersList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseMembersEndpoint) DeleteRequest(authorization Authorization, requestBody enterprisesMembersRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#enterprises-organizations-users-index
// https://docs.deploygate.com/reference#enterprises-organizations-users-create
// https://docs.deploygate.com/reference#enterprises-organizations-users-destroy

type EnterpriseOrganizationsMembersEndpoint struct {
	BaseURL          string
	EnterpriseName   string
	OrganizationName string
	UserName         string
}

func (e EnterpriseOrganizationsMembersEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/enterprises/%s/organizations/%s/users", e.BaseURL, e.EnterpriseName, e.OrganizationName)

	if e.UserName != "" {
		url = fmt.Sprintf("%s/%s", url, e.UserName)
	}

	logrus.Debugln(url)

	return url
}

func (e EnterpriseOrganizationsMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody enterprisesOrganizationsMembersAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseOrganizationsMembersEndpoint) GetListRequest(authorization Authorization, requestParams enterprisesOrganizationsMembersList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseOrganizationsMembersEndpoint) DeleteRequest(authorization Authorization, requestBody enterprisesOrganizationsMembersRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

// https://docs.deploygate.com/reference#enterprises-shared-teams-index
// https://docs.deploygate.com/reference#enterprises-shared-teams-create
// https://docs.deploygate.com/reference#enterprises-shared-teams-destroy

type EnterpriseSharedTeamsEndpoint struct {
	BaseURL        string
	EnterpriseName string
	SharedTeamName string
}

func (e EnterpriseSharedTeamsEndpoint) ToURL() string {
	url := fmt.Sprintf("%s/api/enterprises/%s/shared_teams", e.BaseURL, e.EnterpriseName)

	if e.SharedTeamName != "" {
		url = fmt.Sprintf("%s/%s", url, e.SharedTeamName)
	}

	logrus.Debugln(url)

	return url
}

func (e EnterpriseSharedTeamsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody enterprisesSharedTeamsAdd.Request) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseSharedTeamsEndpoint) GetListRequest(authorization Authorization, requestParams enterprisesSharedTeamsList.Request) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseSharedTeamsEndpoint) DeleteRequest(authorization Authorization, requestBody enterprisesSharedTeamsRemove.Request) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}
