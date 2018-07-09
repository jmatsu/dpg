package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/organizations/members/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/jmatsu/dpg/command/enterprises/organizations"
	"github.com/urfave/cli"
)

func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified enterprise's organization",
		Action: command.AuthorizedCommandAction(newAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.EnterpriseOrganizationsMembersEndpoint
	requestBody *add.Request
}

func newAddCommand(c *cli.Context) (command.Command, error) {
	cmd := addCommand{
		endpoint: &api.EnterpriseOrganizationsMembersEndpoint{
			BaseURL:          api.EndpointURL,
			EnterpriseName:   enterprises.GetEnterpriseName(c),
			OrganizationName: organizations.GetOrganizationName(c),
		},
		requestBody: &add.Request{
			UserName: getUserName(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd addCommand) VerifyInput() error {
	if err := enterprises.RequireEnterpriseName(cmd.endpoint.EnterpriseName); err != nil {
		return err
	}

	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := requireUserName(cmd.requestBody.UserName); err != nil {
		return err
	}

	return nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
