package organizations_destroy

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/api/request/organizations/destroy"
	"github.com/jmatsu/dpg/command/organizations"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "destroy",
		Usage:  "Remove the specified organization",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = destroyOrganization(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationsEndpoint, *api.Authority, *destroy.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.OrganizationsEndpoint{
		BaseURL:          api.EndpointURL,
		OrganizationName: organizations.GetOrganizationName(c),
	}

	requestBody := destroy.Request{}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.OrganizationsEndpoint, authority api.Authority, _ destroy.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName == "" {
		return errors.New("organization must be specified")
	}

	return nil
}

func destroyOrganization(e api.OrganizationsEndpoint, authority api.Authority, requestBody destroy.Request) (string, error) {
	if bytes, err := e.DeleteRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
