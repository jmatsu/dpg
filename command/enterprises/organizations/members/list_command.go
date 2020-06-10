package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/organization_member_relations"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show users who have joined to the specified enterprise's organization",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	enterprise       api.Enterprise
	organizationName string
	requestParams    organization_member_relations.ListRequest
}

func NewListCommand(c *cli.Context) (command.Command, error) {
	enterprise, err := command.RequireEnterprise(c)

	if err != nil {
		return nil, err
	}

	organizationName, err := command.RequireOrganizationName(c)

	if err != nil {
		return nil, err
	}

	cmd := listCommand{
		enterprise:       *enterprise,
		organizationName: organizationName,
		requestParams:    organization_member_relations.ListRequest{},
	}

	return cmd, nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).ListEnterpriseMembersInOrganization(cmd.enterprise, cmd.organizationName, cmd.requestParams)
}
