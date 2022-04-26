#!/bin/bash

##进入脚本所在路径
basedir=`cd $(dirname $0); pwd -P`
cd $basedir

for i in {1..5} ; do
  sleep 1
    echo ${i}
done
