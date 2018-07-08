package organizations_show

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/show"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "show",
		Usage:  "Show the specified organization",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestParams, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = showOrganization(
		*endpoint,
		*authority,
		*requestParams,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationsEndpoint, *api.Authority, *show.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.OrganizationsEndpoint{
		BaseURL:          api.EndpointURL,
		OrganizationName: organizations.GetOrganizationName(c),
	}

	requestParams := show.Request{}

	if err := verifyInput(endpoint, authority, requestParams); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestParams, nil
}

func verifyInput(e api.OrganizationsEndpoint, authority api.Authority, _ show.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName != "" {
		logrus.Fatalln("organization name must not be specified")
	}

	return nil
}

func showOrganization(e api.OrganizationsEndpoint, authority api.Authority, requestParams show.Request) (string, error) {
	if bytes, err := e.GetSingleRequest(authority, requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
