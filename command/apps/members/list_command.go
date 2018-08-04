package members

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/members/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show users who belong to the specified application (expect the apps owner)",
		Action: command.AuthorizedCommandAction(newListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.AppMembersEndpoint
	requestParams *list.Request
}

func newListCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := listCommand{
		endpoint: &api.AppMembersEndpoint{
			BaseURL:      api.EndpointURL,
			AppOwnerName: apps.GetAppOwnerName(c),
			AppId:        apps.GetAppId(c),
			AppPlatform:  platform,
		},
		requestParams: &list.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	app owner's name is required
	app id is required
	app platform is required
Parameters:
	none
*/
func (cmd listCommand) VerifyInput() error {
	if err := apps.RequireAppOwnerName(cmd.endpoint.AppOwnerName); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppId); err != nil {
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
