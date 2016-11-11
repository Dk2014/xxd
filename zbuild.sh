#!/usr/bin/env bash

set -e

PJ=$(cd `dirname $0`; pwd)

DB=$PJ/database
DOC=$PJ/protocol
SRV=$PJ/server/src/game_server
BIN=$PJ/server/bin
BM=$PJ/server/src/test/benchmark.go

OLD_GOPATH=$GOPATH
export GOPATH=$PJ/server

cd $PJ
trap "rm $PJ/zbuilding" EXIT
echo '' > ./zbuilding

echo "update protocol..."
cd $DOC
sh zcode

echo "update database..."
cd $DB
sh zupgrade

echo "generate code..."
sh zcode

echo "compiling server..."
cd $SRV
case $1 in
	static)
		# -a force rebuild ; -tags netgo pure go net package
		go install  -a -tags netgo  -ldflags="-w -extldflags '-static'"
		;;
	*)
		go install
	;;
esac

cp -fv $BIN/game_server $PJ/tools/api_viewer/
cp -fv $BIN/game_server $BIN/hd_server

echo "compiling benchmark tool..."
cd $PJ
go build $BM
mv benchmark $BIN

if [ ! -f $BIN/xxd.conf ];then
	cp -fv $PJ/xxd.conf $BIN
fi

export GOPATH=$OLD_GOPATH

echo "Successed in `date "+%Y-%m-%d %H:%M:%S"`"
