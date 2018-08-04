package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/enterprises/organizations/members/remove"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/jmatsu/dpg/command/enterprises/organizations"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified enterprise's organization",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	endpoint    *api.EnterpriseOrganizationsMembersEndpoint
	requestBody *remove.Request
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	cmd := removeCommand{
		endpoint: &api.EnterpriseOrganizationsMembersEndpoint{
			BaseURL:          api.EndpointURL,
			EnterpriseName:   enterprises.GetEnterpriseName(c),
			OrganizationName: organizations.GetOrganizationName(c),
			UserName:         getUserName(c),
		},
		requestBody: &remove.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	enterprise name is required
	organization name is required
	user name is required
Parameters:
	none
*/
func (cmd removeCommand) VerifyInput() error {
	if err := enterprises.RequireEnterpriseName(cmd.endpoint.EnterpriseName); err != nil {
		return err
	}

	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := requireUserName(cmd.endpoint.UserName); err != nil {
		return err
	}

	return nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
