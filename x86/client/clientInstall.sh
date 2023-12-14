#!/bin/bash
mkdir -p  /zrtx/config/cyberspace
mkdir -p  /zrtx/log/cyberspace
cd /zrtx/apt
chmod +x apt-scan
#go build -o worker worker.go
ps -ef | grep ./apt-scan | grep -v grep | awk '{print $2}' | xargs kill -9
./apt-scan  >> /zrtx/log/cyberspace/apt-scan.log &