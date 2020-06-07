package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/members"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified application (available only for users' apps)",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	app         api.UserApp
	requestBody members.AddRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireUserApp(c)

	if err != nil {
		return nil, err
	}

	invitees, err := command.RequireInvitees(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		app: *app,
		requestBody: members.AddRequest{
			UserNamesOrEmails: invitees,
			DeveloperRole:     command.IsDeveloperRole(c),
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).AddMember(cmd.app, cmd.requestBody)
}
