package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/members"
	"gopkg.in/urfave/cli.v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "Show users who belong to the specified application (expect the apps owner)",
		Action: command.AuthorizedCommandAction(NewListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	app           api.UserApp
	requestParams members.ListRequest
}

func NewListCommand(c *cli.Context) (command.Command, error) {
	app, err := command.RequireUserApp(c)

	if err != nil {
		return nil, err
	}

	cmd := listCommand{
		app:           *app,
		requestParams: members.ListRequest{},
	}

	return cmd, nil
}

func (cmd listCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).ListMembers(cmd.app, cmd.requestParams)
}
