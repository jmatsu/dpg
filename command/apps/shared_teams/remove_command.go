package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/shared_team_relations"
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
	app            api.OrganizationApp
	sharedTeamName string
	requestBody    shared_teams.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireOrganizationApp(c)

	if err != nil {
		return nil, err
	}

	sharedTeamName, err := command.RequireSharedTeamName(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		app:            *app,
		sharedTeamName: sharedTeamName,
		requestBody:    shared_teams.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveSharedTeamFromApp(cmd.app, cmd.sharedTeamName, cmd.requestBody)
}
