package distributions_destroy

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/distributions"
	"github.com/urfave/cli"
)

//type option string

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		distributions.DistributionKey.Flag(),
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
