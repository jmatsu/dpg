package command

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"gopkg.in/urfave/cli.v2"
)

type Command interface {
	VerifyInput() error
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
		authorization := &api.Authorization{
			Token: GetApiToken(c),
		}

		if authorization.Token == "" {
			return errors.New("api token must be specified")
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
