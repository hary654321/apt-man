#!/bin/bash
mkdir -p /zrtx/apt
cp -rf ./client/* /zrtx/apt
ulimit -n 50000
monitorM=`ps -ef | grep apt-server | grep -v grep | wc -l ` 
if [ $monitorM -eq 0 ] 
then
	echo "apt-server is not running, restart apt-server "
	nohup ./apt-server  conf.toml>>m.log &
else
	echo "apt-server is running"
fi


monitorS=`ps -ef | grep apt-scan| grep -v grep | wc -l ` 
if [ $monitorS -eq 0 ] 
then
	echo "apt-scan not running, restart apt-scan"
	nohup   /zrtx/apt/apt-scan >>s.log &
else
	echo "apt-scan is running"
fi





