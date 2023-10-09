#!/bin/bash
docker stop crocodile_server
docker stop crocodile_scaner
mkdir -p /zrtx/apt/scan
mkdir -p /app
cp -rf  ./client/* /zrtx/apt/scan
ulimit -n 50000
monitorM=`ps -ef | grep apt-server | grep -v grep | wc -l ` 
if [ $monitorM -eq 0 ] 
then
	echo "apt-server is not running, restart apt-server "
	 ./apt-server  wl.toml>>m.log &
else
	echo "apt-server is running"
fi


monitorS=`ps -ef | grep apt-scan| grep -v grep | wc -l ` 
if [ $monitorS -eq 0 ] 
then
	echo "apt-scan not running, restart apt-scan"
        cd /zrtx/apt/scan  && ./apt-scan >>s.log &
else
	echo "apt-scan is running"
fi

