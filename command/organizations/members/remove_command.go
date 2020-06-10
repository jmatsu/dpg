package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/organizations/members"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified organization",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	organization api.Organization
	userName null.String
	userEmail null.String
	requestBody members.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
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

	cmd := removeCommand{
		organization: *organization,
		userName:name,
		userEmail:email,
		requestBody: members.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveOrganizationMember(cmd.organization, cmd.userName, cmd.userEmail, cmd.requestBody)
}
