#!/usr/bin/env bash

set -e

LANG=en_US.UTF-8

G_NAME=game_server
P_NAME=platform_server

GRPNAME=$1
BASEDIR=$2
GSCONF=$3
HDCONF=$4
PPORT=$5
DBHOST=$6
DBUSER=$7
DBPSWD=$8
DBPORT=$9
DBNAME_GS=${10}
DBNAME_PLAT=${11}

if [ -z ${11} ] 
then
  echo "usage: ./start_xxds.sh <group_name> <basedir> <gsconf> <hdconf> <platform_port> <db_host> <db_usrname> <db_passwd> <db_port> <dbname_gs> <dbname_platform> "
  exit 1
fi

BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $BASE
cd $BASEDIR

for CONF in `echo $GSCONF | tr "," "\n"`
do
    echo -e $CONF "\n"
    echo -e 'GOMAXPROCS=16 GODEBUG="gctrace=1" nohup ./'$G_NAME' -conf='$BASE'/'$CONF' gs_'$GRPNAME' 1> game_server_'$CONF'.out 2> game_server_'$CONF'.log &'"\n"
    GOMAXPROCS=16 GODEBUG="gctrace=1" nohup ./$G_NAME -conf=$BASE/$CONF gs_$GRPNAME 1> game_server_$CONF.out 2> game_server_$CONF.log &
done

echo -e 'GOMAXPROCS=16 GODEBUG="gctrace=1" nohup ./'$G_NAME' -conf='$BASE'/'$HDCONF' hd_'$GRPNAME' 1> hd_server.out 2> hd_server.log &'"\n"
GOMAXPROCS=16 GODEBUG="gctrace=1" nohup ./$G_NAME -conf=$BASE/$HDCONF hd_$GRPNAME 1> hd_server.out 2> hd_server.log &
echo -e 'GOMAXPROCS=16 GODEBUG="gctrace=1" nohup ./'$P_NAME' -port='$PPORT' -db_name='$DBNAME_PLAT' -db_port='$DBPORT' -db_host='$DBHOST' -db_user='$DBUSER' -db_pass='$DBPSWD' '$GRPNAME' 1> platform_server.out 2> platform_server.log &'"\n"
GOMAXPROCS=16 GODEBUG="gctrace=1" nohup ./$P_NAME -port=$PPORT -db_name=$DBNAME_PLAT -db_port=$DBPORT -db_host=$DBHOST -db_user=$DBUSER -db_pass=$DBPSWD $GRPNAME 1> platform_server.out 2> platform_server.log &
