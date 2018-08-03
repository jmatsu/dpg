package distributions

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/distributions/destroy"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
)

func DestroyCommand() cli.Command {
	return cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified distribution",
		Action: command.AuthorizedCommandAction(newDestroyCommand),
		Flags:  removeFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.AppDistributionsEndpoint
	requestBody *destroy.Request
}

func newDestroyCommand(c *cli.Context) (command.Command, error) {
	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, err
	}

	cmd := destroyCommand{
		endpoint: &api.AppDistributionsEndpoint{
			BaseURL:      api.EndpointURL,
			AppOwnerName: apps.GetAppOwnerName(c),
			AppId:        apps.GetAppId(c),
			AppPlatform:  platform,
		},
		requestBody: &destroy.Request{
			DistributionName: GetDistributionName(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd destroyCommand) VerifyInput() error {
	if cmd.requestBody.DistributionName == "" {
		return errors.New(fmt.Sprintf("--%s must not be empty if specified", DistributionName.name()))
	}

	return nil
}

func (cmd destroyCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
