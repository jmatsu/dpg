package shared_teams

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/shared_teams/remove"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
	"strings"
)

func RemoveCommand() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Removed a shared team from the specified application",
		Action: command.AuthorizedCommandAction(newRemoveCommand),
		Flags:  removeFlags(),
	}
}

type removeCommand struct {
	endpoint    *api.OrganizationAppSharedTeamsEndpoint
	requestBody *remove.Request
}

func newRemoveCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := removeCommand{
		endpoint: &api.OrganizationAppSharedTeamsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: apps.GetAppOwnerName(c),
			AppId:            apps.GetAppId(c),
			AppPlatform:      platform,
			SharedTeamName:   getSharedTeamName(c),
		},
		requestBody: &remove.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd removeCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("an app owner must be specified")
	}

	if cmd.endpoint.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(cmd.endpoint.AppPlatform, "android") && !strings.EqualFold(cmd.endpoint.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	if cmd.endpoint.SharedTeamName == "" {
		return errors.New("a shared team name must be specified")
	}

	return nil
}

func (cmd removeCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
