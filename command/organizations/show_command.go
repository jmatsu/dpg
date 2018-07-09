package organizations

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/show"
	"github.com/jmatsu/dpg/command"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func ShowCommand() cli.Command {
	return cli.Command{
		Name:   "show",
		Usage:  "Show the specified organization",
		Action: command.CommandAction(newShowCommand),
		Flags:  showFlags(),
	}
}

type showCommand struct {
	endpoint      *api.OrganizationsEndpoint
	authority     *api.Authority
	requestParams *show.Request
}

func newShowCommand(c *cli.Context) (command.Command, error) {
	cmd := showCommand{
		authority: &api.Authority{
			Token: command.GetApiToken(c),
		},
		endpoint: &api.OrganizationsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: GetOrganizationName(c),
		},
		requestParams: &show.Request{},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd showCommand) VerifyInput() error {
	if cmd.authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if cmd.endpoint.OrganizationName != "" {
		logrus.Fatalln("organization name must not be specified")
	}

	return nil
}

func (cmd showCommand) Run() (string, error) {
	if bytes, err := cmd.endpoint.GetSingleRequest(*cmd.authority, *cmd.requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
