package api

import (
	"fmt"
	add2 "github.com/jmatsu/dpg/request/apps/shared_teams/add"
	list2 "github.com/jmatsu/dpg/request/apps/shared_teams/list"
	remove2 "github.com/jmatsu/dpg/request/apps/shared_teams/remove"
	"github.com/jmatsu/dpg/request/apps/teams/add"
	"github.com/jmatsu/dpg/request/apps/teams/list"
	"github.com/jmatsu/dpg/request/apps/teams/remove"
)

func (c Client) AddTeam(app OrganizationApp, request add.Request) (string, error) {
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

func (c Client) ListTeams(app UserApp, request list.Request) (string, error) {
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

func (c Client) RemoveTeam(app UserApp, teamName string, request remove.Request) (string, error) {
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

func (c Client) AddSharedTeam(app OrganizationApp, request add2.Request) (string, error) {
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

func (c Client) ListSharedTeams(app UserApp, request list2.Request) (string, error) {
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

func (c Client) RemoveSharedTeam(app UserApp, teamName string, request remove2.Request) (string, error) {
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
