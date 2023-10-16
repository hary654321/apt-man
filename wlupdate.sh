#!/bin/bash
docker stop crocodile_server
docker stop crocodile_scaner

ps -ef | grep apt-scan| grep -v grep | awk '{print $2}' | xargs kill -9
ps -ef | grep apt-server| grep -v grep | awk '{print $2}' | xargs kill -9
mkdir -p /zrtx/apt
mkdir -p /app
yes | cp -rf  ./client/* /zrtx/apt
ulimit -n 50000

nohup ./apt-server  wl.toml>>m.log &
cd /zrtx/apt  && nohup  ./apt-scan >>s.log &

cd - && nohup /wlaqxc.sh &





