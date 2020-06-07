package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/members"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified enterprise",
		Action: command.AuthorizedCommandAction(NewRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	enterprise  api.Enterprise
	userName    string
	requestBody members.RemoveRequest
}

func NewRemoveCommand(c *cli.Context) (command.Command, error) {
	enterprise, err := command.RequireEnterprise(c)

	if err != nil {
		return nil, err
	}

	userName, err := command.RequireUserName(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		enterprise:  *enterprise,
		userName:    userName,
		requestBody: members.RemoveRequest{},
	}

	return cmd, nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).RemoveEnterpriseMember(cmd.enterprise, cmd.userName, cmd.requestBody)
}
