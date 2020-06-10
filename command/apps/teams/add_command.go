package teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/teams"
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
	app         api.OrganizationApp
	requestBody teams.CreateRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireOrganizationApp(c)

	if err != nil {
		return nil, err
	}

	teamName, err := command.RequireTeamName(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		app: *app,
		requestBody: teams.CreateRequest{
			TeamName: teamName,
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).AddTeamToApp(cmd.app, cmd.requestBody)
}
