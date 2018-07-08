package organizations_list

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/api/request/organizations/list"
	"github.com/jmatsu/dpg/command"
	"github.com/sirupsen/logrus"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show organizations which the user has",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestParams, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = listOrganizations(
		*endpoint,
		*authority,
		*requestParams,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationsEndpoint, *api.Authority, *list.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.OrganizationsEndpoint{
		BaseURL: api.EndpointURL,
	}

	requestParams := list.Request{}

	if err := verifyInput(endpoint, authority, requestParams); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestParams, nil
}

func verifyInput(e api.OrganizationsEndpoint, authority api.Authority, _ list.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName != "" {
		logrus.Fatalln("organization name must not be specified")
	}

	return nil
}

func listOrganizations(e api.OrganizationsEndpoint, authority api.Authority, requestParams list.Request) (string, error) {
	if bytes, err := e.GetListRequest(authority, requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
