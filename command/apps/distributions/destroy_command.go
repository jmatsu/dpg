package distributions

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/distributions/destroy"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified distribution",
		Action: command.AuthorizedCommandAction(NewDestroyCommand),
		Flags:  destroyFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.AppDistributionsEndpoint
	requestBody *destroy.Request
}

func NewDestroyCommand(c *cli.Context) (command.Command, error) {
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

/*
Endpoint:
	app owner's name is required
	app id is required
	app platform is required
Parameters:
	distribution name is required
*/
func (cmd destroyCommand) VerifyInput() error {
	if err := apps.RequireAppOwnerName(cmd.endpoint.AppOwnerName); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppId); err != nil {
		return err
	}

	if err := apps.RequireAppId(cmd.endpoint.AppPlatform); err != nil {
		return err
	}

	if cmd.requestBody.DistributionName == "" {
		return errors.New(fmt.Sprintf("--%s must not be empty if specified", constant.DistributionName))
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
