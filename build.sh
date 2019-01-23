#!/usr/bin/env bash
set -e

if [ "$(uname)" == "Darwin" ]; then
	export GOOS=linux
fi

PROJ="go_netlink_example"
ORG_PATH="github.com/example"
REPO_PATH="${ORG_PATH}/${PROJ}"

if [ ! -h .gopath/src/${REPO_PATH} ]; then
	mkdir -p .gopath/src/${ORG_PATH}
	ln -s ../../../.. .gopath/src/${REPO_PATH} || exit 255
fi

export GO15VENDOREXPERIMENT=1
export GOPATH=${PWD}/gopath
export GO="${GO:-go}"

mkdir -p "${PWD}/bin"
export GOBIN=${PWD}/bin


echo "Building go netlink examples"
$GO install "$@" ${REPO_PATH}/examples/rename

echo "Done!"
