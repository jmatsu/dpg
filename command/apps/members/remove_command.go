package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/members"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified application",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	app         api.UserApp
	requestBody members.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireUserApp(c)

	if err != nil {
		return nil, err
	}

	removees, err := command.RequireRemovees(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		app: *app,
		requestBody: members.RemoveRequest{
			UserNamesOrEmails: removees,
		},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveAppMember(cmd.app, cmd.requestBody)
}
