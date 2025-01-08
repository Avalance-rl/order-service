#!/usr/bin/env bash

set -eu -o pipefail

if ! command -v golangci-lint &> /dev/null ; then
    echo "golangci-lint not installed or available in the PATH" >&2
    echo "you have to check https://github.com/golangci/golangci-lint" >&2
    exit 1
fi

exec golangci-lint run --config githooks/.golangci.yaml --new-from-rev HEAD "$@"