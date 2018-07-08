package organizations_show

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/urfave/cli"
)

//type option string

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
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
