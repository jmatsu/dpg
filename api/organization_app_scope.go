package api

import (
	"fmt"
	"github.com/jmatsu/dpg/request/apps/shared_teams"
	"github.com/jmatsu/dpg/request/apps/teams"
)

func (c Client) AddTeam(app OrganizationApp, request teams.AddRequest) (string, error) {
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

func (c Client) ListTeams(app OrganizationApp, request teams.ListRequest) (string, error) {
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

func (c Client) RemoveTeam(app OrganizationApp, teamName string, request teams.RemoveRequest) (string, error) {
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

func (c Client) AddSharedTeam(app OrganizationApp, request shared_teams.AddRequest) (string, error) {
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

func (c Client) ListSharedTeams(app OrganizationApp, request shared_teams.ListRequest) (string, error) {
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

func (c Client) RemoveSharedTeam(app OrganizationApp, teamName string, request shared_teams.RemoveRequest) (string, error) {
	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	if teamName == "" {
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
