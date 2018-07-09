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
		Action: command.AuthorizedCommandAction(newCreateCommand),
		Flags:  createFlags(),
	}
}

type createCommand struct {
	endpoint    *api.OrganizationsEndpoint
	requestBody *create.Request
}

func newCreateCommand(c *cli.Context) (command.Command, error) {
	cmd := createCommand{
		endpoint: &api.OrganizationsEndpoint{
			BaseURL: api.EndpointURL,
		},
		requestBody: &create.Request{
			OrganizationName: GetOrganizationName(c),
			Description:      getCreateDescription(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd createCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName != "" {
		logrus.Fatalln("an organization name must not be specified")
	}

	if cmd.requestBody.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	return nil
}

func (cmd createCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
