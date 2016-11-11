#!/usr/bin/env bash
HOST=$1
DB=$2;
USERNAME=$3;
OUTPUT=$4

mysqldump --skip-opt -h$1 $2 -u$3 -p > $4
