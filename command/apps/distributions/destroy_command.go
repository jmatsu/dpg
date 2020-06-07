package distributions

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/apps/distributions"
	"gopkg.in/urfave/cli.v2"
)

func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified distribution by name",
		Action: command.AuthorizedCommandAction(NewDestroyCommand),
		Flags:  destroyFlags(),
	}
}

type destroyCommand struct {
	app         api.App
	requestBody distributions.DestroyRequest
}

func NewDestroyCommand(c *cli.Context) (command.Command, error) {
	appOwnerName, err := command.RequireAppOwnerName(c)

	if err != nil {
		return nil, err
	}

	appId, err := command.RequireAppId(c)

	if err != nil {
		return nil, err
	}

	platform, err := command.RequireAppPlatform(c)

	if err != nil {
		return nil, err
	}

	distributionName, err := command.RequireDistributionName(c)

	if err != nil {
		return nil, err
	}

	cmd := destroyCommand{
		app: api.App{
			OwnerName: appOwnerName,
			Id:        appId,
			Platform:  platform,
		},
		requestBody: distributions.DestroyRequest{
			DistributionName: distributionName,
		},
	}

	return cmd, nil
}

func (cmd destroyCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).DestroyDistributionByName(cmd.app, cmd.requestBody)
}
