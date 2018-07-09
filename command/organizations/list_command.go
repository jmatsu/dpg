package organizations

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/list"
	"github.com/jmatsu/dpg/command"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show organizations which the user has",
		Action: command.CommandAction(newListCommand),
		Flags:  listFlags(),
	}
}

type listCommand struct {
	endpoint      *api.OrganizationsEndpoint
	authority     *api.Authority
	requestParams *list.Request
}

func newListCommand(c *cli.Context) (command.Command, error) {
	cmd := listCommand{
		authority: &api.Authority{
			Token: command.GetApiToken(c),
		},
		endpoint: &api.OrganizationsEndpoint{
			BaseURL: api.EndpointURL,
		},
		requestParams: &list.Request{},
	}

	if err := cmd.verifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd listCommand) verifyInput() error {
	if cmd.authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if cmd.endpoint.OrganizationName != "" {
		logrus.Fatalln("organization name must not be specified")
	}

	return nil
}

func (cmd listCommand) run() (string, error) {
	if bytes, err := cmd.endpoint.GetListRequest(*cmd.authority, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
