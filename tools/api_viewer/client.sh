#!/usr/bin/env bash

set -e

if [ -z "$1" ];then
cat<<HELP
like this:
\$ $0 [start|stop|restart|update|upgrade]

you can control this:
    start [chenjian|xiong|lujian|liangdong|lizijian|liyupeng]  -- start game server
    stop  -- you know
    restart -- you know
    update -- update game_server
    upgrade -- update myself

HELP
exit 0
else
 case $1 in
         start)
            if [ -z $2 ];then
                echo "input your name!"
                exit 0
            fi
            ./game_server -db_host=svn.vvv.io -db_user=client -db_name=mobile_xxd_$2 &
            ;;
         stop)
            kill -15 `pgrep game_server`
            ;;
        restart)
            $0 stop
            $0 start $2
            ;;
        update)
            curl -fL http://svn.vvv.io:8765/api_viewer/game_server -o ./game_server
            sudo chmod +x ./game_server
            ;;
        upgrade)
            curl -fL http://svn.vvv.io:8765/api_viewer/client.sh >$0 && chmod +x $0
            ;;            
         *)echo "what???????? no zuo no die!";;
 esac
fi