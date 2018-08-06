package app_manage

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
	"os/exec"
	"reflect"
	"strings"
)

func OnExposeCommand() *cli.Command {
	return &cli.Command{
		Name:   "expose",
		Usage:  "Expose some variables to be used app-manage procedure",
		Action: command.CommandAction(newOnExposeCommand),
		Flags:  exposeFlags(),
	}
}

type onExposeCommand struct {
	prefix             string
	ApiToken           null.String `expose:"DPG_API_TOKEN"`
	AppOwnerName       null.String `expose:"DPG_APP_OWNER_NAME"`
	DistributionName   null.String `expose:"DPG_DISTRIBUTION_NAME"`
	DistributionKey    null.String `expose:"DPG_DISTRIBUTION_KEY"`
	Platform           null.String `expose:"DPG_PLATFORM"`
	EnableNotification null.Bool   `expose:"DPG_ENABLE_NOTIFICATION"`
}

func newOnExposeCommand(c *cli.Context) (command.Command, error) {
	cmd := newOnExposeCommandWithoutVerification(c)

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newOnExposeCommandWithoutVerification(c *cli.Context) onExposeCommand {
	return onExposeCommand{
		prefix:             c.String(constant.Prefix),
		ApiToken:           inferApiToken(c),
		AppOwnerName:       inferAppOwnerName(c),
		DistributionName:   inferDistributionName(c),
		DistributionKey:    inferDistributionKey(c),
		Platform:           inferPlatform(c),
		EnableNotification: inferEnableNotification(c),
	}
}

func (cmd onExposeCommand) VerifyInput() error {
	val := reflect.ValueOf(&cmd).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanInterface() {
			if x, ok := field.Interface().(null.String); ok && x.Valid {
				if x.String == "" {
					return errors.New(fmt.Sprintf("%s must not be empty", val.Type().Field(i).Name))
				}
			}
		}
	}

	return nil
}

func (cmd onExposeCommand) Run(authorization *api.Authorization) (template string, _ error) {
	val := reflect.ValueOf(&cmd).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanInterface() {
			value := field.Interface()
			envName := val.Type().Field(i).Tag.Get("expose")

			if x, ok := value.(null.String); ok && x.Valid {
				template += fmt.Sprintf("%s%s='%s'\n", cmd.prefix, envName, x.String)
			} else if x, ok := value.(null.Bool); ok && x.Valid {
				template += fmt.Sprintf("%s%s='%t'\n", cmd.prefix, envName, x.Bool)
			}
		}
	}

	fmt.Println(template)

	return template, nil
}

func inferApiToken(c *cli.Context) null.String {
	if c.IsSet(constant.ApiToken) {
		return null.StringFrom(c.String(constant.ApiToken))
	} else {
		return null.StringFromPtr(nil)
	}
}

func inferAppOwnerName(c *cli.Context) null.String {
	if c.IsSet(constant.AppOwnerName) {
		return null.StringFrom(c.String(constant.AppOwnerName))
	} else {
		return null.StringFromPtr(nil)
	}
}

func inferDistributionName(c *cli.Context) null.String {
	var getCmd string

	if c.Bool(constant.IsFeatureBranch) {
		getCmd = `git rev-parse --abbrev-ref HEAD`
	} else {
		getCmd = `[ -n $(git show --merges HEAD -q) ] && git show HEAD -q --format=%s --merges | sed 's/^.* from [^\/]*\/\(.*\)$/\1/'`
	}

	if c.IsSet(constant.DistributionName) {
		return null.StringFrom(c.String(constant.DistributionName))
	} else if branchRef, err := exec.Command("sh", "-c", getCmd).Output(); err == nil {
		return null.StringFrom(strings.TrimRight(string(branchRef), "\n"))
	} else {
		return null.StringFromPtr(nil)
	}
}

func inferDistributionKey(c *cli.Context) null.String {
	if c.IsSet(constant.DistributionKey) {
		return null.StringFrom(c.String(constant.DistributionKey))
	} else {
		return null.StringFromPtr(nil)
	}
}

func inferPlatform(c *cli.Context) null.String {
	if c.IsSet(constant.Android) {
		return null.StringFrom(constant.Android)
	} else if c.IsSet(constant.IOS) {
		return null.StringFrom(constant.IOS)
	} else if _, err := exec.Command(`test`, `-n`, `$(find . -name "*.apk")`).Output(); err == nil {
		return null.StringFrom(constant.Android)
	} else if _, err := exec.Command(`test`, `-n`, `$(find . -name "*.ipa")`).Output(); err == nil {
		return null.StringFrom(constant.IOS)
	} else {
		return null.StringFromPtr(nil)
	}
}

func inferEnableNotification(c *cli.Context) null.Bool {
	if c.IsSet(constant.EnableNotification) {
		return null.BoolFrom(c.Bool(constant.EnableNotification))
	} else {
		return null.BoolFromPtr(nil)
	}
}
