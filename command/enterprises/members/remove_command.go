package members

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/members/remove"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/urfave/cli"
)

func RemoveCommand() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified enterprise",
		Action: command.AuthorizedCommandAction(newRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	endpoint    *api.EnterpriseMembersEndpoint
	requestBody *remove.Request
}

func newRemoveCommand(c *cli.Context) (command.Command, error) {
	cmd := removeCommand{
		endpoint: &api.EnterpriseMembersEndpoint{
			BaseURL:        api.EndpointURL,
			EnterpriseName: enterprises.GetEnterpriseName(c),
			UserName:       getUserName(c),
		},
		requestBody: &remove.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd removeCommand) VerifyInput() error {
	if cmd.endpoint.EnterpriseName == "" {
		return errors.New("enterprise name must be specified")
	}

	if cmd.endpoint.UserName == "" {
		return errors.New("username must not be empty")
	}

	return nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
