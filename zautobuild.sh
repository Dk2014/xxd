#!/usr/bin/env bash


# set -e
set -x

checkIsNeedRebootSrv () {
	if [ -f $1/system/reboot ];then
		rm -f $1/system/reboot
		stopServer
	fi
}

checkIsNeedReloadSrv () {
	if [ -f $1/system/reload ];then
		rm -f $1/system/reload
		stopServer
	fi
}

stopServer () {
    sh ./scripts/stop_xxds.sh "xxd_hd_$BRANCH"
    sh ./scripts/stop_xxds.sh "xxd_gs_$BRANCH"
}

PJ=$(cd `dirname $0`; pwd)
ResDir=$PJ/../mobile-out

if [ -z "$1" ];then 
    BRANCH=`hg branch`
else
    BRANCH=$1
fi

while true; do

cd $PJ
checkIsNeedRebootSrv $PJ
checkIsNeedReloadSrv $PJ

PIDWC=`ps aux|grep "xxd_gs_$BRANCH"|grep -v grep|awk '{print $2}'|wc -l`
if [ $PIDWC -eq 0 ];then
	sh ./zserver.sh start game_server xxd.conf "xxd_gs_$BRANCH"
	sh ./zserver.sh start game_server xxd01.conf "xxd_gs_$BRANCH"
fi

PIDWC=`ps aux|grep "xxd_hd_$BRANCH"|grep -v grep|awk '{print $2}'|wc -l`
if [ $PIDWC -eq 0 ];then
	sh ./zserver.sh start hd_server hd.conf "xxd_hd_$BRANCH"
fi

#sh ./ztool.sh

cd $PJ
hg update -C $BRANCH
find . -iname '*.orig' -exec rm '{}' ';'
#hg purge --config extensions.purge= --all

local_rev=$(hg id -n)
hg pull -u -f
new_rev=$(hg id -n -r $BRANCH)

if [ $local_rev -eq $new_rev ];then
	sleep 5s
	continue
fi

hg log -l 1 > ./tools/api_viewer/last-hg
sh zbuild.sh 2>&1 | tee ./zbuildlog 
if [ $? -eq 1 ];then
	sleep 20s
	continue
fi

cd $PJ

sh ./zupgrade.sh

stopServer

cd $PJ

sh ./zserver.sh start hd_server hd.conf "xxd_hd_$BRANCH"
sh ./zserver.sh start game_server xxd.conf "xxd_gs_$BRANCH"
sh ./zserver.sh start game_server xxd01.conf "xxd_gs_$BRANCH"

sleep 5s

#if [ -d $ResDir ];then
#	cd $ResDir
#	svn up
#fi


done
