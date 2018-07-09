package members

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/members/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/urfave/cli"
)

func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified group",
		Action: command.AuthorizedCommandAction(newAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.OrganizationMembersEndpoint
	requestBody *add.Request
}

func newAddCommand(c *cli.Context) (command.Command, error) {
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

func (cmd addCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	if !cmd.requestBody.UserName.Valid && !cmd.requestBody.UserEmail.Valid {
		return errors.New("either of user name or user email must be specified")
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
