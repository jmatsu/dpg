package organizations

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/destroy"
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
)

func DestroyCommand() cli.Command {
	return cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified organization",
		Action: command.AuthorizedCommandAction(newDestroyCommand),
		Flags:  destroyFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.OrganizationsEndpoint
	requestBody *destroy.Request
}

func newDestroyCommand(c *cli.Context) (command.Command, error) {
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

func (cmd destroyCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("organization name must be specified")
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
