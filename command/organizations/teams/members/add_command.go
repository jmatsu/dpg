package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/jmatsu/dpg/command/organizations/teams"
	"github.com/jmatsu/dpg/request/organizations/teams/members/add"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified team",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.OrganizationTeamsMembersEndpoint
	requestBody *add.Request
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	cmd := addCommand{
		endpoint: &api.OrganizationTeamsMembersEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: organizations.GetOrganizationName(c),
			TeamName:         teams.GetTeamName(c),
		},
		requestBody: &add.Request{
			UserName: getUserName(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	organization name is required
	team name is required
Parameters:
	user name is required
*/
func (cmd addCommand) VerifyInput() error {
	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := teams.RequireTeamName(cmd.endpoint.TeamName); err != nil {
		return err
	}

	if err := requireUserName(cmd.requestBody.UserName); err != nil {
		return err
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
