[![CircleCI](https://circleci.com/gh/jmatsu/dpg/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/dpg/tree/master)

# dpg

dpg - Golang implementation of  DeployGate API Client CLI

DeployGate API reference is https://docs.deploygate.com/reference#deploygate-api

## Usage

The basic syntax is:

    dpg command [command options] [arguments...]

## Installation

```
go get github.com/jmatsu/dpg
```

### COMMANDS

`help, h` option is avaiable for all commands.
If you'd like to see the version, then run `dpg -v`.

- `dpg app upload` [HELP](./docs/app.upload.txt)
- `dpg app member add` [HELP](./docs/app.member.add.txt)
- `dpg app member list` [HELP](./docs/app.member.list.txt)
- `dpg app member remove` [HELP](./docs/app.member.remove.txt)
- `dpg app team add` [HELP](./docs/app.team.add.txt)
- `dpg app team remove` [HELP](./docs/app.team.remove.txt)
- `dpg app team list` [HELP](./docs/app.team.list.txt)
- `dpg app shared-team add` [HELP](./docs/app.shared-team.add.txt)
- `dpg app shared-team remove` [HELP](./docs/app.shared-team.remove.txt)
- `dpg app shared-team list` [HELP](./docs/app.shared-team.list.txt)
- `dpg app distributions destroy` [HELP](./docs/app.distributions.destroy.txt)
- `dpg distribution destroy` [HELP](./docs/distribution.destroy.txt)
- `dpg organization create` [HELP](./docs/organization.create.txt)
- `dpg organization destroy` [HELP](./docs/organization.destroy.txt)
- `dpg organization list` [HELP](./docs/organization.list.txt)
- `dpg organization show` [HELP](./docs/organization.show.txt)
- `dpg organization update` [HELP](./docs/organization.update.txt)
- `dpg organization member add` [HELP](./docs/organization.member.add.txt)
- `dpg organization member remove` [HELP](./docs/organization.member.remove.txt)
- `dpg organization member list` [HELP](./docs/organization.member.list.txt)
- `dpg organization team member add` [HELP](./docs/organization.team.member.add.txt)
- `dpg organization team member remove` [HELP](./docs/organization.team.member.remove.txt)
- `dpg organization team member list` [HELP](./docs/organization.team.member.list.txt)
- `dpg enterprise member add` [HELP](./docs/enterprise.member.add.txt)
- `dpg enterprise member remove` [HELP](./docs/enterprise.member.remove.txt)
- `dpg enterprise member list` [HELP](./docs/enterprise.member.list.txt)
- `dpg enterprise organization members add` [HELP](./docs/enterprise.organization.members.add.txt)
- `dpg enterprise organization members remove` [HELP](./docs/enterprise.organization.members.remove.txt)
- `dpg enterprise organization members list` [HELP](./docs/enterprise.organization.members.list.txt)
- `dpg enterprise shared-team add` [HELP](./docs/enterprise.shared-team.add.txt)
- `dpg enterprise shared-team remove` [HELP](./docs/enterprise.shared-team.remove.txt)
- `dpg enterprise shared-team list` [HELP](./docs/enterprise.shared-team.list.txt)

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)
