#! /bin/bash
git pull
go build -o apt-man main.go
ps -ef | grep ./apt-man | grep -v grep | awk '{print $2}' | xargs kill -9
cd /u2/apt-man && ./apt-man 138.toml>> m.log &
