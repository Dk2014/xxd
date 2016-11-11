#!/usr/bin/env bash


set -e

PJ=$(cd `dirname $0`; pwd)

SRVNAME=game_server
if [ ! -z "$2" ];then
   SRVNAME=$2
fi

EVENT=event.json
CONF=xxd.conf
if [ ! -z "$3" ];then
   CONF=$3
fi

GROUP=$4

if [ -z "$1" ];then
cat<<HELP
like this:
\$ $0 [start|stop|restart]

HELP
exit 0
else
 case $1 in
         start)
			   cd $PJ/server/bin
            
            if [ ! -d ./out-log ];then
               mkdir ./out-log
            fi

            if [ -f $CONF.log ];then
               mv $CONF.log ./out-log/$CONF-`date +%s`.log
            fi
   			
            nohup ./$SRVNAME -conf=$CONF -event=$EVENT $GROUP >$CONF.log 2>&1 &
            
            ;;
         stop)
            PID=`ps aux|grep "\./$SRVNAME"|grep -v grep|awk '{print $2}'`
            kill -15 $PID
            cd $PJ/server/bin/

            if [ ! -d ./prof ];then
               mkdir ./prof
            fi

            for Prof in `ls game_server_*.csv | tr " " "\n"`
            do
               if [ -f $Prof ]; then
                  ProfName=`basename ${Prof} .csv`
                  mv $Prof ./prof/$ProfName-`date +%s`.csv
               fi
            done 
            ;;
        restart)
            sh $0 stop $SRVNAME $CONF
            sh $0 start $SRVNAME $CONF
            ;;        
         *)echo "what???????? no zuo no die!";;
 esac
fi