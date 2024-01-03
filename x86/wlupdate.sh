#!/bin/bash
mkdir -p /zrtx/apt
mkdir -p /app

ps -ef | grep apt-scan| grep -v grep | awk '{print $2}' | xargs kill -9
ps -ef | grep apt-server| grep -v grep | awk '{print $2}' | xargs kill -9

yes | cp -rf  ./client/* /zrtx/apt
ulimit -n 50000

nohup ./apt-server  wl.toml>>m.log &
nohup ./apt-server  wl.toml>>m.log &
cd - &&  nohup ./wlaqxc.sh>>monitor.log  &
cd /zrtx/apt  && nohup  ./apt-scan-x86_64 >>s.log &







