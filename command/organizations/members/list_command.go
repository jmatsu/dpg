package members

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/members/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/urfave/cli"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show users who have joined to the specified group",
		Action: command.AuthorizedCommandAction(newListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.OrganizationMembersEndpoint
	requestParams *list.Request
}

func newListCommand(c *cli.Context) (command.Command, error) {
	cmd := listCommand{
		endpoint: &api.OrganizationMembersEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: organizations.GetOrganizationName(c),
		},
		requestParams: &list.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd listCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("organization name must be specified")
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
