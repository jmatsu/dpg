package command

import (
	"fmt"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
	"os"
)

type Option int

const (
	ApiToken Option = iota
	AppOwnerName
	AppId
	Android
	IOS
	AppFilePath
	Public
	EnableNotification
	ShortMessage
	DistributionKey
	DistributionName
	ReleaseNote
	Invitees
	DeveloperRole
	Removees
	SharedTeamName
	TeamName
	UserName
	EnterpriseName
)

func ToFlags(options []Option) []cli.Flag {
	flags := make([]cli.Flag, len(options))

	for i, o := range options {
		flags[i] = o.Flag()
	}

	return flags
}

func (o Option) Flag() cli.Flag {
	switch o {
	case ApiToken:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "[Required] API token",
			EnvVars: []string{constant.ApiTokenEnv, constant.DeployGateApiTokenEnv},
		}
	case AppOwnerName:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "[Required] The owner of the application",
			EnvVars: []string{constant.AppOwnerNameEnv, constant.DeployGateUserNameEnv},
		}
	case AppId:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "[Required] The application id. e.g. com.deploygate.sample",
			EnvVars: []string{constant.AppIdEnv},
		}
	case Android:
		return &cli.BoolFlag{
			Name:  o.name(),
			Usage: "[Required] Specify this if the application is an android application",
		}
	case IOS:
		return &cli.BoolFlag{
			Name:  o.name(),
			Usage: "[Required] Specify this if the application is an iOS application",
		}
	case AppFilePath:
		return &cli.PathFlag{
			Name:    o.name(),
			Usage:   "[Required] The file path of the application to be uploaded",
			EnvVars: []string{constant.AppFilePathEnv},
		}
	case Public:
		return &cli.BoolFlag{
			Name:  o.name(),
			Usage: "Specify true if an application to be uploaded should be public",
		}
	case EnableNotification:
		return &cli.BoolFlag{
			Name:    o.name(),
			EnvVars: []string{constant.EnableNotificationEnv},
			Usage:   "[iOS only] Specify true if iOS's notifications should be enabled",
		}
	case ShortMessage:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "A short message to explain this update",
			EnvVars: []string{constant.ShortMessageEnv, constant.DeployGateShortMessageEnv},
		}
	case DistributionKey:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "A key of a distribution to be updated",
			EnvVars: []string{constant.DistributionKeyEnv, constant.DeployGateDistributionKeyEnv},
		}
	case DistributionName:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "A name of a distribution to be updated",
			EnvVars: []string{constant.DistributionNameEnv},
		}
	case ReleaseNote:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "A release note for this revision",
			EnvVars: []string{constant.ReleaseNoteEnv, constant.DeployGateReleaseNoteEnv},
		}
	case Invitees:
		return &cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] Comma-separated names or e-mails of those whom you want to invite",
		}
	case DeveloperRole:
		return &cli.BoolFlag{
			Name:   o.name(),
			Usage:  "[Old Free/Lite/Pro/Biz plans only] Specify this if invitee(s) should be Developer Role, otherwise they would have Tester Role. Tester Role will be selected by default",
			Hidden: true,
		}
	case Removees:
		return &cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] Comma-separated names or e-mails of those who you want to remove",
		}
	case SharedTeamName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The name of the shared team",
		}
	case TeamName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the target team",
		}
	case UserName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The name of the user",
		}
	case EnterpriseName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the target enterprise.",
		}
	}

	panic("Option name mapping is not found")
}

func (o Option) name() string {
	switch o {
	case ApiToken:
		return constant.ApiToken
	case AppOwnerName:
		return constant.AppOwnerName
	case AppId:
		return constant.AppId
	case Android:
		return constant.Android
	case IOS:
		return constant.IOS
	case AppFilePath:
		return constant.AppFilePath
	case Public:
		return constant.IsPublic
	case EnableNotification:
		return constant.EnableNotification
	case ShortMessage:
		return constant.ShortMessage
	case DistributionKey:
		return constant.DistributionKey
	case DistributionName:
		return constant.DistributionName
	case ReleaseNote:
		return constant.ReleaseNote
	case Invitees:
		return constant.Invitees
	case DeveloperRole:
		return constant.DeveloperRole
	case Removees:
		return constant.Removees
	case SharedTeamName:
		return constant.SharedTeamName
	case TeamName:
		return constant.TeamName
	case UserName:
		return constant.UserName
	case EnterpriseName:
		return constant.EnterpriseName
	}

	panic("Option name mapping is not found")
}

func getString(c *cli.Context, option Option) (null.String, error) {
	if x := c.String(option.name()); x != "" {
		return null.StringFrom(x), nil
	} else if c.IsSet(option.name()) {
		return null.StringFromPtr(nil), fmt.Errorf("--%s must not be empty if specified", option.name())
	} else {
		return null.StringFromPtr(nil), nil
	}
}

func requireString(c *cli.Context, option Option) (string, error) {
	if value, err := getString(c, option); err != nil {
		return "", err
	} else if !value.Valid {
		return "", fmt.Errorf("--%s is required", option.name())
	} else {
		return value.String, nil
	}
}

func getStingSlice(c *cli.Context, option Option) ([]string, error) {
	return c.StringSlice(option.name()), nil
}

func requireStingSlice(c *cli.Context, option Option) ([]string, error) {
	if xs, err := getStingSlice(c, option); err != nil {
		return nil, err
	} else if len(xs) > 0 {
		return xs, nil
	} else {
		return nil, fmt.Errorf("the size of --%s must be at least 1", option.name())
	}
}

func GetApiToken(c *cli.Context) (null.String, error) {
	return getString(c, ApiToken)
}

func RequireApiToken(c *cli.Context) (string, error) {
	return requireString(c, ApiToken)
}

func GetAppOwnerName(c *cli.Context) (null.String, error) {
	return getString(c, AppOwnerName)
}

func RequireAppOwnerName(c *cli.Context) (string, error) {
	return requireString(c, AppOwnerName)
}

func GetAppId(c *cli.Context) (null.String, error) {
	return getString(c, AppId)
}

func RequireAppId(c *cli.Context) (string, error) {
	return requireString(c, AppId)
}

func GetAppPlatform(c *cli.Context) (null.String, error) {
	isAndroid := c.Bool(Android.name())
	isIOS := c.Bool(IOS.name())

	logrus.Debugf("android : %s, ios : %s\n", isAndroid, isIOS)

	if isAndroid && isIOS {
		return null.StringFromPtr(nil), fmt.Errorf("only one option of --%s or --%s is allowed", Android.name(), IOS.name())
	}

	if !isAndroid && !isIOS {
		platform := os.Getenv(constant.PlatformEnv)

		if platform == "" {
			return null.StringFromPtr(nil), nil
		} else if platform != constant.Android && platform != constant.IOS {
			return null.StringFromPtr(nil), fmt.Errorf("%s is not allowed. Only %s or %s are allowed", platform, constant.Android, constant.IOS)
		}
	}

	if isAndroid {
		return null.StringFrom(constant.Android), nil
	} else {
		return null.StringFrom(constant.IOS), nil
	}
}

func RequireAppPlatform(c *cli.Context) (string, error) {
	if platform, err := GetAppPlatform(c); err != nil {
		return "", err
	} else if !platform.Valid {
		return "", fmt.Errorf("either of --%s or --%s is required", Android.name(), IOS.name())
	} else {
		return platform.String, nil
	}
}

func GetAppFilePath(c *cli.Context) (null.String, error) {
	return getString(c, AppFilePath)
}

func RequireAppFilePath(c *cli.Context) (string, error) {
	return requireString(c, AppFilePath)
}

func IsPublic(c *cli.Context) bool {
	return c.Bool(Public.name())
}

func IsEnabledNotification(c *cli.Context) bool {
	return c.Bool(EnableNotification.name())
}

func GetShortMessage(c *cli.Context) (null.String, error) {
	return getString(c, ShortMessage)
}

func RequireShortMessage(c *cli.Context) (string, error) {
	return requireString(c, ShortMessage)
}

func GetDistributionKey(c *cli.Context) (null.String, error) {
	return getString(c, DistributionKey)
}

func RequireDistributionKey(c *cli.Context) (string, error) {
	return requireString(c, DistributionKey)
}

func GetDistributionName(c *cli.Context) (null.String, error) {
	return getString(c, DistributionName)
}

func RequireDistributionName(c *cli.Context) (string, error) {
	return requireString(c, DistributionName)
}

func GetReleaseNote(c *cli.Context) (null.String, error) {
	return getString(c, ReleaseNote)
}

func RequireReleaseNote(c *cli.Context) (string, error) {
	return requireString(c, ReleaseNote)
}

func GetInvitees(c *cli.Context) ([]string, error) {
	return getStingSlice(c, Invitees)
}

func RequireInvitees(c *cli.Context) ([]string, error) {
	return requireStingSlice(c, Invitees)
}

func IsDeveloperRole(c *cli.Context) bool {
	return c.Bool(DeveloperRole.name())
}

func GetRemovees(c *cli.Context) ([]string, error) {
	return getStingSlice(c, Removees)
}

func RequireRemovees(c *cli.Context) ([]string, error) {
	return requireStingSlice(c, Removees)
}

func GetSharedTeamName(c *cli.Context) (null.String, error) {
	return getString(c, SharedTeamName)
}

func RequireSharedTeamName(c *cli.Context) (string, error) {
	return requireString(c, SharedTeamName)
}

func GetTeamName(c *cli.Context) (null.String, error) {
	return getString(c, TeamName)
}

func RequireTeamName(c *cli.Context) (string, error) {
	return requireString(c, TeamName)
}

func GetUserName(c *cli.Context) (null.String, error) {
	return getString(c, UserName)
}

func RequireUserName(c *cli.Context) (string, error) {
	return requireString(c, UserName)
}

func GetEnterpriseName(c *cli.Context) (null.String, error) {
	return getString(c, EnterpriseName)
}

func RequireEnterpriseName(c *cli.Context) (string, error) {
	return requireString(c, EnterpriseName)
}
