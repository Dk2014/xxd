#!/bin/bash

## useage: sh zstart.sh ./server/bin/xxd.conf /usr/local/Cellar/jemalloc/3.6.0/lib/libjemalloc.1.dylib

LD_PRELOAD=$2 GODEBUG="gctrace=1" nohup ./server/bin/game_server -conf=$1 1> game_server.out 2> game_server.log &