package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/shared_teams"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove a shared team from the specified enterprise",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	enterprise     api.Enterprise
	sharedTeamName string
	requestBody    shared_teams.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	enterprise, err := command.RequireEnterprise(c)

	if err != nil {
		return nil, err
	}

	sharedTeamName, err := command.RequireSharedTeamName(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		enterprise:     *enterprise,
		sharedTeamName: sharedTeamName,
		requestBody:    shared_teams.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveSharedTeam2(cmd.enterprise, cmd.sharedTeamName, cmd.requestBody)
}
