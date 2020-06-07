package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/shared_teams"
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
	app         api.OrganizationApp
	requestBody shared_teams.AddRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireOrganizationApp(c)

	if err != nil {
		return nil, err
	}

	sharedTeamName, err := command.RequireSharedTeamName(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		app: *app,
		requestBody: shared_teams.AddRequest{
			SharedTeamName: sharedTeamName,
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).AddSharedTeam(cmd.app, cmd.requestBody)
}
