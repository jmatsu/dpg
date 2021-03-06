package organizations

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/list"
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show organizations which the user has",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.OrganizationsEndpoint
	requestParams *list.Request
}

func NewListCommand(_ *cli.Context) (command.Command, error) {
	cmd := listCommand{
		endpoint: &api.OrganizationsEndpoint{
			BaseURL: api.EndpointURL,
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
	none
Parameters:
	none
*/
func (cmd listCommand) VerifyInput() error {
	return nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.GetListRequest(*authorization, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
