package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/jmatsu/dpg/request/organizations/members"
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
	organization api.Organization
	requestBody members.CreateRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	organization, err := command.RequireOrganization(c)

	if err != nil {
		return nil, err
	}

	name, err := command.GetUserName(c)

	if err != nil {
		return nil, err
	}

	email, err := command.GetUserEmail(c)

	if err != nil {
		return nil, err
	}

	if err := assertUserNameOrUserEmail(name, email); err != nil {
		return nil, err
	}

	cmd := addCommand{
		organization: *organization,
		requestBody: members.CreateRequest {
			UserName:  name,
			UserEmail: email,
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
	return api.NewClient(*authorization).AddAppMember()
}
