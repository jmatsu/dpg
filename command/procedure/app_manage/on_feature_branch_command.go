package app_manage

import (
	"errors"
	"flag"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	appsCommand "github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
	"os/exec"
	"strings"
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
	set := flag.NewFlagSet("hub", 0)
	set.Bool(constant.IsFeatureBranch, true, "_")
	set.String(constant.DistributionName, "", "_")
	set.String(constant.DistributionKey, "", "_")
	set.String(constant.ShortMessage, "", "_")
	set.String(constant.ReleaseNote, "", "_")

	nc := cli.NewContext(c.App, set, c)
	nc.Set(constant.IsFeatureBranch, "true")

	variableCatalog := newOnExposeCommandWithoutVerification(nc)

	if distributionName := variableCatalog.DistributionName; distributionName.Valid {
		nc.Set(constant.DistributionName, distributionName.String)
	} else if distributionKey := variableCatalog.DistributionKey; distributionKey.Valid {
		nc.Set(constant.DistributionKey, distributionKey.String)
	} else {
		return nil, errors.New("either distribution name or key is required")
	}

	if x := inferShortMessage(c); x.Valid {
		nc.Set(constant.ShortMessage, x.String)
	}

	if x := inferReleaseNote(c); x.Valid {
		nc.Set(constant.ReleaseNote, x.String)
	}

	uploadCommand, err := appsCommand.NewUploadCommand(nc)

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

func inferShortMessage(c *cli.Context) null.String {
	if c.IsSet(constant.ShortMessage) {
		return null.StringFrom(c.String(constant.ShortMessage))
	} else if x, err := exec.Command(`git`, `log`, `--format=%s`, `-1`).Output(); err == nil {
		return null.StringFrom(strings.TrimRight(string(x), "\n"))
	} else {
		return null.StringFromPtr(nil)
	}
}

func inferReleaseNote(c *cli.Context) null.String {
	if c.IsSet(constant.ReleaseNote) {
		return null.StringFrom(c.String(constant.ReleaseNote))
	} else if x, err := exec.Command(`git`, `log`, `--format=%b`, `-1`).Output(); err == nil {
		return null.StringFrom(strings.TrimRight(string(x), "\n"))
	} else {
		return null.StringFromPtr(nil)
	}
}
