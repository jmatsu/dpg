package shared_teams

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/shared_teams/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/urfave/cli"
)

func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified shared team",
		Action: command.AuthorizedCommandAction(newAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.EnterpriseSharedTeamsEndpoint
	requestBody *add.Request
}

func newAddCommand(c *cli.Context) (command.Command, error) {
	cmd := addCommand{
		endpoint: &api.EnterpriseSharedTeamsEndpoint{
			BaseURL:        api.EndpointURL,
			EnterpriseName: enterprises.GetEnterpriseName(c),
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

func (cmd addCommand) VerifyInput() error {
	if cmd.endpoint.EnterpriseName == "" {
		return errors.New("an enterprise name must be specified")
	}

	if cmd.requestBody.SharedTeamName == "" {
		return errors.New("a team name must be specified")
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
