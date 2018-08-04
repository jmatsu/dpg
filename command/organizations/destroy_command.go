package organizations

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/destroy"
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified organization",
		Action: command.AuthorizedCommandAction(NewDestroyCommand),
		Flags:  destroyFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.OrganizationsEndpoint
	requestBody *destroy.Request
}

func NewDestroyCommand(c *cli.Context) (command.Command, error) {
	cmd := destroyCommand{
		endpoint: &api.OrganizationsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: GetOrganizationName(c),
		},
		requestBody: &destroy.Request{},
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
func (cmd destroyCommand) VerifyInput() error {
	if err := RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (cmd destroyCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
