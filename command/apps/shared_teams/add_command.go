package shared_teams

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/shared_teams/add"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
	"strings"
)

func AddCommand() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Add a shared team to the specified application",
		Action: command.AuthorizedCommandAction(newAddCommand),
		Flags:  addFlags(),
	}
}

type addCommand struct {
	endpoint    *api.OrganizationAppSharedTeamsEndpoint
	requestBody *add.Request
}

func newAddCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := addCommand{
		endpoint: &api.OrganizationAppSharedTeamsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: apps.GetAppOwnerName(c),
			AppId:            apps.GetAppId(c),
			AppPlatform:      platform,
		},
		requestBody: &add.Request{
			SharedTeamName: getSharedTeamName(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd addCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("an app owner name must be specified")
	}

	if cmd.endpoint.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(cmd.endpoint.AppPlatform, "android") && !strings.EqualFold(cmd.endpoint.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	if cmd.requestBody.SharedTeamName == "" {
		return errors.New("a shared team name must be specified")
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
