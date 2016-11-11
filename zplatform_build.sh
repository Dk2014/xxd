#!/usr/bin/env bash


case $1 in
	run)
		killall platform_server
		;;
	*)
		
	;;
esac

set -e

PJ=$(cd `dirname $0`; pwd)
PJ=$PJ/server

cd "${PJ}"

# 引用的 package 都要放在src/core目录下而不是常规的引用
# 因此将GOPATH设置为当前目录

OLD_GOPATH=$GOPATH
export GOPATH=$PJ

cd src/platform_tool

case $1 in
	build)
		;;
	*)
		go test
	;;
esac

cd platform_tool

go install 

cd "${PJ}"

cd src/platform_server

case $1 in
	build)
		;;
	*)
		go test
	;;
esac


cd platform_server

go install

cd ../platform_admin

go install 


cd "${PJ}"

cd src/idip_server

case $1 in
	build)
		;;
	*)
		go test
	;;
esac

cd idip_server

#go install -ldflags="-w -extldflags '-static'"
go install 

cd ../parser

go install 

cd "${PJ}"

cd src/xdgm_server

case $1 in
	build)
		;;
	*)
		go test
	;;
esac

cd xdgm_server

go install


cd "${PJ}"

export GOPATH=$OLD_GOPATH

case $1 in
	run)
	sudo ./bin/platform_server -redis="114.112.58.162:56379"

	tail -f log/`date +%Y-%0m-%0d`.log
		;;
	*)
		
	;;
esac

