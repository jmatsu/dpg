package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/organizations/members"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show users who belong to the specified organization",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	organization api.Organization
	requestParams members.ListRequest
}

func NewListCommand(c *cli.Context) (command.Command, error) {
	organization, err := command.RequireOrganization(c)

	if err != nil {
		return nil, err
	}

	cmd := listCommand{
		organization:*organization,
		requestParams: members.ListRequest{},
	}

	return cmd, nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).ListOrganizationMembers(cmd.organization, cmd.requestParams)
}
