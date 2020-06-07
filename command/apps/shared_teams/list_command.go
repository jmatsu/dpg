package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/shared_teams"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show shared teams which belong to the specified application",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	app           api.OrganizationApp
	requestParams shared_teams.ListRequest
}

func NewListCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireOrganizationApp(c)

	if err != nil {
		return nil, err
	}

	cmd := listCommand{
		app:           *app,
		requestParams: shared_teams.ListRequest{},
	}

	return cmd, nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).ListSharedTeams(cmd.app, cmd.requestParams)
}
