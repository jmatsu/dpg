package organizations_update

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/update"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "update",
		Usage:  "Update an organization",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = updateOrganization(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationsEndpoint, *api.Authority, *update.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.OrganizationsEndpoint{
		BaseURL:          api.EndpointURL,
		OrganizationName: organizations.GetOrganizationName(c),
	}

	description := getDescription(c)

	if !description.Valid {
		return nil, nil, nil, errors.New("description must be specified")
	}

	requestBody := update.Request{
		Description: description.String,
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.OrganizationsEndpoint, authority api.Authority, _ update.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName == "" {
		logrus.Fatalln("an organization name must not be specified")
	}

	return nil
}

func updateOrganization(e api.OrganizationsEndpoint, authority api.Authority, requestBody update.Request) (string, error) {
	if bytes, err := e.PatchRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
