package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/organizations/team_members"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified team",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	organization api.Organization
	teamName string
	userName string
	requestBody team_members.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	organization, err := command.RequireOrganization(c)

	if err != nil {
		return nil, err
	}

	teamName, err := command.RequireTeamName(c)

	if err != nil {
		return nil, err
	}

	userName, err := command.RequireUserName(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		organization: *organization,
		teamName:teamName,
		userName:userName,
		requestBody: team_members.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveOrganizationMemberFromTeam(cmd.organization, cmd.teamName, cmd.userName, cmd.requestBody)
}
