#!/bin/bash
array=("cyberspace-resources_hg" "cyberspace-resources_yg" "cyberspace-resources_zg" "cyberspace-resources_fg" "cyberspace-resources_gd1" "cyberspace-resources_ads2")

for element in "${array[@]}"
do


    ps -ef | grep qr | grep -v grep | awk '{print $2}' | xargs kill -9
    cd /u4/logstashqr

    yes|cp -f config/qr.base config/qr.conf
    sed -i "s/day/$element/g" config/qr.conf
    rm -rf data
    ./bin/logstash -f config/qr.conf



 done

