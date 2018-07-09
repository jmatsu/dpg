package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/teams/members/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/urfave/cli"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show users who belong to the specified team",
		Action: command.AuthorizedCommandAction(newListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.OrganizationTeamsMembersEndpoint
	requestParams *list.Request
}

func newListCommand(c *cli.Context) (command.Command, error) {
	cmd := listCommand{
		endpoint: &api.OrganizationTeamsMembersEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: organizations.GetOrganizationName(c),
			TeamName:         getTeamName(c),
		},
		requestParams: &list.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd listCommand) VerifyInput() error {
	if err := organizations.RequireOrganizationName(cmd.endpoint.OrganizationName); err != nil {
		return err
	}

	if err := requireTeamName(cmd.endpoint.TeamName); err != nil {
		return err
	}

	return nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.GetListRequest(*authorization, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
