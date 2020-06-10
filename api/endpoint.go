package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/apps"
	appsDistributions "github.com/jmatsu/dpg/request/apps/distributions"
	appsMembers "github.com/jmatsu/dpg/request/apps/members"
	appsSharedTeams "github.com/jmatsu/dpg/request/apps/shared_team_relations"
	appsTeams "github.com/jmatsu/dpg/request/apps/teams"
	"github.com/jmatsu/dpg/request/distributions"
	enterprisesMembers "github.com/jmatsu/dpg/request/enterprises/member_relations"
	enterprisesOrganizationsMembers "github.com/jmatsu/dpg/request/enterprises/organization_member_relations"
	enterprisesSharedTeams "github.com/jmatsu/dpg/request/enterprises/shared_teams"
	"github.com/jmatsu/dpg/request/organizations"
	organizationsMembers "github.com/jmatsu/dpg/request/organizations/members"
	"github.com/jmatsu/dpg/request/organizations/team_members"
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

func (e AppsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody apps.UploadRequest) ([]byte, error) {
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

func (e AppMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsMembers.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e AppMembersEndpoint) GetListRequest(authorization Authorization, requestParams appsMembers.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e AppMembersEndpoint) DeleteRequest(authorization Authorization, requestBody appsMembers.RemoveRequest) ([]byte, error) {
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

func (e AppDistributionsEndpoint) DeleteRequest(authorization Authorization, requestBody appsDistributions.DestroyRequest) ([]byte, error) {
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

func (e OrganizationAppTeamsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsTeams.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationAppTeamsEndpoint) GetListRequest(authorization Authorization, requestParams appsTeams.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationAppTeamsEndpoint) DeleteRequest(authorization Authorization, requestBody appsTeams.RemoveRequest) ([]byte, error) {
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

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody appsSharedTeams.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) GetListRequest(authorization Authorization, requestParams appsSharedTeams.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseOrganizationAppSharedTeamsEndpoint) DeleteRequest(authorization Authorization, requestBody appsSharedTeams.RemoveRequest) ([]byte, error) {
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

func (e DistributionsEndpoint) DeleteRequest(authorization Authorization, requestBody distributions.DestroyRequest) ([]byte, error) {
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

func (e OrganizationsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody organizations.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationsEndpoint) GetListRequest(authorization Authorization, requestParams organizations.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationsEndpoint) GetSingleRequest(authorization Authorization, requestParams organizations.ShowRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationsEndpoint) DeleteRequest(authorization Authorization, requestBody organizations.DestroyRequest) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}

func (e OrganizationsEndpoint) PatchRequest(authorization Authorization, requestBody organizations.UpdateRequest) ([]byte, error) {
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

func (e OrganizationMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody organizationsMembers.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationMembersEndpoint) GetListRequest(authorization Authorization, requestParams organizationsMembers.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationMembersEndpoint) DeleteRequest(authorization Authorization, requestBody organizationsMembers.RemoveRequest) ([]byte, error) {
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

func (e OrganizationTeamsMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody team_members.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e OrganizationTeamsMembersEndpoint) GetListRequest(authorization Authorization, requestParams team_members.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e OrganizationTeamsMembersEndpoint) DeleteRequest(authorization Authorization, requestBody team_members.RemoveRequest) ([]byte, error) {
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

func (e EnterpriseMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody enterprisesMembers.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseMembersEndpoint) GetListRequest(authorization Authorization, requestParams enterprisesMembers.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseMembersEndpoint) DeleteRequest(authorization Authorization, requestBody enterprisesMembers.RemoveRequest) ([]byte, error) {
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

func (e EnterpriseOrganizationsMembersEndpoint) MultiPartFormRequest(authorization Authorization, requestBody enterprisesOrganizationsMembers.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseOrganizationsMembersEndpoint) GetListRequest(authorization Authorization, requestParams enterprisesOrganizationsMembers.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseOrganizationsMembersEndpoint) DeleteRequest(authorization Authorization, requestBody enterprisesOrganizationsMembers.RemoveRequest) ([]byte, error) {
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

func (e EnterpriseSharedTeamsEndpoint) MultiPartFormRequest(authorization Authorization, requestBody enterprisesSharedTeams.CreateRequest) ([]byte, error) {
	return multiPartFormRequest(e, &authorization, requestBody)
}

func (e EnterpriseSharedTeamsEndpoint) GetListRequest(authorization Authorization, requestParams enterprisesSharedTeams.ListRequest) ([]byte, error) {
	return getRequest(e, &authorization, requestParams)
}

func (e EnterpriseSharedTeamsEndpoint) DeleteRequest(authorization Authorization, requestBody enterprisesSharedTeams.RemoveRequest) ([]byte, error) {
	return deleteRequest(e, &authorization, requestBody)
}
