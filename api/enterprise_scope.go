package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/enterprises/member_relations"
	"github.com/jmatsu/dpg/request/enterprises/organization_member_relations"
	"github.com/jmatsu/dpg/request/enterprises/shared_teams"
	"gopkg.in/guregu/null.v3"
)

func (c Client) AddEnterpriseMember(enterprise Enterprise, userName string) (string, error) {
	request := member_relations.CreateRequest {
		UserName: userName,
	}

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

func (c Client) ListEnterpriseMembers(enterprise Enterprise) (string, error) {
	request := member_relations.ListRequest{}

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

func (c Client) RemoveEnterpriseMember(enterprise Enterprise, userName string) (string, error) {
	request := member_relations.RemoveRequest{}

	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if userName == "" {
		return "", fmt.Errorf("user name must be present")
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

func (c Client) AddEnterpriseMemberToOrganization(enterprise Enterprise, organization Organization, userName string) (string, error) {
	request := organization_member_relations.CreateRequest{
		UserName:userName,
	}

	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationsMembersEndpoint{
		BaseURL:          c.baseURL,
		EnterpriseName:   enterprise.Name,
		OrganizationName: organization.Name,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListEnterpriseMembersInOrganization(enterprise Enterprise, organization Organization) (string, error) {
	request := organization_member_relations.ListRequest{}

	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationsMembersEndpoint{
		BaseURL:          c.baseURL,
		EnterpriseName:   enterprise.Name,
		OrganizationName: organization.Name,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveEnterpriseMemberFromOrganization(enterprise Enterprise, organization Organization, userName string) (string, error) {
	request := organization_member_relations.RemoveRequest{}

	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err:=organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if userName == "" {
		return "", fmt.Errorf("user name must be present")
	}

	endpoint := EnterpriseOrganizationsMembersEndpoint{
		BaseURL:          c.baseURL,
		EnterpriseName:   enterprise.Name,
		OrganizationName: organization.Name,
		UserName:         userName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) CreateSharedTeam(enterprise Enterprise, sharedTeamName string, description null.String) (string, error) {
	request := shared_teams.CreateRequest {
		SharedTeamName:sharedTeamName,
		Description:description,
	}

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

func (c Client) ListSharedTeams(enterprise Enterprise) (string, error) {
	request := shared_teams.ListRequest {}

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

func (c Client) DestroySharedTeam(enterprise Enterprise, sharedTeamName string) (string, error) {
	request := shared_teams.DestroyRequest{}

	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if sharedTeamName == "" {
		return "", fmt.Errorf("shared team name must be present")
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
