package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/shared_teams"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Added a shared team to the specified enterprise",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	enterprise  api.Enterprise
	requestBody shared_teams.CreateRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	enterprise, err := command.RequireEnterprise(c)

	if err != nil {
		return nil, err
	}

	sharedTeamName, err := command.RequireSharedTeamName(c)

	if err != nil {
		return nil, err
	}

	description, err := command.GetSharedTeamDescription(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		enterprise: *enterprise,
		requestBody: shared_teams.CreateRequest{
			SharedTeamName: sharedTeamName,
			Description:    description.String,
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).CreateSharedTeam(cmd.enterprise, cmd.requestBody)
}
