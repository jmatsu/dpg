package teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/organizations/teams"
	"github.com/jmatsu/dpg/request/apps/teams/add"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Add a team to the specified application",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.OrganizationAppTeamsEndpoint
	requestBody *add.Request
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		endpoint: &api.OrganizationAppTeamsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: apps.GetAppOwnerName(c),
			AppId:            apps.GetAppId(c),
			AppPlatform:      platform,
		},
		requestBody: &add.Request{
			TeamName: teams.GetTeamName(c),
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
	team name is required
*/
func (cmd addCommand) VerifyInput() error {
	if err := apps.RequireAppOwnerName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppId); err != nil {
		return err
	}

	if err := teams.RequireTeamName(cmd.requestBody.TeamName); err != nil {
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
