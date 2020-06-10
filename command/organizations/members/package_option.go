package members

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

func addFlags() []cli.Flag {
	options := append(
		command.OrganizationOptions(),
		command.ApiToken,
		command.UserName,
		command.UserEmail,
	)

	return command.ToFlags(options)
}

func listFlags() []cli.Flag {
	options := append(
		command.OrganizationOptions(),
		command.ApiToken,
	)

	return command.ToFlags(options)
}

func removeFlags() []cli.Flag {
	options := append(
		command.OrganizationOptions(),
		command.ApiToken,
		command.UserName,
		command.UserEmail,
	)

	return command.ToFlags(options)
}

func assertUserNameOrUserEmail(name, email null.String) error {
	if name.String != "" && email.String != "" {
		return errors.New(fmt.Sprintf("only one of --%s or --%s is allowed", command.UserName.Flag().Names()[0], command.UserEmail.Flag().Names()[0]))
	} else if name.String == "" && email.String == "" {
		return errors.New(fmt.Sprintf("either of --%s or --%s must be specified", command.UserName.Flag().Names()[0], command.UserEmail.Flag().Names()[0]))
	}

	return nil
}
