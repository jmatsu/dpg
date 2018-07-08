package apps_teams_add

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/request/apps/teams/add"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps/teams"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Add a team to the specified application",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = addTeamToApp(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationAppTeamsEndpoint, *api.Authority, *add.Request, error) {
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

	requestBody := add.Request{
		TeamName: teams.GetTeamName(c),
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.OrganizationAppTeamsEndpoint, authority api.Authority, request add.Request) error {
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

	if request.TeamName == "" {
		return errors.New("team name must be specified")
	}

	return nil
}

func addTeamToApp(e api.OrganizationAppTeamsEndpoint, authority api.Authority, requestBody add.Request) (string, error) {
	if bytes, err := e.MultiPartFormRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
