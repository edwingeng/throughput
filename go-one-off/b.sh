#!/usr/bin/env bash

[[ "$TRACE" ]] && set -x
pushd `dirname $0` > /dev/null
trap __EXIT EXIT

colorful=false
tput setaf 7 > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    colorful=true
fi

function __EXIT() {
    popd > /dev/null
}

function printError() {
    $colorful && tput setaf 1
    >&2  echo "Error: $@"
    $colorful && tput setaf 7
}

function printImportantMessage() {
    $colorful && tput setaf 3
    echo "$@"
    $colorful && tput setaf 7
}

function printUsage() {
    $colorful && tput setaf 3
    >&2  echo "$@"
    $colorful && tput setaf 7
}

go build -o server/server server/server.go
go build -o client/client client/client.go

server/server -n=1 &
sleep 1
client/client
sleep 1

echo
server/server -n=1 &
sleep 1
client/client -nodelay
sleep 1

wait > /dev/null 2>&1
