package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/shared_teams/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Add a shared team to the specified application",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.EnterpriseOrganizationAppSharedTeamsEndpoint
	requestBody *add.Request
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		endpoint: &api.EnterpriseOrganizationAppSharedTeamsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: apps.GetAppOwnerName(c),
			AppId:            apps.GetAppId(c),
			AppPlatform:      platform,
		},
		requestBody: &add.Request{
			SharedTeamName: getSharedTeamName(c),
		},
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
Parameters:
	shared team name is required
*/
func (cmd addCommand) VerifyInput() error {
	if err := apps.RequireAppOwnerName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppId); err != nil {
		return err
	}

	if err := requireSharedTeamName(cmd.requestBody.SharedTeamName); err != nil {
		return err
	}

	return nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
