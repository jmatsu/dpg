package app_manage

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	distributionsCommand "github.com/jmatsu/dpg/command/apps/distributions"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

func OnDeployBranchCommand() *cli.Command {
	return &cli.Command{
		Name:   "on-deploy-branch",
		Usage:  "Delete associated distributions on deploy branch",
		Action: command.AuthorizedCommandAction(newOnDeployBranchCommand),
		Flags:  onDeployBranchFlags(),
	}
}

type onDeployBranchCommand struct {
	destroyDistributionCommand command.Command
}

func newOnDeployBranchCommand(c *cli.Context) (command.Command, error) {
	// Don't need to control IsFeatureBranch option
	variableCatalog := newOnExposeCommandWithoutVerification(c)

	if distributionName := variableCatalog.DistributionName; distributionName.Valid {
		c.Set(constant.DistributionName, distributionName.String)
	} else if distributionKey := variableCatalog.DistributionKey; distributionKey.Valid {
		c.Set(constant.DistributionKey, distributionKey.String)
	} else {
		return nil, errors.New("either distribution name or key is required")
	}

	destroyDistributionCommand, err := distributionsCommand.NewDestroyCommand(c)

	if err != nil {
		return nil, err
	}

	cmd := onDeployBranchCommand{
		destroyDistributionCommand,
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd onDeployBranchCommand) VerifyInput() error {
	if err := cmd.destroyDistributionCommand.VerifyInput(); err != nil {
		return err
	}

	return nil
}

func (cmd onDeployBranchCommand) Run(authorization *api.Authorization) (string, error) {
	if str, err := cmd.destroyDistributionCommand.Run(authorization); err != nil {
		return "", err
	} else {
		return str, nil
	}
}
