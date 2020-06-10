package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/apps/shared_team_relations"
	"github.com/jmatsu/dpg/request/apps/teams"
)

func (c Client) AddTeamToApp(app OrganizationApp, teamName string) (string, error) {
	request := teams.CreateRequest{
		TeamName:teamName,
	}

	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationAppTeamsEndpoint{
		BaseURL:          c.baseURL,
		OrganizationName: app.OwnerName,
		AppId:            app.Id,
		AppPlatform:      app.Platform,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListTeamsInApp(app OrganizationApp) (string, error) {
	request := teams.ListRequest {}
	
	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := OrganizationAppTeamsEndpoint{
		BaseURL:          c.baseURL,
		OrganizationName: app.OwnerName,
		AppId:            app.Id,
		AppPlatform:      app.Platform,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveTeamFromApp(app OrganizationApp, teamName string) (string, error) {
	request := teams.RemoveRequest{}
	
	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if teamName == "" {
		return "", fmt.Errorf("team name must be present")
	}

	endpoint := OrganizationAppTeamsEndpoint{
		BaseURL:          c.baseURL,
		OrganizationName: app.OwnerName,
		AppId:            app.Id,
		AppPlatform:      app.Platform,
		TeamName:         teamName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) AddSharedTeamToApp(app OrganizationApp, sharedTeamName string) (string, error) {
	request := shared_team_relations.CreateRequest{
		SharedTeamName:sharedTeamName,
	}

	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationAppSharedTeamsEndpoint{
		BaseURL:          c.baseURL,
		OrganizationName: app.OwnerName,
		AppId:            app.Id,
		AppPlatform:      app.Platform,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListSharedTeamsInApp(app OrganizationApp) (string, error) {
	request := shared_team_relations.ListRequest{}

	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseOrganizationAppSharedTeamsEndpoint{
		BaseURL:          c.baseURL,
		OrganizationName: app.OwnerName,
		AppId:            app.Id,
		AppPlatform:      app.Platform,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveSharedTeamFromApp(app OrganizationApp, sharedTeamName string) (string, error) {
	request := shared_team_relations.DestroyRequest{}

	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if sharedTeamName == "" {
		return "", fmt.Errorf("team name must be present")
	}

	endpoint := EnterpriseOrganizationAppSharedTeamsEndpoint{
		BaseURL:          c.baseURL,
		OrganizationName: app.OwnerName,
		AppId:            app.Id,
		AppPlatform:      app.Platform,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
