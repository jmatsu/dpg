#!/usr/bin/env bash

set -eu

split_by_space_and_get_tail() {
    (
        set -- $@
        shift 1
        echo $@
    )
}

create_helps() {
    rm -f docs/*.txt || :
    mkdir -p docs

    while read COMMAND; do
        ${COMMAND} -h > "docs/$(split_by_space_and_get_tail ${COMMAND} | tr " " ".").txt"
    done < <($(dirname "$0")/list_all_command.bash)
}

create_readme() {
cat<<'EOF'
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

EOF

while read COMMAND; do
    echo  "- \`${COMMAND}\` [HELP](./"docs/$(split_by_space_and_get_tail ${COMMAND} | tr " " ".").txt")"
done < <($(dirname "$0")/list_all_command.bash)
}

create_pr() {
  local -r branch_name="$1"
  local -r api_url="https://api.github.com/repos/jmatsu/dpg/pulls"

  local body=("\"head\": \"$branch_name\"", "\"base\": \"master\"", "\"title\": \"Update Documents via CI\"")
  local json_body="{${body[*]}}"

  curl -s -H "Authorization: token ${GITHUB_TOKEN}" -H "Content-Type: application/json" -d "${json_body}" "${api_url}" || :
}

if [[ $(git show --merges HEAD -q --format='%s' | awk '$0=$NF') == "jmatsu/update_doc" ]]; then
    echo "Merged from doc update branch."
    exit 0
fi

create_helps
create_readme > README.md

if [[ -z $(git diff) ]]; then
    exit 0
fi

if [[ "${CI:-false}" == "true" ]]; then
    git config user.email "jmatsu.drm+github@gmail.com"
    git config user.name "CircleCI job"
fi

branch_name="update_doc"

git checkout -b "$branch_name"
git add .
git commit -m "Updated docs ${CIRCLE_SHA1:-$(git rev-parse HEAD)}"
git push origin "$branch_name" -f

create_pr "${branch_name}"
