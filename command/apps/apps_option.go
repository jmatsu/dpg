package apps

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type option int

const (
	AppOwnerName option = iota
	AppId
	Android
	IOS
)

func (o option) name() string {
	switch o {
	case AppOwnerName:
		return "app-owner"
	case AppId:
		return "app-id"
	case Android:
		return "android"
	case IOS:
		return "ios"
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case AppOwnerName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The owner of the application",
		}
	case AppId:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The application id. e.g. com.deploygate.sample",
		}
	case Android:
		return cli.BoolFlag{
			Name:  o.name(),
			Usage: "[Required] Specify this if the application is an android application",
		}
	case IOS:
		return cli.BoolFlag{
			Name:  o.name(),
			Usage: "[Required] Specify this if the application is an iOS application",
		}
	}

	panic("Option name mapping is not found")
}

func GetAppOwnerName(c *cli.Context) string {
	return c.String(AppOwnerName.name())
}

func RequireAppOwnerName(appOwner string) error {
	if appOwner != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", AppOwnerName.name()))
}

func GetAppId(c *cli.Context) string {
	return c.String(AppId.name())
}

func RequireAppId(appId string) error {
	if appId != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", AppId.name()))
}

func GetAppPlatform(c *cli.Context) (string, error) {
	isAndroid := c.Bool(Android.name())
	isIOS := c.Bool(IOS.name())

	logrus.Debugf("android : %s, ios : %s\n", isAndroid, isIOS)

	if isAndroid && isIOS {
		return "", errors.New(fmt.Sprintf("only one option of --%s or --%s is allowed", Android.name(), IOS.name()))
	}

	if !isAndroid && !isIOS {
		return "", errors.New(fmt.Sprintf("either of --%s or --%s is needed", Android.name(), IOS.name()))
	}

	if isAndroid {
		return "android", nil
	} else {
		return "ios", nil
	}
}
