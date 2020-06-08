package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/organization_members"
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
	enterprise       api.Enterprise
	organizationName string
	userName         string
	requestBody      organization_members.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	enterprise, err := command.RequireEnterprise(c)

	if err != nil {
		return nil, err
	}

	organizationName, err := command.RequireOrganizationName(c)

	if err != nil {
		return nil, err
	}

	userName, err := command.RequireUserName(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		enterprise:       *enterprise,
		organizationName: organizationName,
		userName:         userName,
		requestBody:      organization_members.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveEnterpriseOrganizationMember(cmd.enterprise, cmd.organizationName, cmd.userName, cmd.requestBody)
}
