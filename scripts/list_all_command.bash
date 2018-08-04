#!/usr/bin/env bash

set -eu

die() {
    echo $@ 1>&2
    exit 1
}

dpg_help() {
    dpg $@ -h
}

dfs() {
    _dfs ""
}

_dfs() {
    local command="$@"

    if [[ "$command" =~ .*help$ ]]; then
        return 0
    fi

    if dpg_help "${command}" | grep "COMMANDS:" >/dev/null 2>&1; then
        local -r head=$(($(dpg_help ${command} | grep -n "COMMANDS:" | awk -F: '$0=$1')))
        local -r last=$(($(dpg_help ${command} | grep -n "OPTIONS:" | awk -F: '$0=$1') - 1))

        local new_command=

        for new_command in $(dpg_help ${command} | awk "${head} < NR && NR < ${last} { print \$1 }" | xargs); do
            _dfs "${command} ${new_command%,}"
        done
    else
        echo dpg ${command}
    fi
}

dfs