package apps_teams_list

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/apps/teams"
	"github.com/urfave/cli"
)

//type option string

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		teams.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
	}
}

//func (o option) name() string {
//	switch name {
//	}
//
//	panic("Option name mapping is not found")
//}
//
//func (o option) flag() cli.flag {
//	switch name {
//	}
//
//	panic("Option name mapping is not found")
//}
