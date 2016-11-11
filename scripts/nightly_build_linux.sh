#!/usr/bin/env bash

set -e

LANG=en_US.UTF-8
REPO="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

BRANCH=$1
PRJ=$2
BINS=$3

if [ -z $3 ] 
then
  echo "usage: ./nightly_build_linux.sh tip xxd-server /data/xxds/binaries"
  exit 0
fi

LOCKFILE="/tmp/nightlybuild-${PRJ}.isrunning"
LASTREVFILE="${REPO}/lastrev-${PRJ}"

if [ ! -e $LOCKFILE ]
then
    echo $$ >"$LOCKFILE"
else
    PID=$(cat "$LOCKFILE")
    if kill -0 "$PID" >&/dev/null
    then
        echo "nightly build still running"
        exit 0
    else
        echo $$ >"$LOCKFILE"
        echo "Warning: previous nightly-build appears to have not finished correctly"
    fi
fi

cd $REPO
cd $PRJ

hg update -C $BRANCH
oldrev=`hg id -n`
hg pull -u -f
hg update -C $BRANCH
rev=`hg id -n`
find . -iname '*.orig' -exec rm '{}' ';'

if [ -e $LASTREVFILE ]
then
    lastrev=$(cat "$LASTREVFILE")
    if [ $lastrev == $rev ]
    then
        echo "nightly build is not necessery"
        exit 0
    fi
fi

if [ $oldrev != $rev ]
then
  echo quiting for now... complie only when repo is already the lastest revision
  exit 0
fi


./zbuild.sh static
./zplatform_build.sh static

cd $REPO
rm -rf $BINS/$rev 
mkdir $BINS/$rev

cp $PRJ/server/bin/game_server $BINS/$rev/
cp $PRJ/server/bin/platform_server $BINS/$rev/

cd $PRJ/database 

php sql.php game_server 0 9999999999 $BINS/$rev/database-game_server-upgrade.sql

php sql.php game_platform 0 9999999999 $BINS/$rev/database-game_platform-upgrade.sql

cd $BINS/$rev/

echo "${rev}" > version
md5sum * > .md5sum

rm -f ../xxds-linux-amd64-$rev.tgz

tar -czf ../xxds-linux-amd64-$rev.tgz ./

rm -rf $BINS/$rev/

echo $rev > $BINS/lastrev.txt
echo $rev > "$LASTREVFILE"
