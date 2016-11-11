#!/usr/bin/env bash

set -e

#./xxds_setup.sh latest gs1.conf,gs2.conf hd.conf _xxds_server_000_ 8888 rdsyejaq3yejaq3.mysql.rds.aliyuncs.com xxd000 xxd000KJS7 3306 xxd000 xxd000_platform
#./xxds_setup.sh latest gs.conf hd.conf _xxds_server_000_ 8888 rdsyejaq3yejaq3.mysql.rds.aliyuncs.com xxd000 xxd000KJS7 3306 xxd000 xxd000_platform

LATEST_REV=$1
GSCONF=$2
HDCONF=$3
GRPNAME=$4
PPORT=$5
DBHOST=$6
DBUSER=$7
DBPSWD=$8
DBPORT=$9
DBNAME_GS=${10}
DBNAME_PLAT=${11}

if [ -z ${11} ] 
then
  echo "usage: ./xxds_setup.sh <version|latest> <gs.conf> <hd.conf> <process_groupname> <platform_port> <db_host> <db_usrname> <db_passwd> <db_port> <dbname_gs> <dbname_platform> "
  exit 0
fi

BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

LOCKFILE="/tmp/setup_xxds.isrunning"
LASTREVFILE="${BASE}/lastrev-xxds"

if [ ! -e $LOCKFILE ]
then
    echo $$ >"$LOCKFILE"
else
    PID=$(cat "$LOCKFILE")
    if kill -0 "$PID" >&/dev/null
    then
        echo "nightly upgrade still running"
        exit 0
    else
        echo $$ >"$LOCKFILE"
        echo "Warning: previous nightly-upgrade appears to have not finished correctly"
    fi
fi

case $LATEST_REV in
 latest)
  LATEST_REV=`curl http://builder.xxd.pinidea.us/xxds/lastrev.txt`
  
  if [ -e $LASTREVFILE ]
  then
      lastrev=$(cat "$LASTREVFILE")
      if [ $lastrev == $LATEST_REV ]
      then
          echo "nightly upgrade is not necessery"
          exit 0
      fi
  fi
 
 ;;
 *)
 ;;
esac


echo -e "fetch ${LATEST_REV} \n"

cd $BASE

TARGETDIR=./xxds_$LATEST_REV/ 

./fetch_xxds.sh $LATEST_REV ./xxds_$LATEST_REV/ 

cd $BASE

./stop_xxds.sh $GRPNAME

cd $BASE

echo -e "database-game_server-upgrade.sql\n"
mysql -h$DBHOST -u$DBUSER -p$DBPSWD -P$DBPORT --default-character-set=utf8 $DBNAME_GS < $TARGETDIR/database-game_server-upgrade.sql

echo -e "database-game_platform-upgrade.sql\n"
mysql -h$DBHOST -u$DBUSER -p$DBPSWD -P$DBPORT --default-character-set=utf8 $DBNAME_PLAT < $TARGETDIR/database-game_platform-upgrade.sql

cd $BASE

./start_xxds.sh "${GRPNAME}_${LATEST_REV}" $TARGETDIR $GSCONF $HDCONF $PPORT $DBHOST $DBUSER $DBPSWD $DBPORT $DBNAME_GS $DBNAME_PLAT

echo $LATEST_REV > "$LASTREVFILE"

echo -e "done\n"