package apps_invite

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/apps/invite"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "invite",
		Aliases: []string{"i"},
		Usage:   "Invite users to the specified application space",
		Action:  action,
		Flags:   flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	_, err = inviteUsers(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.AppMemberEndpoint, *api.Authority, *invite.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, nil, nil, err
	}

	endpoint := api.AppMemberEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: apps.GetAppOwnerName(c),
		AppId:        apps.GetAppId(c),
		AppPlatform:  platform,
	}

	requestBody := invite.Request{
		UserNamesOrEmails: getInvitees(c),
		DeveloperRole:     isDeveloperRole(c),
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.AppMemberEndpoint, authority api.Authority, requestBody invite.Request) error {
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

func inviteUsers(e api.AppMemberEndpoint, authority api.Authority, requestBody invite.Request) (response.AppInviteResponse, error) {
	var r response.AppInviteResponse

	if bytes, err := e.MultiPartFormRequest(authority, requestBody); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}
