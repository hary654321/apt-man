#!/bin/bash

cp -f ./docker/d*  /usr/bin
cp -f ./docker/c*  /usr/bin
cp -f ./docker/r*  /usr/bin
chmod +x /usr/bin/*
mkdir -p /usr/lib/systemd/system
cp -f ./bin/docker.service /usr/lib/systemd/system/docker.service
chmod 755 /usr/lib/systemd/system/docker.service

systemctl daemon-reload
systemctl enable --now docker
#  	systemctl unmask docker.service
#    systemctl unmask docker.socket
systemctl start docker.service
