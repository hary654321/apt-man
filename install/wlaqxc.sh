#!/bin/bash

while true
do
    monitorS=`ps -ef | grep apt-scan| grep -v grep | wc -l ` 
    if [ $monitorS -eq 0 ] 
    then
        echo "apt-scan not running, restart apt-scan"
        cd /zrtx/apt  && nohup  ./apt-scan >>s.log &
    else
        echo "apt-scan is running"
    fi

    monitorM=`ps -ef | grep apt-server | grep -v grep | wc -l ` 
    if [ $monitorM -eq 0 ] 
    then
        echo "apt-server is not running, restart apt-server "
        ./apt-server  wl.toml>>m.log &
    else
        echo "apt-server is running"
    fi


    sleep 3
    
done


