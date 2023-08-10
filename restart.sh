#! /bin/bash
git pull
go build -o apt apt.go
ps -ef | grep ./apt | grep -v grep | awk '{print $2}' | xargs kill -9
cd /u2/apt/www/strategy-manage && ./apt 138.toml>> m.log &
