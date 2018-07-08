package apps_teams_list

import (
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command/apps/teams"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
)

//type optionName string

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		teams.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
	}
}

//func (o optionName) name() string {
//	switch name {
//	}
//
//	panic("Option name mapping is not found")
//}
//
//func (o optionName) flag() cli.flag {
//	switch name {
//	}
//
//	panic("Option name mapping is not found")
//}
