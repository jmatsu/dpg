package distributions

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/request/distributions"
	"gopkg.in/urfave/cli.v2"
)

func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified distribution",
		Action: command.AuthorizedCommandAction(NewDestroyCommand),
		Flags:  removeFlags(),
	}
}

type destroyCommand struct {
	distributionKey string
	requestBody     distributions.DestroyRequest
}

func NewDestroyCommand(c *cli.Context) (command.Command, error) {
	distributionKey, err := command.RequireDistributionKey(c)

	if err != nil {
		return nil, err
	}

	cmd := destroyCommand{
		distributionKey: distributionKey,
		requestBody:     distributions.DestroyRequest{},
	}

	return cmd, nil
}

func (cmd destroyCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).DestroyDistributionByKey(cmd.distributionKey, cmd.requestBody)
}
