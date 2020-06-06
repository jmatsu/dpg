package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/jmatsu/dpg/request/organizations/members/add"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified organization",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.OrganizationMembersEndpoint
	requestBody *add.Request
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	cmd := addCommand{
		endpoint: &api.OrganizationMembersEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: organizations.GetOrganizationName(c),
		},
		requestBody: &add.Request{
			UserName:  getUserName(c),
			UserEmail: getUserEmail(c),
		},
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
	user name or user email is required
*/
func (cmd addCommand) VerifyInput() error {
	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := requireUserNameOrUserEmail(cmd.requestBody.UserName, cmd.requestBody.UserEmail); err != nil {
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
