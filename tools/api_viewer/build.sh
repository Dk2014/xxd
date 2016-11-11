#!/usr/bin/env bash


CUR=$(cd `dirname $0`; pwd)
OUT=$CUR/genclient
DB=$CUR/../../database
DOC=$CUR/../../protocol


if [ ! -d $OUT ];then
	mkdir $OUT
fi

echo "generate protodoc..."
cd $DOC
php protodoc.php invoke game_server client_code $OUT


echo "generate code..."
cd $DB
php database.php invoke local_game client_data


cd $OUT

if [ -f $OUT/data.zip ];then
	rm -rf $OUT/data.zip
fi
zip -rq $OUT/data.zip ./
