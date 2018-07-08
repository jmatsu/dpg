package apps_shared_teams_add

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/apps/shared_teams"
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
		shared_teams.TeamName.Flag(),
	}
}

//func (o option) name() string {
//	switch o {
//	}
//
//	panic("Option name mapping is not found")
//}

//func (o option) flag() cli.Flag {
//	switch o {
//	}
//
//	panic("Option name mapping is not found")
//}
