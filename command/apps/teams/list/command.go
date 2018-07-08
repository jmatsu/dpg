package apps_teams_list

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/teams/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
	"strings"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Show teams which belong to the specified application",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestParams, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = listTeams(
		*endpoint,
		*authority,
		*requestParams,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationAppTeamsEndpoint, *api.Authority, *list.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, nil, nil, err
	}

	endpoint := api.OrganizationAppTeamsEndpoint{
		BaseURL:          api.EndpointURL,
		OrganizationName: apps.GetAppOwnerName(c),
		AppId:            apps.GetAppId(c),
		AppPlatform:      platform,
	}

	requestParams := list.Request{}

	if err := verifyInput(endpoint, authority, requestParams); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestParams, nil
}

func verifyInput(e api.OrganizationAppTeamsEndpoint, authority api.Authority, _ list.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	if e.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(e.AppPlatform, "android") && !strings.EqualFold(e.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	return nil
}

func listTeams(e api.OrganizationAppTeamsEndpoint, authority api.Authority, requestParams list.Request) (string, error) {
	if bytes, err := e.GetRequest(authority, requestParams); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
