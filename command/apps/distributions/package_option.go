package distributions

import (
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
)

func destroyFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		DistributionName.Flag(),
	}
}
