#!/usr/bin/env bash


set -e

PJ=$(cd `dirname $0`; pwd)
PJ=$PJ/server

cd "${PJ}"

OLD_GOPATH=$GOPATH
export GOPATH=$PJ


cd src/payment/bin/payment


go install

cd "${PJ}"


export GOPATH=$OLD_GOPATH
