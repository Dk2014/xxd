#!/usr/bin/env bash

set -e

PJ=$(cd `dirname $0`; pwd)

SRV=$PJ/server/src/tss_server

OLD_GOPATH=$GOPATH
export GOPATH=$PJ/server

cd $SRV
case $1 in
	static)
		go install -ldflags="-w -extldflags '-static'"
		;;
	*)
		go install
	;;
esac


export GOPATH=$OLD_GOPATH