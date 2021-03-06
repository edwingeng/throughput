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

rm server client > /dev/null 2>&1

gcc -o2 -o server server.c
gcc -o2 -o client client.c

./server &
sleep 1
./client
sleep 1

echo
./server &
sleep 1
./client -nodelay
sleep 1

wait > /dev/null 2>&1
