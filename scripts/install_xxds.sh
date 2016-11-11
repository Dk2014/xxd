#!/usr/bin/env bash

set -e

LANG=en_US.UTF-8
BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

VERSION=$1


curl -O -v "http://115.182.64.106/xxds/xxds-linux-amd64-${VERSION}.tgz" -H "Host: builder.xxd.pinidea.us"

mkdir xxds_$VERSION

cd xxds_$VERSION

tar -xzf "../xxds-linux-amd64-${VERSION}.tgz"

cd ..

rsync -av "xxds_${VERSION}/game_server" ./
rsync -av "xxds_${VERSION}/idip_server" ./
rsync -av "xxds_${VERSION}/platform_server" ./
rsync -av "xxds_${VERSION}/platform_tool" ./

echo "mysql -uroot -S /data/mysqldata/20000/mysql.sock mobile_xxd003 < xxd_${VERSION}/database-game_server-upgrade.sql"
