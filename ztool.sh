#!/usr/bin/env bash


set -x

PJ=$(cd `dirname $0`; pwd)

cd $PJ/../mobile-drama

svn up --accept theirs-full

php ./tool/drama_config/drama.php
php ./tool/assets_version/version.php