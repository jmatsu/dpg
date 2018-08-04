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

EOF

while read COMMAND; do
    echo  "- \`${COMMAND}\` [HELP](./"docs/$(split_by_space_and_get_tail ${COMMAND} | tr " " ".").md")"
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

create_pr() {
  local -r branch_name="$1"
  local -r api_url="https://api.github.com/repos/jmatsu/dpg/pulls"

  local body=("\"head\": \"$branch_name\"", "\"base\": \"master\"", "\"title\": \"Update Documents via CI\"")
  local json_body="{${body[*]}}"

  curl -s -H "Authorization: token ${GITHUB_ACCESS_TOKEN}" -H "Content-Type: application/json" -d "${json_body}" "${api_url}" || :
}

if [[ $(git log --merges --format='%s' -1|awk '$0=$NF') == "jmatsu/update_doc" ]]; then
    echo "Merged from doc update branch."
    return 0
fi

create_helps
create_readme > README.md

if [[ -z $(git diff) ]]; then
    return 0
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
