package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/enterprises/members"
	"gopkg.in/urfave/cli.v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified enterprise",
		Action: command.AuthorizedCommandAction(NewAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	enterprise  api.Enterprise
	requestBody members.AddRequest
}

func NewAddCommand(c *cli.Context) (command.Command, error) {
	enterprise, err := command.RequireEnterprise(c)

	if err != nil {
		return nil, err
	}

	userName, err := command.RequireUserName(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		enterprise: *enterprise,
		requestBody: members.AddRequest{
			UserName: userName,
		},
	}

	return cmd, nil
}

func (cmd addCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).AddEnterpriseMember(cmd.enterprise, cmd.requestBody)
}
