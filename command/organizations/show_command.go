package organizations

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/show"
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
)

func ShowCommand() cli.Command {
	return cli.Command{
		Name:   "show",
		Usage:  "Show the specified organization",
		Action: command.AuthorizedCommandAction(newShowCommand),
		Flags:  showFlags(),
	}
}

type showCommand struct {
	endpoint      *api.OrganizationsEndpoint
	requestParams *show.Request
}

func newShowCommand(c *cli.Context) (command.Command, error) {
	cmd := showCommand{
		endpoint: &api.OrganizationsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: GetOrganizationName(c),
		},
		requestParams: &show.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	organization name is required
Parameters:
	none
*/
func (cmd showCommand) VerifyInput() error {
	if err := RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (cmd showCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.GetSingleRequest(*authorization, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
