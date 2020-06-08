package api

import (
	"github.com/jmatsu/dpg/request/enterprises/members"
	"github.com/jmatsu/dpg/request/enterprises/organization_members"
	"github.com/jmatsu/dpg/request/enterprises/shared_teams"
)

func (c Client) AddEnterpriseMember(enterprise Enterprise, request members.AddRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseMembersEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListEnterpriseMembers(enterprise Enterprise, request members.ListRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseMembersEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveEnterpriseMember(enterprise Enterprise, userName string, request members.RemoveRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseMembersEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
		UserName:       userName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) AddEnterpriseOrganizationMember(enterprise Enterprise, organizationName string, request organization_members.AddRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationsMembersEndpoint{
		BaseURL:          c.baseURL,
		EnterpriseName:   enterprise.Name,
		OrganizationName: organizationName,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListEnterpriseOrganizationMembers(enterprise Enterprise, organizationName string, request organization_members.ListRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationsMembersEndpoint{
		BaseURL:          c.baseURL,
		EnterpriseName:   enterprise.Name,
		OrganizationName: organizationName,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveEnterpriseOrganizationMember(enterprise Enterprise, organizationName string, userName string, request organization_members.RemoveRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationsMembersEndpoint{
		BaseURL:          c.baseURL,
		EnterpriseName:   enterprise.Name,
		OrganizationName: organizationName,
		UserName:         userName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) AddSharedTeam2(enterprise Enterprise, request shared_teams.AddRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseSharedTeamsEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListSharedTeams2(enterprise Enterprise, request shared_teams.ListRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseSharedTeamsEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveSharedTeam2(enterprise Enterprise, sharedTeamName string, request shared_teams.RemoveRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseSharedTeamsEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
		SharedTeamName: sharedTeamName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
