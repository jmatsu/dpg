package distributions_remove

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/api/request/distributions/remove"
	"github.com/jmatsu/dpg/command/distributions"
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

	_, err = removeDistribution(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.DistributionsEndpoint, *api.Authority, *remove.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.DistributionsEndpoint{
		BaseURL:         api.EndpointURL,
		DistributionKey: distributions.GetDistributionKey(c),
	}

	requestBody := remove.Request{}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.DistributionsEndpoint, authority api.Authority, _ remove.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.DistributionKey == "" {
		return errors.New("distribution key must be specified")
	}

	return nil
}

func removeDistribution(e api.DistributionsEndpoint, authority api.Authority, requestBody remove.Request) (response.DistributionsRemoveResponse, error) {
	var r response.DistributionsRemoveResponse

	if err := verifyInput(e, authority, requestBody); err != nil {
		return r, err
	}

	if bytes, err := e.DeleteRequest(authority, requestBody); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}
