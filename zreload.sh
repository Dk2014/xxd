#!/usr/bin/env bash


set -e

PJ=$(cd `dirname $0`; pwd)

SrvName=$1

if [ ! -f $PJ/system/reload ];then
	exit 0
fi

rm -rf $PJ/system/reload

cd $PJ

echo "Stop $SrvName..."
sh zserver.sh stop $SrvName

echo "building..."
sh zbuild.sh

cd $PJ/server/bin
nohup ./$SrvName -conf=gs.conf &
nohup ./$SrvName -conf=hd.conf &