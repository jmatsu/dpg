package shared_teams

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/shared_teams/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/urfave/cli"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show users who have joined to the specified shared team",
		Action: command.AuthorizedCommandAction(newListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.EnterpriseSharedTeamsEndpoint
	requestParams *list.Request
}

func newListCommand(c *cli.Context) (command.Command, error) {
	cmd := listCommand{
		endpoint: &api.EnterpriseSharedTeamsEndpoint{
			BaseURL:        api.EndpointURL,
			EnterpriseName: enterprises.GetEnterpriseName(c),
		},
		requestParams: &list.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd listCommand) VerifyInput() error {
	if cmd.endpoint.EnterpriseName == "" {
		return errors.New("an enterprise name must be specified")
	}

	return nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.GetListRequest(*authorization, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
