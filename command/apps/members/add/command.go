package apps_members_add

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command"
	"github.com/sirupsen/logrus"
	"github.com/jmatsu/dpg/api/request/apps/members/add"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "add",
		Usage:  "Invite users to the specified application space",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	logrus.Debugln("apps/members/add")

	endpoint, authority, requestBody, err := buildResource(c)

	_, err = addUsersToApp(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.AppMemberEndpoint, *api.Authority, *add.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, nil, nil, err
	}

	endpoint := api.AppMemberEndpoint{
		BaseURL:      api.EndpointURL,
		AppOwnerName: apps.GetAppOwnerName(c),
		AppId:        apps.GetAppId(c),
		AppPlatform:  platform,
	}

	requestBody := add.Request{
		UserNamesOrEmails: getInvitees(c),
		DeveloperRole:     isDeveloperRole(c),
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.AppMemberEndpoint, authority api.Authority, requestBody add.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.AppOwnerName == "" {
		return errors.New("application owner must be specified")
	}

	if e.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(e.AppPlatform, "android") && !strings.EqualFold(e.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	if len(requestBody.UserNamesOrEmails) == 0 {
		return errors.New("the number of invitees must be greater than 0")
	}

	return nil
}

func addUsersToApp(e api.AppMemberEndpoint, authority api.Authority, requestBody add.Request) (string, error) {
	if bytes, err := e.MultiPartFormRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
