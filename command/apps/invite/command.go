package apps_invite

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/apps/invite"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "invite",
		Aliases: []string{"i"},
		Usage:   "Invite users to the specified application space",
		Action:  action,
		Flags:   allFlags(),
		Subcommands: []cli.Command{
			{
				Name:    "android",
				Aliases: []string{"a"},
				Usage:   "Invite users to the specified android application space",
				Action:  androidAppAction,
			},
			{
				Name:    "ios",
				Aliases: []string{"i"},
				Usage:   "Invite users to the specified iOS application space",
				Action:  iOSAppAction,
			},
		},
	}
}

func androidAppAction(c *cli.Context) error {
	if err := verifyAndroidApp(c); err != nil {
		return err
	}

	return action(c)
}

func iOSAppAction(c *cli.Context) error {
	if err := verifyIOSApp(c); err != nil {
		return err
	}

	return action(c)
}

func action(c *cli.Context) error {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.AppMemberEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: appOwnerName.Value(c).(string),
		AppId:        appId.Value(c).(string),
		AppPlatform:  appPlatform.Value(c).(string),
	}

	resp, err := inviteUsers(
		endpoint,
		authority,
		invite.Request{
			Invitees:      invitees.Value(c).([]string),
			DeveloperRole: role.Value(c).(null.Bool),
		},
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func inviteUsers(e api.AppMemberEndpoint, authority api.Authority, requestBody invite.Request, verbose bool) (response.AppInviteResponse, error) {
	var r response.AppInviteResponse

	if err := verifyInput(e, authority, requestBody); err != nil {
		return r, err
	}

	if bytes, err := e.MultiPartFormRequest(authority, requestBody, verbose); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}

func verifyAndroidApp(c *cli.Context) error {
	appPlatform := appPlatform.Value(c).(string)

	if !strings.EqualFold(appPlatform, "android") {
		return errors.New("A platform must be `android`")
	}

	return nil
}

func verifyIOSApp(c *cli.Context) error {
	appPlatform := appPlatform.Value(c).(string)

	if !strings.EqualFold(appPlatform, "ios") {
		return errors.New("A platform must be `ios`")
	}

	return nil
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

	if len(requestBody.Invitees) == 0 {
		return errors.New("the number of invitees must be greater than 0")
	}

	return nil
}
