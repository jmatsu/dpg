package organizations

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/create"
	"github.com/jmatsu/dpg/command"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func CreateCommand() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "Create an organization",
		Action: command.CommandAction(newCreateCommand),
		Flags:  createFlags(),
	}
}

type createCommand struct {
	endpoint    *api.OrganizationsEndpoint
	authority   *api.Authority
	requestBody *create.Request
}

func newCreateCommand(c *cli.Context) (command.Command, error) {
	cmd := createCommand{
		authority: &api.Authority{
			Token: command.GetApiToken(c),
		},
		endpoint: &api.OrganizationsEndpoint{
			BaseURL: api.EndpointURL,
		},
		requestBody: &create.Request{
			OrganizationName: GetOrganizationName(c),
			Description:      getCreateDescription(c),
		},
	}

	if err := cmd.verifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd createCommand) verifyInput() error {
	if cmd.authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if cmd.endpoint.OrganizationName != "" {
		logrus.Fatalln("an organization name must not be specified")
	}

	if cmd.requestBody.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	return nil
}

func (cmd createCommand) run() (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*cmd.authority, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
