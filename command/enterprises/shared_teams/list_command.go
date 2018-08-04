package shared_teams

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/shared_teams/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show shared teams which belong to the specified enterprise",
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

/*
Endpoint:
	enterprise name is required
Parameters:
	none
*/
func (cmd listCommand) VerifyInput() error {
	if err := enterprises.RequireEnterpriseName(cmd.endpoint.EnterpriseName); err != nil {
		return err
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
