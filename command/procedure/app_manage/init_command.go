package app_manage

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func InitCommand() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Action: command.CommandAction(newInitCommand),
	}
}

type initCommand struct {
	config config
}

func newInitCommand(_ *cli.Context) (command.Command, error) {
	cmd := initCommand{
		config: config{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd initCommand) VerifyInput() error {
	return nil
}

func (cmd initCommand) Run(_ *api.Authorization) (string, error) {
	//if str, err := cmd.destroyDistributionCommand.Run(authorization); err != nil {
	//	return "", err
	//} else {
	//	return str, nil
	//}

	// create a config?

	return "", nil
}
