package api

import (
	"github.com/jmatsu/dpg/request/organizations/members"
	"github.com/jmatsu/dpg/request/organizations/team_members"
	"gopkg.in/guregu/null.v3"
)

func (c Client) AddOrganizationMember(organization Organization, request members.AddRequest) (string, error) {
	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationMembersEndpoint{
		BaseURL:      c.baseURL,
		OrganizationName:organization.Name,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListOrganizationMembers(organization Organization, request members.ListRequest) (string, error) {
	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationMembersEndpoint{
		BaseURL:      c.baseURL,
		OrganizationName:organization.Name,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveOrganizationMember(organization Organization, userName null.String, userEmail null.String, request members.RemoveRequest) (string, error) {
	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationMembersEndpoint{
		BaseURL:      c.baseURL,
		OrganizationName:organization.Name,
		UserName:userName,
		UserEmail:userEmail,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) AddTeamMember(organization Organization, teamName string, request team_members.AddRequest) (string, error) {
	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationTeamsMembersEndpoint{
		BaseURL:      c.baseURL,
		OrganizationName:organization.Name,
		TeamName:  teamName,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}


func (c Client) ListTeamMembers(organization Organization, teamName string, request team_members.ListRequest) (string, error) {
	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationTeamsMembersEndpoint{
		BaseURL:      c.baseURL,
		OrganizationName:organization.Name,
		TeamName:  teamName,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}


func (c Client) RemoveTeamMember(organization Organization, teamName string, userName string, request team_members.RemoveRequest) (string, error) {
	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationTeamsMembersEndpoint{
		BaseURL:      c.baseURL,
		OrganizationName:organization.Name,
		TeamName:  teamName,
		UserName:userName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
