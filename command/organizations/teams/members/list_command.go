package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/organizations/team_members"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show users who belong to the specified team",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	organization api.Organization
	teamName string
	requestParams team_members.ListRequest
}

func NewListCommand(c *cli.Context) (command.Command, error) {
	organization, err := command.RequireOrganization(c)

	if err != nil {
		return nil, err
	}

	teamName, err := command.RequireTeamName(c)

	if err != nil {
		return nil, err
	}


	cmd := listCommand{
		organization:*organization,
		teamName:teamName,
		requestParams: team_members.ListRequest{},
	}

	return cmd, nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).ListTeamMembers(cmd.organization, cmd.teamName, cmd.requestParams)
}
