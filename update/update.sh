#!/bin/bash

warning() { echo ; echo -e "\033[31m  $1  \033[0m"; echo ; }
info() { echo -e "\033[32m  $1  \033[0m"; }

runContainer() {
  sleep 0.5
  echo "开始生成 Docker 容器。耗时比较长，请耐心等待..."
   docker-compose up -d 
  info "生成 Docker 容器完成"
}


getIpAddr() {
  ipaddr=`ifconfig -a|grep inet|grep -v .0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"`
  array=(`echo $ipaddr | tr '\n' ' '`)
  num=${#array[@]}
  if [ $num -eq 1 ]; then
    #echo "*单网卡"
    local_ip=${array[*]}
  elif [ $num -gt 1 ];then
    local_ip=${array[1]}
  else
    warning "未设置网卡IP，请检查服务器环境！"
    exit 1
  fi
}


update()
{
  docker build -t hary654321/crocodile . -f DockerfileServer

  cd client
  docker build -t hary654321/scaner . -f DockerfileScaner

  cd -

  runContainer
  getIpAddr
  info "后台地址：http://${local_ip}:61665/crocodile/"
  info '更新顺利完成！'
}

update


