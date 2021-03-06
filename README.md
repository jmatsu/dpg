[![CircleCI](https://circleci.com/gh/jmatsu/dpg/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/dpg/tree/master)

# dpg

dpg - Golang implementation of Unofficial DeployGate API Client CLI

DeployGate API reference is https://docs.deploygate.com/reference#deploygate-api

## Usage

The basic syntax is:

    dpg command [command options] [arguments...]

Command list is [here](#COMMANDS)

`help, h` option is avaiable for all commands.
If you'd like to see the version, then run `dpg -v`.

### Bash/Zsh completion

```
// For Bash
eval $(dpg --init-completion bash)
dpg --init-completion bash >> ~/.bashrc


// For Zsh
eval $(dpg --init-completion zsh)
dpg --init-completion zsh >> ~/.zshrc
```

## Installation

```
go get github.com/jmatsu/dpg
```

Or download the binary

```
curl -sL "https://raw.githubusercontent.com/jmatsu/dpg/master/install.bash" | bash
curl -sL "https://raw.githubusercontent.com/jmatsu/dpg/master/install.bash" | VERSION=<0.2.1 or greater> bash
```

Or build on your local

```
go get -v -t -d ./...
go build
```

Docker containers are also available at https://hub.docker.com/r/jmatsu/dpg

## Advanced

`dpg` is providing some procedures to improve your deployment experience.  
They would be great help for you. You can see examples at [docs/procedure.md](docs/procedure.md).

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)

---

### COMMANDS

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
- `dpg procedure app-manage expose` [HELP](./docs/procedure.app-manage.expose.txt)
- `dpg procedure app-manage on-feature-branch` [HELP](./docs/procedure.app-manage.on-feature-branch.txt)
- `dpg procedure app-manage on-deploy-branch` [HELP](./docs/procedure.app-manage.on-deploy-branch.txt)
