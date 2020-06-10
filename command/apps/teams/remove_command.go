package teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/teams"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Removed a team from the specified application",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	app         api.OrganizationApp
	teamName    string
	requestBody teams.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireOrganizationApp(c)

	if err != nil {
		return nil, err
	}

	teamName, err := command.RequireTeamName(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		app:         *app,
		teamName:    teamName,
		requestBody: teams.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveTeamFromApp(cmd.app, cmd.teamName, cmd.requestBody)
}
