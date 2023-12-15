#!/bin/bash
cd /zrtx/apt
chmod +x apt-scan
ps -ef | grep ./apt-scan | grep -v grep | awk '{print $2}' | xargs kill -9
/zrtx/apt/apt-scan  >> /zrtx/apt/s.log &