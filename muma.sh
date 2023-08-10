#! /bin/bash
/muma/gost/gost -L=:8081 &
/muma/frp/frps &
/muma/fuso/fus &
/muma/gortcp/server &
/muma/nps/nps &
/muma/termite/agent  -l 8888 &

