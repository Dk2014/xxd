#!/usr/bin/env bash


set -e

PJ=$(cd `dirname $0`; pwd)
PJ=$PJ/server

cd "${PJ}"

OLD_GOPATH=$GOPATH
export GOPATH=$PJ


cd src/cron_server/


go install -ldflags="-w -extldflags '-static'"

cd "${PJ}"


export GOPATH=$OLD_GOPATH
