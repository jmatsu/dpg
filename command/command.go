package command

import "github.com/urfave/cli"

type Command interface {
	VerifyInput() error
	Run() (string, error)
}

type generateCommandFunc func(c *cli.Context) (Command, error)

func CommandAction(fun generateCommandFunc) func(ctx *cli.Context) error {
	return func(c *cli.Context) error {
		if cmd, err := fun(c); err != nil {
			return err
		} else if _, err := cmd.Run(); err != nil {
			return err
		} else {
			return nil
		}
	}
}
