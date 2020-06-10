package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/organizations/members"
	"github.com/jmatsu/dpg/request/organizations/team_members"
	"gopkg.in/guregu/null.v3"
)

func (c Client) AddOrganizationMember(organization Organization, userName null.String, userEmail null.String) (string, error) {
	request := members.CreateRequest{
		UserName:userName,
		UserEmail:userEmail,
	}

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

func (c Client) AddOrganizationMemberToTeam(organization Organization, teamName string, userName string) (string, error) {
	request := team_members.CreateRequest{
		UserName:userName,
	}

	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if teamName == "" {
		return "", fmt.Errorf("team name must be present")
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


func (c Client) ListOrganizationMembersInTeam(organization Organization, teamName string) (string, error) {
	request :=  team_members.ListRequest {

	}

	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if teamName == "" {
		return "", fmt.Errorf("team name must be present")
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


func (c Client) RemoveOrganizationMemberFromTeam(organization Organization, teamName string, userName string) (string, error) {
	request := team_members.RemoveRequest {
	}

	if err := organization.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if teamName == "" {
		return "", fmt.Errorf("team name must be present")
	}

	if userName == "" {
		return "", fmt.Errorf("user name must be present")
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
