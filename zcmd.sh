#!/usr/bin/env bash


set -x

PJ=$(cd `dirname $0`; pwd)

SrvName=$1
Cmd=$2
echo $Cmd>$PJ/server/bin/cmd.in
kill -USR1 `pgrep $SrvName`
