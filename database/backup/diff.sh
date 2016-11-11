#!/usr/bin/env bash
diff $1 $2|grep INSERT|cut -d' ' -f  4,4|uniq -c
