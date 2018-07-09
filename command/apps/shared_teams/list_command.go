package shared_teams

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/shared_teams/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
	"strings"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show shared teams which belong to the specified application",
		Action: command.AuthorizedCommandAction(newListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.OrganizationAppSharedTeamsEndpoint
	requestParams *list.Request
}

func newListCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := listCommand{
		endpoint: &api.OrganizationAppSharedTeamsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: apps.GetAppOwnerName(c),
			AppId:            apps.GetAppId(c),
			AppPlatform:      platform,
		},
		requestParams: &list.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd listCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("an app owner must be specified")
	}

	if cmd.endpoint.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(cmd.endpoint.AppPlatform, "android") && !strings.EqualFold(cmd.endpoint.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
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
