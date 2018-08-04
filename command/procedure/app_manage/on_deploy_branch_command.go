package app_manage

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	distributionsCommand "github.com/jmatsu/dpg/command/apps/distributions"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"os"
	"os/exec"
)

func OnDeployBranchCommand() *cli.Command {
	return &cli.Command{
		Name:   "on-deploy",
		Usage:  "Delete associated distributions on deploy branch",
		Action: command.AuthorizedCommandAction(newOnDeployBranchCommand),
		Flags:  onDeployBranchFlags(),
	}
}

type onDeployBranchCommand struct {
	destroyDistributionCommand command.Command
}

func newOnDeployBranchCommand(c *cli.Context) (command.Command, error) {
	if c.IsSet(constant.DistributionName) {
		logrus.Debugf("Use the specified distribution name { %s }", c.String(constant.DistributionName))
	} else if name := os.Getenv("DPG_DISTRIBUTION_NAME"); name != "" {
		c.Set(constant.DistributionName, name)
	} else if branchRef, err := exec.Command("sh", "-c", `git log --format=%s --merges -1|sed 's/.*\/\([^/]*\)$/\1/'`).Output(); err == nil {
		c.Set(constant.DistributionName, string(branchRef))
	} else {
		return nil, err
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
