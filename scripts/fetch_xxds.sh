#!/usr/bin/env bash

set -e

LANG=en_US.UTF-8
BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

VERSION=$1
TARGETDIR=$2

if [ -z $2 ] 
then
  echo "usage: ./fetch_xxds.sh <version> <output_dir> <db_host> <db_usrname> <db_passwd> <db_port> <dbname_gs> <dbname_platform> "
  exit 0
fi

rm -rf $TARGETDIR
mkdir $TARGETDIR
cd $TARGETDIR

curl --compressed "http://builder.xxd.pinidea.us/xxds/xxds-linux-amd64-${VERSION}.tgz" | tar -xz

cd $BASE

echo -e "binary file extracted\n"

echo -e "executing .sql file to upgrade mysql\n"



