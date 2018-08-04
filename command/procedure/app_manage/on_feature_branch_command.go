package app_manage

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	appsCommand "github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"os"
	"os/exec"
)

func OnFeatureBranchCommand() *cli.Command {
	return &cli.Command{
		Name:   "on-feature-branch",
		Usage:  "Upload applications on updating feature branches and create distributions ",
		Action: command.AuthorizedCommandAction(newOnFeatureBranchCommand),
		Flags:  onFeatureBranchFlags(),
	}
}

type onFeatureBranchCommand struct {
	uploadCommand command.Command
}

func newOnFeatureBranchCommand(c *cli.Context) (command.Command, error) {
	if c.IsSet(constant.DistributionName) {
		logrus.Debugf("Use the specified distribution name { %s }", c.String(constant.DistributionName))
	} else if name := os.Getenv("DPG_DISTRIBUTION_NAME"); name != "" {
		c.Set(constant.DistributionName, name)
	} else if branchRef, err := exec.Command("sh", "-c", `git rev-parse --abbrev-ref HEAD`).Output(); err == nil {
		c.Set(constant.DistributionName, string(branchRef))
	} else {
		return nil, errors.New("distribution name is not found")
	}

	uploadCommand, err := appsCommand.NewUploadCommand(c)

	if err != nil {
		return nil, err
	}

	cmd := onFeatureBranchCommand{
		uploadCommand,
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd onFeatureBranchCommand) VerifyInput() error {
	if err := cmd.uploadCommand.VerifyInput(); err != nil {
		return err
	}

	return nil
}

func (cmd onFeatureBranchCommand) Run(authorization *api.Authorization) (string, error) {
	if str, err := cmd.uploadCommand.Run(authorization); err != nil {
		return "", err
	} else {
		return str, nil
	}
}
