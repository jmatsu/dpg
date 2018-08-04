[![CircleCI](https://circleci.com/gh/jmatsu/dpg/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/dpg/tree/master)

# dpg

    dpg - Golang implementation of  DeployGate API Client CLI
    DeployGate API reference is https://docs.deploygate.com/reference#deploygate-api

## Usage

The basic syntax is:

   dpg command [command options] [arguments...]

If you'd like to see the version, then run `dpg version`.

### COMMANDS


`help, h` option is avaiable for all commands.

	dpg app upload [HELP](./docs/app.upload.md)
	dpg app member add [HELP](./docs/app.member.add.md)
	dpg app member list [HELP](./docs/app.member.list.md)
	dpg app member remove [HELP](./docs/app.member.remove.md)
	dpg app team add [HELP](./docs/app.team.add.md)
	dpg app team remove [HELP](./docs/app.team.remove.md)
	dpg app team list [HELP](./docs/app.team.list.md)
	dpg app shared-team add [HELP](./docs/app.shared-team.add.md)
	dpg app shared-team remove [HELP](./docs/app.shared-team.remove.md)
	dpg app shared-team list [HELP](./docs/app.shared-team.list.md)
	dpg app distributions destroy [HELP](./docs/app.distributions.destroy.md)
	dpg distribution destroy [HELP](./docs/distribution.destroy.md)
	dpg organization create [HELP](./docs/organization.create.md)
	dpg organization destroy [HELP](./docs/organization.destroy.md)
	dpg organization list [HELP](./docs/organization.list.md)
	dpg organization show [HELP](./docs/organization.show.md)
	dpg organization update [HELP](./docs/organization.update.md)
	dpg organization member add [HELP](./docs/organization.member.add.md)
	dpg organization member remove [HELP](./docs/organization.member.remove.md)
	dpg organization member list [HELP](./docs/organization.member.list.md)
	dpg organization team member add [HELP](./docs/organization.team.member.add.md)
	dpg organization team member remove [HELP](./docs/organization.team.member.remove.md)
	dpg organization team member list [HELP](./docs/organization.team.member.list.md)
	dpg enterprise member add [HELP](./docs/enterprise.member.add.md)
	dpg enterprise member remove [HELP](./docs/enterprise.member.remove.md)
	dpg enterprise member list [HELP](./docs/enterprise.member.list.md)
	dpg enterprise organization members add [HELP](./docs/enterprise.organization.members.add.md)
	dpg enterprise organization members remove [HELP](./docs/enterprise.organization.members.remove.md)
	dpg enterprise organization members list [HELP](./docs/enterprise.organization.members.list.md)
	dpg enterprise shared-team add [HELP](./docs/enterprise.shared-team.add.md)
	dpg enterprise shared-team remove [HELP](./docs/enterprise.shared-team.remove.md)
	dpg enterprise shared-team list [HELP](./docs/enterprise.shared-team.list.md)

## Installation

```
go get github.com/jmatsu/dpg
```

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)
