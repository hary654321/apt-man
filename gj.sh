#!/bin/bash

array=("cyberspace-resources_hg" "cyberspace-resources_gd1" "cyberspace-resources_fg" "cyberspace-resources_ads2")

for element in "${array[@]}"
do
    echo $element
    ps -ef | grep cyqy | grep -v grep | awk '{print $2}' | xargs kill -9
    cd /u4/logstash-7.9.3

    yes|cp -f config/cyqy.base config/cyqy.conf
    sed -i "s/day/$element/g" config/cyqy.conf
    rm -rf data
    ./bin/logstash -f config/cyqy.conf

done