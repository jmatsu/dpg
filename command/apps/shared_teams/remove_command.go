package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/enterprises/shared_teams"
	"github.com/jmatsu/dpg/request/apps/shared_teams/remove"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Removed a shared team from the specified application",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	endpoint    *api.EnterpriseOrganizationAppSharedTeamsEndpoint
	requestBody *remove.Request
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		endpoint: &api.EnterpriseOrganizationAppSharedTeamsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: apps.GetAppOwnerName(c),
			AppId:            apps.GetAppId(c),
			AppPlatform:      platform,
			SharedTeamName:   shared_teams.GetSharedTeamName(c),
		},
		requestBody: &remove.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	organization name is required
	app id is required
	app platform is required
	shared team name is required
Parameters:
	none
*/
func (cmd removeCommand) VerifyInput() error {
	if err := apps.RequireAppOwnerName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppId); err != nil {
		return err
	}

	if err := shared_teams.RequireSharedTeamName(cmd.endpoint.SharedTeamName); err != nil {
		return err
	}

	return nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
