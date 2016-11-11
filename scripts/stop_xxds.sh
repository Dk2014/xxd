#!/usr/bin/env bash

set -e

LANG=en_US.UTF-8

GRPNAME=$1

if [ -z $1 ] 
then
  echo "usage: stop_xxds.sh <process_group>"
  exit 1
fi

for PID in `ps auxww|grep "${GRPNAME}"|grep -v grep|grep -vF ".sh"|awk '{print $2}'`;do
  if [ ! -z $PID ];then
    kill -9 $PID
  fi
done
