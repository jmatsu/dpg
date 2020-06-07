package command

import (
	"github.com/jmatsu/dpg/api"
	"gopkg.in/urfave/cli.v2"
)

type Command interface {
	Run(authorization *api.Authorization) (string, error)
}

type generateCommandFunc func(c *cli.Context) (Command, error)

func CommandAction(fun generateCommandFunc) func(ctx *cli.Context) error {
	return func(c *cli.Context) error {
		if cmd, err := fun(c); err != nil {
			return err
		} else if _, err := cmd.Run(nil); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func AuthorizedCommandAction(fun generateCommandFunc) func(ctx *cli.Context) error {
	return func(c *cli.Context) error {
		apiToken, err := RequireApiToken(c)

		if err != nil {
			return err
		}

		authorization := &api.Authorization{
			Token: apiToken,
		}

		if cmd, err := fun(c); err != nil {
			return err
		} else if _, err := cmd.Run(authorization); err != nil {
			return err
		} else {
			return nil
		}
	}
}
