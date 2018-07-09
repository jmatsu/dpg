package distributions

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/distributions/destroy"
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
)

func DestroyCommand() cli.Command {
	return cli.Command{
		Name:   "destroy",
		Usage:  "Destroy the specified distribution",
		Action: command.CommandAction(newDestroyCommand),
		Flags:  removeFlags(),
	}
}

type destroyCommand struct {
	endpoint    *api.DistributionsEndpoint
	authority   *api.Authority
	requestBody *destroy.Request
}

func newDestroyCommand(c *cli.Context) (command.Command, error) {
	cmd := destroyCommand{
		authority: &api.Authority{
			Token: command.GetApiToken(c),
		},
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
	if cmd.authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if cmd.endpoint.DistributionKey == "" {
		return errors.New("a distribution key must be specified")
	}

	return nil
}

func (cmd destroyCommand) Run() (string, error) {
	if bytes, err := cmd.endpoint.DeleteRequest(*cmd.authority, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
