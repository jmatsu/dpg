package apps

import (
	"github.com/urfave/cli"
	"errors"
	"github.com/sirupsen/logrus"
)

type option int

const (
	AppOwnerName option = iota
	AppId
	Android
	IOS
)

func (o option) Name() string {
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
			Name:  o.Name(),
			Usage: "[Required] An owner of application(s)",
		}
	case AppId:
		return cli.StringFlag{
			Name:  o.Name(),
			Usage: "[Required] An application id. e.g. com.deploygate",
		}
	case Android:
		return cli.BoolFlag{
			Name:  o.Name(),
			Usage: "[Required] Either of this or ios flag must be specified",
		}
	case IOS:
		return cli.BoolFlag{
			Name:  o.Name(),
			Usage: "[Required] Either of this or android flag must be specified",
		}
	}

	panic("Option name mapping is not found")
}

func GetAppOwnerName(c *cli.Context) string {
	return c.String(AppOwnerName.Name())
}

func GetAppId(c *cli.Context) string {
	return c.String(AppId.Name())
}

func GetAppPlatform(c *cli.Context) (string, error) {
	isAndroid := c.Bool(Android.Name())
	isIOS := c.Bool(IOS.Name())

	logrus.Debugf("android : %s, ios : %s\n", isAndroid, isIOS)

	if isAndroid && isIOS {
		return "", errors.New("only one option of android or ios is allowed")
	}

	if !isAndroid && !isIOS {
		return "", errors.New("either of android or ios must be specified")
	}

	if isAndroid {
		return "android", nil
	} else {
		return "ios", nil
	}
}
