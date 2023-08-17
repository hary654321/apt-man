#!/bin/bash

cd ..

go-bindata -o=core/utils/asset/asset.go -pkg=asset web/crocodile/... sql

go build -o ./install/apt-server  main.go 

cd /mnt/hgfs/go-pro/apt-scan
rm -f ../cyberspacemapping/www/strategy-manage/install/client/apt-scan
go build -o ../cyberspacemapping/www/strategy-manage/install/client/apt-scan main.go 


cd /mnt/hgfs/go-pro/cyberspacemapping/www/strategy-manage/install/client
# docker build -t hary654321/scaner . -f DockerfileScaner

cd /mnt/hgfs/go-pro/cyberspacemapping/www/strategy-manage/install
# docker build -t hary654321/crocodile . -f DockerfileServer


# docker image save -o crocodile.tar hary654321/crocodile mysql:5.7 redis:latest

cd client

tar -cvf ../client.tar  *

cd ..


tar -cvf ../install.tar  *

tar -cvf update.tar client apt-server  DockerfileServer update.sh 