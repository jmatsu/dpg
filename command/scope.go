package command

import (
	"github.com/jmatsu/dpg/api"
	"gopkg.in/urfave/cli.v2"
)

func UserAppOptions() []Option {
	return []Option{
		AppOwnerName,
		AppId,
		Android,
		IOS,
	}
}

func OrganizationAppOptions() []Option {
	return []Option{
		AppOwnerName,
		AppId,
		Android,
		IOS,
	}
}

func EnterpriseOptions() []Option {
	return []Option{
		EnterpriseName,
	}
}

func RequireUserApp(c *cli.Context) (*api.UserApp, error) {
	appOwnerName, err := RequireAppOwnerName(c)

	if err != nil {
		return nil, err
	}

	appId, err := RequireAppId(c)

	if err != nil {
		return nil, err
	}

	platform, err := RequireAppPlatform(c)

	if err != nil {
		return nil, err
	}

	return &api.UserApp{
		OwnerName: appOwnerName,
		Id:        appId,
		Platform:  platform,
	}, err
}

func RequireOrganizationApp(c *cli.Context) (*api.OrganizationApp, error) {
	appOwnerName, err := RequireAppOwnerName(c)

	if err != nil {
		return nil, err
	}

	appId, err := RequireAppId(c)

	if err != nil {
		return nil, err
	}

	platform, err := RequireAppPlatform(c)

	if err != nil {
		return nil, err
	}

	return &api.OrganizationApp{
		OwnerName: appOwnerName,
		Id:        appId,
		Platform:  platform,
	}, err
}

func RequireEnterprise(c *cli.Context) (*api.Enterprise, error) {
	name, err := RequireEnterpriseName(c)

	if err != nil {
		return nil, err
	}

	return &api.Enterprise{
		Name: name,
	}, nil
}
