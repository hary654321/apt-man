#!/bin/bash
cd /zrtx/apt
chmod +x apt-scan-`uname -m`
ps -ef | grep ./apt-scan | grep -v grep | awk '{print $2}' | xargs kill -9
/zrtx/apt/apt-scan-`uname -m`  >> /zrtx/apt/s.log &