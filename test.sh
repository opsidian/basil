#!/usr/bin/env bash

set -e
echo -n "" > coverage.txt

export GO111MODULE=on

for dir in $(go list ./... | grep -v vendor); do
    go test -v -race -coverprofile=profile.out -covermode=atomic "$dir"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
