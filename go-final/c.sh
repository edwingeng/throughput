#!/usr/bin/env bash

[[ "$TRACE" ]] && set -x
pushd `dirname $0` > /dev/null
trap __EXIT EXIT

function __EXIT() {
	popd > /dev/null
}

function printError() {
    tput setaf 1
    >&2  echo "Error: $@"
    tput setaf 7
}

function printImportantMessage() {
    tput setaf 3
    echo "$@"
    tput setaf 7
}

function printUsage() {
    tput setaf 3
    >&2  echo "$@"
    tput setaf 7
}

go build server.go
go build client.go

./server -n=3 &
sleep 1

./client &
./client &
./client &

wait > /dev/null 2>&1
sleep 1

echo
./server -n=3 &
sleep 1

./client -nodelay &
./client -nodelay &
./client -nodelay &

wait > /dev/null 2>&1
sleep 1
