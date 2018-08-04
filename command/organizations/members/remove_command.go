package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/members/remove"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified organization",
		Action: command.AuthorizedCommandAction(newRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	endpoint    *api.OrganizationMembersEndpoint
	requestBody *remove.Request
}

func newRemoveCommand(c *cli.Context) (command.Command, error) {
	cmd := removeCommand{
		endpoint: &api.OrganizationMembersEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: organizations.GetOrganizationName(c),
			UserEmail:        getUserEmail(c),
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
	organization name is required
	user name or user email is required
Parameters:
	none
*/
func (cmd removeCommand) VerifyInput() error {
	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := requireUserNameOrUserEmail(cmd.endpoint.UserName, cmd.endpoint.UserEmail); err != nil {
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
