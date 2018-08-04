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
    mkdir -p docs

    while read COMMAND; do
        ${COMMAND} -h > "docs/$(split_by_space_and_get_tail ${COMMAND} | tr " " ".").md"
    done < <($(dirname "$0")/list_all_command.bash)
}

create_readme() {
cat<<'EOF'
[![CircleCI](https://circleci.com/gh/jmatsu/deploygate-cli-go/tree/master.svg?style=svg)](https://circleci.com/gh/jmatsu/deploygate-cli-go/tree/master)

# dpg

    dpg - Golang implementation of  DeployGate API Client CLI
    DeployGate API reference is https://docs.deploygate.com/reference#deploygate-api

## Usage

The basic syntax is:

   dpg command [command options] [arguments...]

If you'd like to see the version, then run `dpg version`.

### COMMANDS


`help, h` option is avaiable for all commands.

EOF

while read COMMAND; do
    echo "\t${COMMAND} [HELP](./"docs/$(split_by_space_and_get_tail ${COMMAND} | tr " " ".").md")"
done < <($(dirname "$0")/list_all_command.bash)

cat<<'EOF'

## Installation

```
go get github.com/jmatsu/dpg
```

## LICENSE

Under MIT License. See [LICENSE](./LICENSE)
EOF
}

if [[ $(git log --merges --format='%s' -1|awk '$0=$NF') =~ ^update_doc_on_.* ]]; then
    echo "Merged from doc update branch."
    return 0
fi

create_helps
create_readme > README.md

if [[ -z $(git diff) ]]; then
    return 0
fi

branch_name="update_doc_on_$(git rev-parse --short HEAD)"

git config user.email "jmatsu.drm+github@gmail.com"
git config user.name "CircleCI job"
git checkout -b "$branch_name"
git add docs
git commit -m "Updated docs based on the script"
git push origin "$branch_name"
