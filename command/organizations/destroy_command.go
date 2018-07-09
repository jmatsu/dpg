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
		Action: command.CommandAction(newDestroyCommand),
		Flags:  destroyFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.OrganizationsEndpoint
	authority   *api.Authority
	requestBody *destroy.Request
}

func newDestroyCommand(c *cli.Context) (command.Command, error) {
	cmd := destroyCommand{
		authority: &api.Authority{
			Token: command.GetApiToken(c),
		},
		endpoint: &api.OrganizationsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: GetOrganizationName(c),
		},
		requestBody: &destroy.Request{},
	}

	if err := cmd.verifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd destroyCommand) verifyInput() error {
	if cmd.authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if cmd.endpoint.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	return nil
}

func (cmd destroyCommand) run() (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*cmd.authority, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
