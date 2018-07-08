package distributions_destroy

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/distributions/destroy"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/distributions"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Remove the specified distribution",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = destroyDistribution(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.DistributionsEndpoint, *api.Authority, *destroy.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.DistributionsEndpoint{
		BaseURL:         api.EndpointURL,
		DistributionKey: distributions.GetDistributionKey(c),
	}

	requestBody := destroy.Request{}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.DistributionsEndpoint, authority api.Authority, _ destroy.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.DistributionKey == "" {
		return errors.New("distribution key must be specified")
	}

	return nil
}

func destroyDistribution(e api.DistributionsEndpoint, authority api.Authority, requestBody destroy.Request) (string, error) {
	if bytes, err := e.DeleteRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
