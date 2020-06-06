package organizations

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/organizations/create"
	"gopkg.in/urfave/cli.v2"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:   "create",
		Usage:  "Create an organization",
		Action: command.AuthorizedCommandAction(NewCreateCommand),
		Flags:  createFlags(),
	}
}

type createCommand struct {
	endpoint    *api.OrganizationsEndpoint
	requestBody *create.Request
}

func NewCreateCommand(c *cli.Context) (command.Command, error) {
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

/*
Endpoint:
	none
Parameters:
	organization name is required
*/
func (cmd createCommand) VerifyInput() error {
	if err := RequireOrganizationName(cmd.requestBody.OrganizationName); err != nil {
		return err
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
