package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/organization_members"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified enterprise's organization",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	enterprise       api.Enterprise
	organizationName string
	requestBody      organization_members.AddRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
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

	cmd := addCommand{
		enterprise:       *enterprise,
		organizationName: organizationName,
		requestBody: organization_members.AddRequest{
			UserName: userName,
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).AddEnterpriseOrganizationMember(cmd.enterprise, cmd.organizationName, cmd.requestBody)
}
