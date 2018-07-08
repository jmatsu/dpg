package apps_teams_remove

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/api/request/apps/teams/remove"
	"github.com/jmatsu/dpg/command/apps/teams"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Remove the team from the specified application",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = removeTeamFromApp(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationAppTeamsEndpoint, *api.Authority, *remove.Request, error) {
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
		TeamName:         teams.GetTeamName(c),
	}

	requestBody := remove.Request{}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.OrganizationAppTeamsEndpoint, authority api.Authority, _ remove.Request) error {
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

	if e.TeamName == "" {
		return errors.New("team name must be specified")
	}

	return nil
}

func removeTeamFromApp(e api.OrganizationAppTeamsEndpoint, authority api.Authority, requestBody remove.Request) (string, error) {
	if bytes, err := e.DeleteRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
