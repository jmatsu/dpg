package organizations_create

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/create"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "Create an organization",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = createOrganization(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationsEndpoint, *api.Authority, *create.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.OrganizationsEndpoint{
		BaseURL: api.EndpointURL,
	}

	requestBody := create.Request{
		OrganizationName: organizations.GetOrganizationName(c),
		Description:      getDescription(c),
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.OrganizationsEndpoint, authority api.Authority, request create.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName != "" {
		logrus.Fatalln("an organization name must not be specified")
	}

	if request.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	return nil
}

func createOrganization(e api.OrganizationsEndpoint, authority api.Authority, requestBody create.Request) (string, error) {
	if bytes, err := e.MultiPartFormRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
