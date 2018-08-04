package distributions

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/distributions/destroy"
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified distribution",
		Action: command.AuthorizedCommandAction(newDestroyCommand),
		Flags:  removeFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.DistributionsEndpoint
	requestBody *destroy.Request
}

func newDestroyCommand(c *cli.Context) (command.Command, error) {
	cmd := destroyCommand{
		endpoint: &api.DistributionsEndpoint{
			BaseURL:         api.EndpointURL,
			DistributionKey: GetDistributionKey(c),
		},
		requestBody: &destroy.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd destroyCommand) VerifyInput() error {
	if cmd.endpoint.DistributionKey == "" {
		return errors.New(fmt.Sprintf("--%s must not be empty if specified", DistributionKey.name()))
	}

	return nil
}

func (cmd destroyCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
