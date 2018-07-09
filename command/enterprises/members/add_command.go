package members

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/members/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/urfave/cli"
)

func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified enterprise",
		Action: command.AuthorizedCommandAction(newAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.EnterpriseMembersEndpoint
	requestBody *add.Request
}

func newAddCommand(c *cli.Context) (command.Command, error) {
	cmd := addCommand{
		endpoint: &api.EnterpriseMembersEndpoint{
			BaseURL:        api.EndpointURL,
			EnterpriseName: enterprises.GetEnterpriseName(c),
		},
		requestBody: &add.Request{
			UserName: getUserName(c),
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

	if cmd.requestBody.UserName == "" {
		return errors.New("username must be specified")
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
